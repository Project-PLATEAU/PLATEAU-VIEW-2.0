package sdkapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/reearth/reearthx/rerror"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	data := `code,type,maxLod
53394452,bldg,1
53394453,bldg,1
53394461,bldg,1
53394462,bldg,1`
	httpmock.RegisterResponder("GET", "https://example.com/max_lod.csv", httpmock.NewBytesResponder(http.StatusOK, []byte(data)))

	e := echo.New()
	cms := NewCMS(&mockCMS{}, nil, "prj", false)
	assert.NoError(t, handler(Config{DisableCache: true}, e.Group(""), cms))

	// GET /dataset
	r := httptest.NewRequest("GET", "/datasets", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	body := map[string]any{}
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, map[string]any{
		"data": []any{
			map[string]any{
				"id":    "pref",
				"title": "pref",
				"data": []any{
					map[string]any{
						"featureTypes": []any{"bldg"},
						"id":           "item",
						"title":        "city",
						"description":  "desc",
					},
				},
			},
		},
	}, body)

	// GET /datase/item/files
	r = httptest.NewRequest("GET", "/datasets/item/files", nil)
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	body = map[string]any{}
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, map[string]any{
		"bldg": []any{
			map[string]any{
				"code":   "53394452",
				"maxLod": float64(1),
				"url":    "https://example.com/citygml/hoge/53394452_bldg_xxx.gml",
			},
			map[string]any{
				"code":   "53394453",
				"maxLod": float64(1),
				"url":    "https://example.com/citygml/hoge/53394453_bldg_xxx.gml",
			},
			map[string]any{
				"code":   "53394461",
				"maxLod": float64(1),
				"url":    "https://example.com/citygml/hoge/53394461_bldg_xxx.gml",
			},
		},
	}, body)
}

func TestGetMaxLOD(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	data := `code,type,maxLod
53394452,bldg,1
53394452,tran,1
53394453,bldg,1
53394453,tran,1
53394461,bldg,1
53394461,tran,1
53394462,bldg,1`
	httpmock.RegisterResponder("GET", "https://example.com", httpmock.NewBytesResponder(http.StatusOK, []byte(data)))

	ctx := context.Background()
	res, err := getMaxLOD(ctx, "https://example.com")
	assert.NoError(t, err)
	assert.Equal(t, MaxLODColumns{
		{Code: "53394452", Type: "bldg", MaxLOD: 1},
		{Code: "53394452", Type: "tran", MaxLOD: 1},
		{Code: "53394453", Type: "bldg", MaxLOD: 1},
		{Code: "53394453", Type: "tran", MaxLOD: 1},
		{Code: "53394461", Type: "bldg", MaxLOD: 1},
		{Code: "53394461", Type: "tran", MaxLOD: 1},
		{Code: "53394462", Type: "bldg", MaxLOD: 1},
	}, res)
}

type mockCMS struct {
	cms.Interface
}

func (c *mockCMS) Asset(ctx context.Context, id string) (*cms.Asset, error) {
	if id != "citygml" {
		return nil, rerror.ErrNotFound
	}
	return &cms.Asset{
		ID:                      "citygml",
		URL:                     "https://example.com/citygml.zip",
		ArchiveExtractionStatus: "done",
		File: &cms.File{
			Children: []cms.File{
				{Path: "/citygml/hoge/53394452_bldg_xxx.gml"},
				{Path: "/citygml/hoge/53394453_bldg_xxx.gml"},
				{Path: "/citygml/hoge/53394461_bldg_xxx.gml"},
				// {Path: "/bldg/hoge/53394462_bldg_xxx.gml"},
			},
		},
	}, nil
}

func (c *mockCMS) GetItem(ctx context.Context, itemID string, asset bool) (*cms.Item, error) {
	if itemID != "item" {
		return nil, rerror.ErrNotFound
	}
	return &mockItem, nil
}

func (c *mockCMS) GetItemsByKey(ctx context.Context, projectIDOrAlias, modelIDOrKey string, asset bool) (*cms.Items, error) {
	if projectIDOrAlias != "prj" || modelIDOrKey != "plateau" {
		return nil, rerror.ErrNotFound
	}
	return &cms.Items{
		Items: []cms.Item{mockItem, mockItem2},
	}, nil
}

var mockItem = cms.Item{
	ID: "item",
	Fields: []cms.Field{
		{
			Key:   "prefecture",
			Value: "pref",
		},
		{
			Key:   "city_name",
			Value: "city",
		},
		{
			Key: "citygml",
			Value: map[string]any{
				"id":                      "citygml",
				"archiveExtractionStatus": "done",
			},
		},
		{
			Key:   "description_bldg",
			Value: "desc",
		},
		{
			Key: "bldg",
			Value: []any{
				map[string]any{
					"url": "https://example.com/bldg.zip",
				},
			},
		},
		{
			Key: "max_lod",
			Value: map[string]any{
				"url": "https://example.com/max_lod.csv",
			},
		},
		{
			Key:   "sdk_publication",
			Value: "公開する",
		},
	},
}

var mockItem2 = cms.Item{
	ID: "aaa",
	Fields: []cms.Field{
		{
			Key:   "prefecture",
			Value: "pref",
		},
		{
			Key:   "city_name",
			Value: "city",
		},
		{
			Key: "citygml",
			Value: map[string]any{
				"id":                      "citygml",
				"archiveExtractionStatus": "done",
			},
		},
		{
			Key: "bldg",
			Value: []any{
				map[string]any{
					"url": "https://example.com/bldg.zip",
				},
			},
		},
		{
			Key: "max_lod",
			Value: map[string]any{
				"url": "https://example.com/max_lod.csv",
			},
		},
		// not published
	},
}
