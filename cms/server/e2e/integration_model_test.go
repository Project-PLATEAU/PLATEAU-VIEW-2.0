package e2e

import (
	"net/http"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/app"
	"github.com/reearth/reearth-cms/server/pkg/id"
)

// GET /models/{modelId}
func TestIntegrationModelGetAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.GET("/api/models/{modelId}", id.NewModelID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/models/{modelId}", id.NewModelID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/models/{modelId}", id.NewModelID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/models/{modelId}", id.NewModelID()).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusNotFound)

	// key cannot be used
	e.GET("/api/models/{modelId}", ikey).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusBadRequest)

	obj := e.GET("/api/models/{modelId}", mId).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("id", mId.String()).
		ValueEqual("name", "m1").
		ValueEqual("description", "m1 desc").
		ValueEqual("public", true).
		ValueEqual("key", ikey.String()).
		ValueEqual("projectId", pid).
		ValueEqual("schemaId", sid)

	obj.Value("createdAt").NotNull()
	obj.Value("updatedAt").NotNull()
	obj.Value("lastModified").NotNull()
}

// GET /projects/{projectIdOrAlias}/models/{modelIdOrKey}
func TestIntegrationModelGetWithProjectAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.GET("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}", palias, id.NewModelID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}", palias, id.NewModelID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}", palias, id.NewModelID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}", palias, id.NewModelID()).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusNotFound)

	obj := e.GET("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}", palias, mId).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("id", mId.String()).
		ValueEqual("name", "m1").
		ValueEqual("description", "m1 desc").
		ValueEqual("public", true).
		ValueEqual("key", ikey.String()).
		ValueEqual("projectId", pid).
		ValueEqual("schemaId", sid)

	obj.Value("createdAt").NotNull()
	obj.Value("updatedAt").NotNull()
	obj.Value("lastModified").NotNull()

	obj = e.GET("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}", palias, ikey).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("id", mId.String()).
		ValueEqual("name", "m1").
		ValueEqual("description", "m1 desc").
		ValueEqual("public", true).
		ValueEqual("key", ikey.String()).
		ValueEqual("projectId", pid).
		ValueEqual("schemaId", sid)

	obj.Value("createdAt").NotNull()
	obj.Value("updatedAt").NotNull()
	obj.Value("lastModified").NotNull()
}
