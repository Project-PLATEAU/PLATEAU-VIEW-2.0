package e2e

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/google/uuid"
	"github.com/reearth/reearth-cms/server/internal/app"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/samber/lo"
)

var (
	secret = "secret_1234567890"
	uId    = id.NewUserID()
	iId    = id.NewIntegrationID()
	mId    = id.NewModelID()
	aid    = id.NewAssetID()
	auuid  = uuid.NewString()
	itmId  = id.NewItemID()
	fId    = id.NewFieldID()
	fId2   = id.NewFieldID()
	thId   = id.NewThreadID()
	icId   = id.NewCommentID()
	ikey   = key.Random()
	pid    = id.NewProjectID()
	sid    = id.NewSchemaID()
	palias = "PROJECT_ALIAS"
	sfKey  = key.Random()
	sfKey2 = id.NewKey("asset")
)

func baseSeeder(ctx context.Context, r *repo.Container) error {
	u := user.New().ID(uId).
		Name("e2e").
		Email("e2e@e2e.com").
		MustBuild()
	if err := r.User.Save(ctx, u); err != nil {
		return err
	}

	i := integration.New().ID(iId).
		Type(integration.TypePublic).
		Name("i1").
		Description("i1 desc").
		LogoUrl(lo.Must(url.Parse("https://test.com"))).
		Token(secret).
		Developer(id.NewUserID()).
		MustBuild()
	if err := r.Integration.Save(ctx, i); err != nil {
		return err
	}

	w := user.NewWorkspace().NewID().
		Name("e2e").
		Personal(false).
		Integrations(map[user.IntegrationID]user.Member{i.ID(): {Role: user.RoleOwner, InvitedBy: u.ID()}}).
		MustBuild()
	if err := r.Workspace.Save(ctx, w); err != nil {
		return err
	}

	p := project.New().ID(pid).
		Name("p1").
		Description("p1 desc").
		ImageURL(lo.Must(url.Parse("https://test.com"))).
		Workspace(w.ID()).
		Alias(palias).
		MustBuild()
	if err := r.Project.Save(ctx, p); err != nil {
		return err
	}

	sf := schema.NewField(schema.NewText(nil).TypeProperty()).ID(fId).Key(sfKey).MustBuild()
	sf2 := schema.NewField(schema.NewAsset().TypeProperty()).ID(fId2).Key(sfKey2).MustBuild()
	s := schema.New().ID(sid).
		Workspace(w.ID()).
		Project(p.ID()).
		Fields([]*schema.Field{sf, sf2}).
		MustBuild()
	if err := r.Schema.Save(ctx, s); err != nil {
		return err
	}

	m := model.New().
		ID(mId).
		Name("m1").
		Description("m1 desc").
		Public(true).
		Key(ikey).
		Project(p.ID()).
		Schema(s.ID()).
		MustBuild()
	if err := r.Model.Save(ctx, m); err != nil {
		return err
	}

	itm := item.New().ID(itmId).
		Schema(s.ID()).
		Model(m.ID()).
		Project(p.ID()).
		Thread(thId).
		Fields([]*item.Field{
			item.NewField(fId2, value.TypeAsset.Value(aid).AsMultiple()),
		}).
		MustBuild()
	if err := r.Item.Save(ctx, itm); err != nil {
		return err
	}

	cmt := thread.NewComment(icId, operator.OperatorFromUser(u.ID()), "test comment")
	th := thread.New().ID(thId).Workspace(w.ID()).Comments([]*thread.Comment{cmt}).MustBuild()
	if err := r.Thread.Save(ctx, th); err != nil {
		return err
	}

	f := asset.NewFile().Name("aaa.jpg").Size(1000).ContentType("image/jpg").Build()
	a := asset.New().ID(aid).
		Project(p.ID()).
		CreatedByUser(u.ID()).
		FileName("aaa.jpg").
		Size(1000).
		UUID(auuid).
		Thread(thId).
		MustBuild()

	if err := r.Asset.Save(ctx, a); err != nil {
		return err
	}

	if err := r.AssetFile.Save(ctx, a.ID(), f); err != nil {
		return err
	}

	return nil
}

// GET /models/{modelId}/items
func TestIntegrationItemListAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.GET("/api/models/{modelId}/items", id.NewModelID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/models/{modelId}/items", id.NewModelID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/models/{modelId}/items", id.NewModelID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/models/{modelId}/items", id.NewModelID()).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusNotFound)

	obj := e.GET("/api/models/{modelId}/items", mId).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("page", 1).
		ValueEqual("perPage", 5).
		ValueEqual("totalCount", 1)

	a := obj.Value("items").Array()
	a.Length().Equal(1)
	assertItem(a.First(), false)

	// asset embeded
	obj = e.GET("/api/models/{modelId}/items", mId).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		WithQuery("asset", "true").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("page", 1).
		ValueEqual("perPage", 5).
		ValueEqual("totalCount", 1)

	a = obj.Value("items").Array()
	a.Length().Equal(1)
	assertItem(a.First(), true)

	// key cannot be used
	e.GET("/api/models/{modelId}/items", ikey).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusBadRequest)
}

