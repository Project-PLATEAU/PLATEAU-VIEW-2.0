package e2e

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/reearth/reearth-cms/server/internal/app"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/samber/lo"
)

var (
	publicAPIProjectID    = id.NewProjectID()
	publicAPIModelID      = id.NewModelID()
	publicAPIItem1ID      = id.NewItemID()
	publicAPIItem2ID      = id.NewItemID()
	publicAPIItem3ID      = id.NewItemID()
	publicAPIItem4ID      = id.NewItemID()
	publicAPIAsset1ID     = id.NewAssetID()
	publicAPIAsset2ID     = id.NewAssetID()
	publicAPIAssetUUID    = uuid.NewString()
	publicAPIProjectAlias = "test-project"
	publicAPIModelKey     = "test-model"
	publicAPIModelKey2    = "test-model-2"
	publicAPIField1Key    = "test-field-1"
	publicAPIField2Key    = "asset"
	publicAPIField3Key    = "test-field-2"
	publicAPIField4Key    = "asset2"
)

func TestPublicAPI(t *testing.T) {
	e, repos := StartServerAndRepos(t, &app.Config{
		AssetBaseURL: "https://example.com",
	}, true, publicAPISeeder)

	// not found
	e.GET("/api/p/{project}/{model}", "invalid-alias", publicAPIModelKey).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Equal(map[string]any{"error": "not found"})

	e.GET("/api/p/{project}/{model}", publicAPIProjectAlias, publicAPIModelKey2).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Equal(map[string]any{"error": "not found"})

	e.GET("/api/p/{project}/{model}", publicAPIProjectAlias, "invalid-key").
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Equal(map[string]any{"error": "not found"})

	e.GET("/api/p/{project}/{model}/{item}", publicAPIProjectAlias, publicAPIModelKey, id.NewItemID()).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Equal(map[string]any{"error": "not found"})

	// ok
	e.GET("/api/p/{project}/{model}", publicAPIProjectAlias, publicAPIModelKey).
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(map[string]any{
			"results": []map[string]any{
				{
					"id":               publicAPIItem1ID.String(),
					publicAPIField1Key: "aaa",
					publicAPIField2Key: map[string]any{
						"type": "asset",
						"id":   publicAPIAsset1ID.String(),
						"url":  fmt.Sprintf("https://example.com/assets/%s/%s/aaa.zip", publicAPIAssetUUID[:2], publicAPIAssetUUID[2:]),
					},
				},
				{
					"id":               publicAPIItem2ID.String(),
					publicAPIField1Key: "bbb",
				},
				{
					"id":               publicAPIItem3ID.String(),
					publicAPIField1Key: "ccc",
					publicAPIField3Key: []string{"aaa", "bbb", "ccc"},
					publicAPIField4Key: []any{
						map[string]any{
							"type": "asset",
							"id":   publicAPIAsset1ID.String(),
							"url":  fmt.Sprintf("https://example.com/assets/%s/%s/aaa.zip", publicAPIAssetUUID[:2], publicAPIAssetUUID[2:]),
						},
					},
				},
			},
			"totalCount": 3,
			"hasMore":    false,
			"limit":      50,
			"offset":     0,
			"page":       1,
		})

	// offset pagination
	e.GET("/api/p/{project}/{model}", publicAPIProjectAlias, publicAPIModelKey).
		WithQuery("limit", "1").
		WithQuery("offset", "1").
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(map[string]any{
			"results": []map[string]any{
				{
					"id":               publicAPIItem2ID.String(),
					publicAPIField1Key: "bbb",
				},
			},
			"totalCount": 3,
			"hasMore":    false,
			"limit":      1,
			"offset":     1,
			"page":       2,
		})

	// cursor pagination
	e.GET("/api/p/{project}/{model}", publicAPIProjectAlias, publicAPIModelKey).
		WithQuery("start_cursor", publicAPIItem1ID.String()).
		WithQuery("page_size", "1").
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(map[string]any{
			"results": []map[string]any{
				{
					"id":               publicAPIItem2ID.String(),
					publicAPIField1Key: "bbb",
				},
			},
			"totalCount": 3,
			"hasMore":    true,
			"nextCursor": publicAPIItem2ID.String(),
		})

	e.GET("/api/p/{project}/{model}/{item}", publicAPIProjectAlias, publicAPIModelKey, publicAPIItem1ID).
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(map[string]any{
			"id":               publicAPIItem1ID.String(),
			publicAPIField1Key: "aaa",
			publicAPIField2Key: map[string]any{
				"type": "asset",
				"id":   publicAPIAsset1ID.String(),
				"url":  fmt.Sprintf("https://example.com/assets/%s/%s/aaa.zip", publicAPIAssetUUID[:2], publicAPIAssetUUID[2:]),
			},
		})

	e.GET("/api/p/{project}/{model}/{item}", publicAPIProjectAlias, "___", publicAPIItem1ID).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Equal(map[string]any{
			"error": "not found",
		})

	e.GET("/api/p/{project}/{model}/{item}", publicAPIProjectAlias, publicAPIModelKey, publicAPIItem4ID).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Equal(map[string]any{
			"error": "not found",
		})

	e.GET("/api/p/{project}/assets/{assetid}", publicAPIProjectAlias, publicAPIAsset1ID).
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(map[string]any{
			"type": "asset",
			"id":   publicAPIAsset1ID.String(),
			"url":  fmt.Sprintf("https://example.com/assets/%s/%s/aaa.zip", publicAPIAssetUUID[:2], publicAPIAssetUUID[2:]),
			"files": []string{
				fmt.Sprintf("https://example.com/assets/%s/%s/aaa/bbb.txt", publicAPIAssetUUID[:2], publicAPIAssetUUID[2:]),
			},
		})

	// make the project's assets private
	ctx := context.Background()
	prj := lo.Must(repos.Project.FindByID(ctx, publicAPIProjectID))
	prj.Publication().SetAssetPublic(false)
	lo.Must0(repos.Project.Save(ctx, prj))

	e.GET("/api/p/{project}/{model}", publicAPIProjectAlias, publicAPIModelKey).
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(map[string]any{
			"results": []map[string]any{
				{
					"id":               publicAPIItem1ID.String(),
					publicAPIField1Key: "aaa",
					// publicAPIField2Key should be removed
				},
				{
					"id":               publicAPIItem2ID.String(),
					publicAPIField1Key: "bbb",
				},
				{
					"id":               publicAPIItem3ID.String(),
					publicAPIField1Key: "ccc",
					publicAPIField3Key: []string{"aaa", "bbb", "ccc"},
					// publicAPIField4Key should be removed
				},
			},
			"totalCount": 3,
			"hasMore":    false,
			"limit":      50,
			"offset":     0,
			"page":       1,
		})

	e.GET("/api/p/{project}/{model}/{item}", publicAPIProjectAlias, publicAPIModelKey, publicAPIItem1ID).
		Expect().
		Status(http.StatusOK).
		JSON().
		Equal(map[string]any{
			"id":               publicAPIItem1ID.String(),
			publicAPIField1Key: "aaa",
			// publicAPIField2Key should be removed
		})

	// make the project private
	prj.Publication().SetScope(project.PublicationScopePrivate)
	lo.Must0(repos.Project.Save(ctx, prj))

	e.GET("/api/p/{project}/{model}", publicAPIProjectAlias, publicAPIModelKey).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Equal(map[string]any{
			"error": "not found",
		})

	e.GET("/api/p/{project}/{model}/{item}", publicAPIProjectAlias, publicAPIModelKey, publicAPIItem1ID).
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Equal(map[string]any{
			"error": "not found",
		})
}

