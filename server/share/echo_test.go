package share

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	mockCMS(t)

	e := echo.New()
	g := e.Group("/share")
	assert.NoError(t, Echo(g, Config{
		CMSBase:  "https://cms.example.com",
		CMSToken: "token",
		CMSModel: "modelmodel",
	}))

	r := httptest.NewRequest("GET", "/share/prj/aaa", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"a":"b"}`, strings.TrimSpace(w.Body.String()))

	r = httptest.NewRequest("GET", "/share/prj/aaaa", nil)
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `"not found"`, strings.TrimSpace(w.Body.String()))

	r = httptest.NewRequest("POST", "/share/prj", strings.NewReader(`{"a":"b"}`))
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `"aaa"`, strings.TrimSpace(w.Body.String()))

	r = httptest.NewRequest("POST", "/share/prj", strings.NewReader(`---`))
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `"invalid json"`, strings.TrimSpace(w.Body.String()))
}

func mockCMS(t *testing.T) {
	t.Helper()

	httpmock.RegisterResponder("GET", "https://cms.example.com/api/items/aaa", func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") != "Bearer token" {
			return httpmock.NewBytesResponse(http.StatusUnauthorized, nil), nil
		}
		return httpmock.NewJsonResponse(http.StatusOK, map[string]any{"id": "aaa", "fields": []map[string]string{{"key": "data", "value": `{"a":"b"}`}}})
	})

	httpmock.RegisterResponder("GET", "https://cms.example.com/api/items/aaaa", lo.Must(httpmock.NewJsonResponder(http.StatusNotFound, "not found")))

	httpmock.RegisterResponder("POST", "https://cms.example.com/api/projects/prj/models/modelmodel/items", func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") != "Bearer token" {
			return httpmock.NewBytesResponse(http.StatusUnauthorized, nil), nil
		}
		return httpmock.NewJsonResponse(http.StatusOK, map[string]string{"id": "aaa"})
	})
}