// GET /projects/{projectIdOrAlias}/models/{modelIdOrKey}/items
func TestIntegrationItemListWithProjectAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.GET("/api/projects/{projectId}/models/{modelId}/items", pid, id.NewModelID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/projects/{projectId}/models/{modelId}/items", pid, id.NewModelID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/projects/{projectId}/models/{modelId}/items", pid, id.NewModelID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/projects/{projectId}/models/{modelId}/items", pid, id.NewModelID()).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusNotFound)

	obj := e.GET("/api/projects/{projectId}/models/{modelId}/items", pid, mId).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Value("page").Equal(1)
	obj.Value("perPage").Equal(5)
	obj.Value("totalCount").Equal(1)

	a := obj.Value("items").Array()
	a.Length().Equal(1)
	assertItem(a.First(), false)

	// model key can be also usable
	obj = e.GET("/api/projects/{projectId}/models/{modelId}/items", pid, ikey).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("page", 1).
		ValueEqual("perPage", 5).
		ValueEqual("totalCount", 1)

	a = obj.Value("items").Array()
	a.Length().Equal(1)
	assertItem(a.First(), false)

	// project alias can be also usable
	obj = e.GET("/api/projects/{projectId}/models/{modelId}/items", palias, ikey).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("page", 1).
		ValueEqual("perPage", 5).
		ValueEqual("totalCount", 1)

	a = obj.Value("items").Array()
	a.Length().Equal(1)
	assertItem(a.First(), false)

	// asset embeded
	obj = e.GET("/api/projects/{projectId}/models/{modelId}/items", pid, mId).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		WithQuery("asset", "true").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ValueEqual("page", 1).
		ValueEqual("perPage", 5).
		ValueEqual("totalCount", 1)

	a = obj.Value("items").Array()
	a.Length().Equal(1)
	assertItem(a.First(), true)

	// invalid key
	e.GET("/api/projects/{projectId}/models/{modelId}/items", pid, "xxx").
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusNotFound)

	// invalid project
	e.GET("/api/projects/{projectId}/models/{modelId}/items", id.NewProjectID(), ikey).
		WithHeader("authorization", "Bearer "+secret).
		WithQuery("page", 1).
		WithQuery("perPage", 5).
		Expect().
		Status(http.StatusNotFound)
}

