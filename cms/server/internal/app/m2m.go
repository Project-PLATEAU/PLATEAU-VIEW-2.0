package app

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reearth/reearth-cms/server/internal/adapter"
	rhttp "github.com/reearth/reearth-cms/server/internal/adapter/http"
	"github.com/reearth/reearthx/log"
)

func NotifyHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input rhttp.NotifyInput
		var b pubsubBody
		if err := c.Bind(&b); err != nil {
			if err := c.Bind(&input); err != nil {
				return err
			}
		} else if data, err := b.Data(); err != nil {
			return err
		} else if err := json.Unmarshal(data, &input); err != nil {
			return err
		}

		ctx := c.Request().Context()

		log.Infof("notified and updating files begin: assetID=%s type=%s status=%s", input.AssetID, input.Type, input.Status)
		assetUC := adapter.Usecases(ctx).Asset
		controller := rhttp.NewTaskController(assetUC)
		if err := controller.Notify(ctx, input); err != nil {
			log.Errorf("failed to update files: assetID=%s, type=%s, status=$s", input.AssetID, input.Type, input.Status)
			return err
		}
		log.Infof("successfully notified and files has been updated: assetID=%s, type=%s, status=$s", input.AssetID, input.Type, input.Status)
		return c.JSON(http.StatusOK, "OK")
	}
}

type pubsubBody struct {
	Message struct {
		Data string `json:"data"`
	} `json:"message"`
}

func (b pubsubBody) Data() ([]byte, error) {
	if b.Message.Data == "" {
		return nil, nil
	}

	return base64.StdEncoding.DecodeString(b.Message.Data)
}