func publicAPISeeder(ctx context.Context, r *repo.Container) error {
	uid := id.NewUserID()
	p1 := project.New().ID(publicAPIProjectID).Workspace(id.NewWorkspaceID()).Alias(publicAPIProjectAlias).Publication(
		project.NewPublication(project.PublicationScopePublic, true),
	).MustBuild()

	a := asset.New().ID(publicAPIAsset1ID).Project(p1.ID()).CreatedByUser(uid).Size(1).Thread(id.NewThreadID()).
		FileName("aaa.zip").UUID(publicAPIAssetUUID).MustBuild()
	af := asset.NewFile().Name("bbb.txt").Path("aaa/bbb.txt").Build()

	s := schema.New().NewID().Project(p1.ID()).Workspace(p1.Workspace()).Fields(schema.FieldList{
		schema.NewField(schema.NewText(nil).TypeProperty()).NewID().Key(key.New(publicAPIField1Key)).MustBuild(),
		schema.NewField(schema.NewAsset().TypeProperty()).NewID().Key(key.New(publicAPIField2Key)).MustBuild(),
		schema.NewField(schema.NewText(nil).TypeProperty()).NewID().Key(key.New(publicAPIField3Key)).Multiple(true).MustBuild(),
		schema.NewField(schema.NewAsset().TypeProperty()).NewID().Key(key.New(publicAPIField4Key)).Multiple(true).MustBuild(),
	}).MustBuild()

	m := model.New().ID(publicAPIModelID).Project(p1.ID()).Schema(s.ID()).Public(true).Key(key.New(publicAPIModelKey)).MustBuild()
	// not public model
	m2 := model.New().ID(publicAPIModelID).Project(p1.ID()).Schema(s.ID()).Key(key.New(publicAPIModelKey2)).Public(false).MustBuild()

	i1 := item.New().ID(publicAPIItem1ID).Model(m.ID()).Schema(s.ID()).Project(p1.ID()).Thread(id.NewThreadID()).User(uid).Fields([]*item.Field{
		item.NewField(s.Fields()[0].ID(), value.TypeText.Value("aaa").AsMultiple()),
		item.NewField(s.Fields()[1].ID(), value.TypeAsset.Value(a.ID()).AsMultiple()),
	}).MustBuild()

	i2 := item.New().ID(publicAPIItem2ID).Model(m.ID()).Schema(s.ID()).Project(p1.ID()).Thread(id.NewThreadID()).User(uid).Fields([]*item.Field{
		item.NewField(s.Fields()[0].ID(), value.TypeText.Value("bbb").AsMultiple()),
	}).MustBuild()

	i3 := item.New().ID(publicAPIItem3ID).Model(m.ID()).Schema(s.ID()).Project(p1.ID()).Thread(id.NewThreadID()).User(uid).Fields([]*item.Field{
		item.NewField(s.Fields()[0].ID(), value.TypeText.Value("ccc").AsMultiple()),
		item.NewField(s.Fields()[1].ID(), value.TypeAsset.Value(publicAPIAsset2ID).AsMultiple()),
		item.NewField(s.Fields()[2].ID(), value.NewMultiple(value.TypeText, []any{"aaa", "bbb", "ccc"})),
		item.NewField(s.Fields()[3].ID(), value.TypeAsset.Value(a.ID()).AsMultiple()),
	}).MustBuild()

	// not public
	i4 := item.New().ID(publicAPIItem4ID).Model(m.ID()).Schema(s.ID()).Project(p1.ID()).Thread(id.NewThreadID()).User(uid).Fields([]*item.Field{
		item.NewField(s.Fields()[0].ID(), value.TypeText.Value("ddd").AsMultiple()),
	}).MustBuild()
	// not public model
	i5 := item.New().ID(publicAPIItem1ID).Model(m2.ID()).Schema(s.ID()).Project(p1.ID()).Thread(id.NewThreadID()).User(uid).Fields([]*item.Field{
		item.NewField(s.Fields()[0].ID(), value.TypeText.Value("aaa").AsMultiple()),
		item.NewField(s.Fields()[1].ID(), value.TypeAsset.Value(a.ID()).AsMultiple()),
	}).MustBuild()

	lo.Must0(r.Project.Save(ctx, p1))
	lo.Must0(r.Asset.Save(ctx, a))
	lo.Must0(r.AssetFile.Save(ctx, a.ID(), af))
	lo.Must0(r.Schema.Save(ctx, s))
	lo.Must0(r.Model.Save(ctx, m))
	lo.Must0(r.Item.Save(ctx, i1))
	lo.Must0(r.Item.Save(ctx, i2))
	lo.Must0(r.Item.Save(ctx, i3))
	lo.Must0(r.Item.Save(ctx, i4))
	lo.Must0(r.Item.Save(ctx, i5))
	lo.Must0(r.Item.UpdateRef(ctx, i1.ID(), version.Public, version.Latest.OrVersion().Ref()))
	lo.Must0(r.Item.UpdateRef(ctx, i2.ID(), version.Public, version.Latest.OrVersion().Ref()))
	lo.Must0(r.Item.UpdateRef(ctx, i3.ID(), version.Public, version.Latest.OrVersion().Ref()))

	return nil
}