// POST /models/{modelId}/items
func TestIntegrationCreateItemAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.POST("/api/models/{modelId}/items", id.NewModelID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.POST("/api/models/{modelId}/items", id.NewModelID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.POST("/api/models/{modelId}/items", id.NewModelID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.POST("/api/models/%s/items", id.NewModelID()).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusBadRequest)

	e.POST("/api/models/{modelId}/items", ikey).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusBadRequest)

	r := e.POST("/api/models/{modelId}/items", mId).
		WithHeader("authorization", "Bearer "+secret).
		WithJSON(map[string]interface{}{
			"fields": []interface{}{
				map[string]string{
					"id":    fId.String(),
					"value": "test value",
				},
			},
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()
	r.Keys().
		Contains("id", "modelId", "fields", "createdAt", "updatedAt", "version", "parents", "refs")
	r.Value("fields").Equal([]any{
		map[string]string{
			"id":    fId.String(),
			"type":  "text",
			"value": "test value",
			"key":   sfKey.String(),
		},
	})
	r.Value("modelId").Equal(mId.String())
	r.Value("refs").Equal([]string{"latest"})

	e.POST("/api/models/{modelId}/items", mId).
		WithHeader("authorization", "Bearer "+secret).
		WithJSON(map[string]interface{}{
			"fields": []interface{}{
				map[string]string{
					"key":   sfKey.String(),
					"value": "test value 2",
				},
			},
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Value("fields").
		Equal([]any{
			map[string]string{
				"id":    fId.String(),
				"type":  "text",
				"value": "test value 2",
				"key":   sfKey.String(),
			},
		})
}

// POST /projects/{projectIdOrAlias}/models/{modelIdOrKey}/items
func TestIntegrationCreateItemWithProjectAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.POST("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items", palias, id.NewModelID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.POST("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items", palias, id.NewModelID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.POST("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items", palias, id.NewModelID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.POST("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items", palias, id.NewModelID()).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusBadRequest)

	r := e.POST("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items", palias, ikey).
		WithHeader("authorization", "Bearer "+secret).
		WithJSON(map[string]interface{}{
			"fields": []interface{}{
				map[string]string{
					"id":    fId.String(),
					"value": "test value",
				},
			},
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()
	r.Keys().
		Contains("id", "modelId", "fields", "createdAt", "updatedAt", "version", "parents", "refs")
	r.Value("fields").Equal([]any{
		map[string]string{
			"id":    fId.String(),
			"type":  "text",
			"value": "test value",
			"key":   sfKey.String(),
		},
	})
	r.Value("modelId").Equal(mId.String())
	r.Value("refs").Equal([]string{"latest"})

	e.POST("/api/projects/{projectIdOrAlias}/models/{modelIdOrKey}/items", palias, ikey).
		WithHeader("authorization", "Bearer "+secret).
		WithJSON(map[string]interface{}{
			"fields": []interface{}{
				map[string]string{
					"key":   sfKey.String(),
					"value": "test value 2",
				},
			},
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Value("fields").
		Equal([]any{
			map[string]string{
				"id":    fId.String(),
				"type":  "text",
				"value": "test value 2",
				"key":   sfKey.String(),
			},
		})
}

// PATCH /items/{itemId}
func TestIntegrationUpdateItemAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.PATCH("/api/items/{itemId}", id.NewItemID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.PATCH("/api/items/{itemId}", id.NewItemID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.PATCH("/api/items/{itemId}", id.NewItemID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.PATCH("/api/items/{itemId}", id.NewItemID()).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusBadRequest)

	r := e.PATCH("/api/items/{itemId}", itmId).
		WithHeader("authorization", "Bearer "+secret).
		WithJSON(map[string]interface{}{
			"fields": []interface{}{
				map[string]string{
					"id":    fId.String(),
					"value": "test value",
				},
			},
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()
	r.Keys().
		Contains("id", "modelId", "fields", "createdAt", "updatedAt", "version", "parents", "refs")
	r.Value("fields").Equal([]interface{}{
		map[string]string{
			"id":    fId2.String(),
			"key":   "asset",
			"type":  "asset",
			"value": aid.String(),
		},
		map[string]string{
			"id":    fId.String(),
			"type":  "text",
			"value": "test value",
			"key":   sfKey.String(),
		},
	})
	r.Value("modelId").Equal(mId.String())
	r.Value("refs").Equal([]string{"latest"})

	e.PATCH("/api/items/{itemId}", itmId).
		WithHeader("authorization", "Bearer "+secret).
		WithJSON(map[string]interface{}{
			"fields": []interface{}{
				map[string]string{
					"id":    fId2.String(),
					"key":   "asset",
					"type":  "asset",
					"value": aid.String(),
				},
				map[string]string{
					"key":   sfKey.String(),
					"value": "test value 2",
				},
			},
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		Value("fields").
		Equal([]any{
			map[string]string{
				"id":    fId2.String(),
				"key":   "asset",
				"type":  "asset",
				"value": aid.String(),
			},
			map[string]string{
				"id":    fId.String(),
				"type":  "text",
				"value": "test value 2",
				"key":   sfKey.String(),
			},
		})
}

// GET /items/{itemId}
func TestIntegrationGetItemAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.GET("/api/items/{itemId}", id.NewItemID()).
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/items/{itemId}", id.NewItemID()).
		WithHeader("authorization", "secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/items/{itemId}", id.NewItemID()).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/items/{itemId}", id.NewItemID()).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusNotFound)

	e.GET("/api/items/{itemId}", itmId).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().Keys().
		Contains("id", "modelId", "fields", "createdAt", "updatedAt", "version", "parents", "refs")
}

// DELETE /items/{itemId}
func TestIntegrationDeleteItemAPI(t *testing.T) {
	e := StartServer(t, &app.Config{}, true, baseSeeder)

	e.DELETE("/api/items/{itemId}", itmId).
		WithHeader("authorization", "Bearer secret_abc").
		Expect().
		Status(http.StatusUnauthorized)

	e.GET("/api/items/{itemId}", itmId).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().Keys().
		Contains("id", "modelId", "fields", "createdAt", "updatedAt", "version", "parents", "refs")

	e.DELETE("/api/items/{itemId}", itmId).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().Keys().
		Contains("id")

	e.GET("/api/items/{itemId}", itmId).
		WithHeader("authorization", "Bearer "+secret).
		Expect().
		Status(http.StatusNotFound)
}

func assertItem(v *httpexpect.Value, assetEmbeded bool) {
	o := v.Object()
	o.Value("id").Equal(itmId.String())
	if assetEmbeded {
		a := o.Value("fields").Array()
		a.Length().Equal(1)
		a.First().Object().Value("value").Object().
			ValueEqual("id", aid.String()).
			NotContainsKey("contentType").
			NotContainsKey("file").
			NotContainsKey("name").
			ValueEqual("previewType", "unknown").
			ValueEqual("projectId", pid.String()).
			ValueEqual("totalSize", 1000).
			ValueEqual("url", fmt.Sprintf("https://example.com/assets/%s/%s/aaa.jpg", auuid[0:2], auuid[2:]))
	} else {
		o.Value("fields").Equal([]map[string]any{
			{
				"id":    fId2.String(),
				"key":   "asset",
				"type":  "asset",
				"value": aid.String(),
			},
		})
	}
	o.Value("parents").Equal([]any{})
	o.Value("refs").Equal([]string{"latest"})
}
