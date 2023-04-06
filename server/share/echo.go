package share

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
)

const (
	cmsModel        = "share"
	cmsDataFieldKey = "data"
)

type Config struct {
	CMSBase  string
	CMSToken string
	// optional
	CMSModel        string
	CMSDataFieldKey string
	Disable         bool
}

func (conf *Config) Default() {
	if conf.CMSModel == "" {
		conf.CMSModel = cmsModel
	}
	if conf.CMSDataFieldKey == "" {
		conf.CMSDataFieldKey = cmsDataFieldKey
	}
}

func Echo(g *echo.Group, conf Config) error {
	conf.Default()
	if conf.Disable {
		return nil
	}

	cmsapi, err := cms.New(conf.CMSBase, conf.CMSToken)
	if err != nil {
		return fmt.Errorf("share: failed to init cms: %w", err)
	}

	g.Use(middleware.CORS())

	g.GET("/:project/:id", func(c echo.Context) error {
		prj := c.Param("project")
		if prj == "" {
			return c.JSON(http.StatusNotFound, "not found")
		}

		res, err := cmsapi.GetItem(c.Request().Context(), c.Param("id"), false)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}

			log.Errorf("share: failed to get an item: %s", err)
			return c.JSON(http.StatusInternalServerError, "internal server error")
		}

		f := res.FieldByKey(conf.CMSDataFieldKey)
		if f == nil {
			log.Errorf("share: item got, but field %s does not contain: %+v", conf.CMSDataFieldKey, res)
			return c.JSON(http.StatusNotFound, "not found")
		}

		v, ok := f.Value.(string)
		if !ok {
			log.Errorf("share: item got, but field %s's value is not a string: %+v", conf.CMSDataFieldKey, res)
			return c.JSON(http.StatusNotFound, "not found")
		}

		return c.Blob(http.StatusOK, "application/json", []byte(v))
	})

	g.POST("/:project", func(c echo.Context) error {
		prj := c.Param("project")
		if prj == "" {
			return c.JSON(http.StatusNotFound, "not found")
		}

		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, "failed to read body")
		}

		if !json.Valid(body) {
			return c.JSON(http.StatusBadRequest, "invalid json")
		}

		res, err := cmsapi.CreateItemByKey(c.Request().Context(), prj, conf.CMSModel, []cms.Field{
			{Key: conf.CMSDataFieldKey, Type: "textarea", Value: string(body)},
		})
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}

			log.Errorf("share: failed to create an item: %s", err)
			return c.JSON(http.StatusInternalServerError, "internal server error")
		}

		return c.JSON(http.StatusOK, res.ID)
	}, middleware.BodyLimit("10M"))

	return nil
}
