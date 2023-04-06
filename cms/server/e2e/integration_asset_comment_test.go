package e2e

import (
	"net/http"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/app"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integrationapi"
)

// GET|/assets/{assetId}/comments
func TestIntegrationGetAssetCommentAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.GET("/api/assets/{assetId}/comments", id.NewAssetID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/assets/{assetId}/comments", id.NewAssetID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/assets/{assetId}/comments", id.NewAssetID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/assets/{assetId}/comments", id.NewAssetID()).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusNotFound)

	r := e.GET("/api/assets/{assetId}/comments", aid).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()
	r.Keys().Contains("comments")
	c := r.Value("comments").Array().First().Object()
	c.Value("id").Equal(icId.String())
	c.Value("authorId").Equal(uId.String())
	c.Value("authorType").Equal(integrationapi.User)
	c.Value("content").Equal("test comment")
}

// POST|/assets/{assetId}/comments
func TestIntegrationCreateAssetCommentAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.POST("/api/assets/{assetId}/comments", id.NewAssetID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.POST("/api/assets/{assetId}/comments", id.NewAssetID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.POST("/api/assets/{assetId}/comments", id.NewAssetID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.POST("/api/assets/{assetId}/comments", id.NewAssetID()).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusNotFound)

	c := e.POST("/api/assets/{assetId}/comments", aid).
		WithHeader("authorization", "Bearer "+secret).
		WithJSON(map[string]interface{}{
			"content": "test",
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	c.Value("authorId").Equal(iId)
	c.Value("authorType").Equal(integrationapi.Integrtaion)
	c.Value("content").Equal("test")
}

// PATCH|/assets/{assetId}/comments/{commentId}
func TestIntegrationUpdateAssetCommentAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.PATCH("/api/assets/{assetId}/comments/{commentId}", id.NewAssetID(), id.NewCommentID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.PATCH("/api/assets/{assetId}/comments/{commentId}", id.NewAssetID(), id.NewCommentID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.PATCH("/api/assets/{assetId}/comments/{commentId}", id.NewAssetID(), id.NewCommentID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.PATCH("/api/assets/{assetId}/comments/{commentId}", id.NewAssetID(), id.NewCommentID()).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusNotFound)

	r := e.PATCH("/api/assets/{assetId}/comments/{commentId}", aid, icId).
		WithHeader("authorization", "Bearer "+secret).
		WithJSON(map[string]interface{}{
			"content": "updated content",
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	r.Keys().
		Contains("id", "authorId", "authorType", "content", "createdAt")

	r.Value("id").Equal(icId.String())
	r.Value("authorId").Equal(uId)
	r.Value("authorType").Equal(integrationapi.User)
	r.Value("content").Equal("updated content")

}

// DELETE|/assets/{assetId}/comments/{commentId}
func TestIntegrationDeleteAssetCommentAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.DELETE("/api/assets/{assetId}/comments/{commentId}", id.NewAssetID(), id.NewCommentID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.DELETE("/api/assets/{assetId}/comments/{commentId}", aid, icId).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().Keys().
		Contains("id")

	e.GET("/api/assets/{assetId}/comments", aid).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusOK).
		JSON().Object().Value("comments").Array().Empty()
}
