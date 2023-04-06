package putil

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CacheControlMiddleware(cacheControl string, capture bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if capture {
				c.Response().Header().Set(echo.HeaderCacheControl, cacheControl)
			}

			if err := next(c); err != nil {
				return err
			}

			if !capture && c.Response().Status == http.StatusOK {
				c.Response().Header().Set(echo.HeaderCacheControl, cacheControl)
			}

			return nil
		}
	}
}

var NoCacheMiddleware = CacheControlMiddleware("no-store", true)
