package putil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCacheControlMiddleware(t *testing.T) {
	e := echo.New()
	e.Use(CacheControlMiddleware("max-age=180", true))
	e.GET("/aaa", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	r := httptest.NewRequest("GET", "/aaa", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, "hello", w.Body.String())
	assert.Equal(t, "max-age=180", w.Header().Get("Cache-Control"))

	e = echo.New()
	e.Use(CacheControlMiddleware("no-store", true))
	e.Use(CacheControlMiddleware("max-age=180", false))
	e.GET("/aaa", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	r = httptest.NewRequest("GET", "/aaa", nil)
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, "hello", w.Body.String())
	assert.Equal(t, "max-age=180", w.Header().Get("Cache-Control"))
}
