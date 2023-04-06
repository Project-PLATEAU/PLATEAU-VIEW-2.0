package sidebar

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/eukarya-inc/reearth-plateauview/server/putil"
	"github.com/labstack/echo/v4"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
)

const (
	dataModelKey     = "sidebar-data"
	templateModelKey = "sidebar-template"
	dataField        = "data"
)

type Handler struct {
	CMS cms.Interface
}

func NewHandler(CMS cms.Interface) *Handler {
	return &Handler{
		CMS: CMS,
	}
}

// GET /:pid
func (h *Handler) fetchRoot() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		prj := c.Param("pid")

		c.Response().Header().Set(echo.HeaderCacheControl, "no-cache, must-revalidate")

		if hit, err := h.lastModified(c, prj, dataModelKey, templateModelKey); err != nil {
			return err
		} else if hit {
			return nil
		}

		data, err := h.CMS.GetItemsByKey(ctx, prj, dataModelKey, false)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		templates, err := h.CMS.GetItemsByKey(ctx, prj, templateModelKey, false)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]any{
			"data":      itemsToJSONs(data.Items),
			"templates": itemsToJSONs(templates.Items),
		})
	}
}

// GET /:pid/data
func (h *Handler) getAllDataHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		prj := c.Param("pid")

		c.Response().Header().Set(echo.HeaderCacheControl, "no-cache, must-revalidate")

		if hit, err := h.lastModified(c, prj, dataModelKey); err != nil {
			return err
		} else if hit {
			return nil
		}

		data, err := h.CMS.GetItemsByKey(ctx, prj, dataModelKey, false)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		return c.JSON(http.StatusOK, itemsToJSONs(data.Items))
	}
}

// GET /:pid/data/:iid
func (h *Handler) getDataHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		itemID := c.Param("iid")
		if itemID == "" {
			return c.JSON(http.StatusNotFound, nil)
		}

		c.Response().Header().Set(echo.HeaderCacheControl, "no-cache, must-revalidate")

		item, err := h.CMS.GetItem(ctx, itemID, false)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		res := itemJSON(item.FieldByKey(dataField), item.ID)
		if res == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		return c.JSON(http.StatusOK, res)
	}
}

// POST /:pid/data
func (h *Handler) createDataHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		prj := c.Param("pid")
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		if !json.Valid(b) {
			return errors.New("invalid json")
		}

		fields := []cms.Field{{
			Key:   dataField,
			Value: string(b),
		}}
		item, err := h.CMS.CreateItemByKey(ctx, prj, dataModelKey, fields)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		res := itemJSON(item.FieldByKey(dataField), item.ID)
		if res == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		return c.JSON(http.StatusOK, res)
	}
}

// PATCH /:pid/data/:did
func (h *Handler) updateDataHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		itemID := c.Param("iid")
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		if !json.Valid(b) {
			return errors.New("invalid json")
		}

		fields := []cms.Field{{
			Key:   dataField,
			Value: string(b),
		}}

		item, err := h.CMS.UpdateItem(ctx, itemID, fields)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		res := itemJSON(item.FieldByKey(dataField), item.ID)
		if res == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		return c.JSON(http.StatusOK, res)
	}
}

// DELETE /:pid/data/:did
func (h *Handler) deleteDataHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		itemID := c.Param("iid")

		if err := h.CMS.DeleteItem(ctx, itemID); err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}
}

// GET /:pid/templates
func (h *Handler) fetchTemplatesHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		prj := c.Param("pid")

		c.Response().Header().Set(echo.HeaderCacheControl, "no-cache, must-revalidate")

		if hit, err := h.lastModified(c, prj, templateModelKey); err != nil {
			return err
		} else if hit {
			return nil
		}

		res, err := h.CMS.GetItemsByKey(ctx, prj, templateModelKey, false)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		return c.JSON(http.StatusOK, itemsToJSONs(res.Items))
	}
}

// GET /:pid/templates/:tid
func (h *Handler) fetchTemplateHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		templateID := c.Param("tid")
		template, err := h.CMS.GetItem(ctx, templateID, false)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		res := itemJSON(template.FieldByKey(dataField), template.ID)
		if res == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		return c.JSON(http.StatusOK, res)
	}
}

// POST /:pid/templates
func (h *Handler) createTemplateHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		prj := c.Param("pid")
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		if !json.Valid(b) {
			return errors.New("invalid json")
		}

		fields := []cms.Field{{
			Key:   dataField,
			Value: string(b),
		}}

		template, err := h.CMS.CreateItemByKey(ctx, prj, templateModelKey, fields)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		res := itemJSON(template.FieldByKey(dataField), template.ID)
		if res == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		return c.JSON(http.StatusOK, res)
	}
}

// PATCH /:id/templates/:tid
func (h *Handler) updateTemplateHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		templateID := c.Param("tid")
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		if !json.Valid(b) {
			return errors.New("invalid json")
		}

		fields := []cms.Field{{
			Key:   dataField,
			Value: string(b),
		}}

		template, err := h.CMS.UpdateItem(ctx, templateID, fields)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		res := itemJSON(template.FieldByKey(dataField), template.ID)
		if res == nil {
			return c.JSON(http.StatusNotFound, "not found")
		}

		return c.JSON(http.StatusOK, res)
	}
}

// DELETE /:id/templates/:tid
func (h *Handler) deleteTemplateHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		templateID := c.Param("tid")

		if err := h.CMS.DeleteItem(ctx, templateID); err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return c.JSON(http.StatusNotFound, "not found")
			}
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func itemsToJSONs(items []cms.Item) []any {
	return lo.FilterMap(items, func(i cms.Item, _ int) (any, bool) {
		j := itemJSON(i.FieldByKey(dataField), i.ID)
		return j, j != nil
	})
}

func itemJSON(f *cms.Field, id string) any {
	j, err := f.ValueJSON()
	if j == nil || err != nil {
		return nil
	}
	if f.ID != "" {
		if o, ok := j.(map[string]any); ok {
			o["id"] = id
			return o
		}
	}
	return j
}

func (h *Handler) lastModified(c echo.Context, prj string, models ...string) (bool, error) {
	mlastModified := time.Time{}

	for _, m := range models {
		model, err := h.CMS.GetModelByKey(c.Request().Context(), prj, m)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return false, c.JSON(http.StatusNotFound, "not found")
			}
			return false, err
		}

		if model != nil && mlastModified.Before(model.LastModified) {
			mlastModified = model.LastModified
		}
	}

	return putil.LastModified(c, mlastModified)
}
