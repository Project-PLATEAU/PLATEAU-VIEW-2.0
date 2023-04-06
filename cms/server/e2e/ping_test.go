package e2e

import (
	"net/http"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/app"
)

func TestPingAPI(t *testing.T) {
	e := StartServer(t, &app.Config{
		Origins: []string{"https://example.com"},
	}, false, nil)

	e.OPTIONS("/api/ping").
		WithHeader("Origin", "https://example.com").
		Expect().
		Status(http.StatusNoContent).
		Header("Access-Control-Allow-Origin").
		Equal("https://example.com")

	r := e.GET("/api/ping").
		WithHeader("Origin", "https://example.com").
		Expect()

	r.Header("Cache-Control").
		Equal("private, no-store, no-cache, must-revalidate")

	r.Header("Access-Control-Allow-Origin").
		Equal("https://example.com")

	r.Status(http.StatusOK).
		JSON().
		String().
		Equal("pong")
}
