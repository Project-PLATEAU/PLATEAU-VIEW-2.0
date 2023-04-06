package putil

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func LastModified(c echo.Context, lastModified time.Time) (bool, error) {
	if !lastModified.IsZero() {
		c.Response().Header().Set(echo.HeaderLastModified, lastModified.Format(time.RFC1123))
	}

	lm, _ := time.Parse(time.RFC1123, c.Request().Header.Get(echo.HeaderIfModifiedSince))
	if !lastModified.IsZero() && !lm.IsZero() && !lm.Before(lastModified) {
		return true, c.NoContent(http.StatusNotModified)
	}

	return false, nil
}
