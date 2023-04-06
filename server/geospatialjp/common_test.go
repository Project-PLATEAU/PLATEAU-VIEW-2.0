package geospatialjp

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/eukarya-inc/reearth-plateauview/server/geospatialjp/ckan"
	"github.com/jarcoal/httpmock"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestService_CheckCatalog(t *testing.T) {
	ctx := context.Background()
	catalogData := lo.Must(os.ReadFile("testdata/xxxxx_xxx_catalog.xlsx"))
	cf := NewCatalogFile(lo.Must(excelize.OpenReader(bytes.NewReader(catalogData))))
	cf.DeleteSheet()
	catalogData2 := lo.Must(cf.File().WriteToBuffer()).Bytes()

	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder("GET", "https://example.com/catalog.xlsx", httpmock.NewBytesResponder(http.StatusOK, catalogData))
	httpmock.RegisterResponder("GET", "https://example.com/catalog2.xlsx", httpmock.NewBytesResponder(http.StatusOK, catalogData2))

	cmsm := &mockCMS{}
	s := &Services{
		CMS: cmsm,
	}

	// case1: catalog file with the metadata sheet
	assert.NoError(t, s.CheckCatalog(ctx, "prj", Item{
		ID:      "item",
		Catalog: "catalog",
	}))
	assert.Equal(t, cms.Item{
		ID: "item",
		Fields: []cms.Field{
			{Key: "catalog_status", Value: "完了", Type: "select"},
		},
	}, cmsm.item)
	cmsm.item = cms.Item{}

	// case2: catalog file with the metadata sheet
	assert.ErrorContains(t, s.CheckCatalog(ctx, "prj", Item{
		ID:      "item",
		Catalog: "catalog2",
	}), "G空間情報センター用メタデータシートが見つかりません。")
	assert.Equal(t, cms.Item{
		ID: "item",
		Fields: []cms.Field{
			{Key: "catalog_status", Value: "エラー", Type: "select"},
		},
	}, cmsm.item)
	cmsm.item = cms.Item{}

	// case3: no catalog file
	assert.ErrorContains(t, s.CheckCatalog(ctx, "prj", Item{
		ID: "item",
	}), "目録アセットの読み込みに失敗しました。該当アセットが削除されていませんか？")
	assert.Equal(t, cms.Item{
		ID: "item",
		Fields: []cms.Field{
			{Key: "catalog_status", Value: "エラー", Type: "select"},
		},
	}, cmsm.item)
	cmsm.item = cms.Item{}
}

func TestService_RegisterCkanResources(t *testing.T) {
	ctx := context.Background()
	catalogData := lo.Must(os.ReadFile("testdata/xxxxx_xxx_catalog.xlsx"))
	cf := NewCatalogFile(lo.Must(excelize.OpenReader(bytes.NewReader(catalogData))))
	cf.DeleteSheet()
	catalogData2 := lo.Must(cf.File().WriteToBuffer()).Bytes()

	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder("GET", "https://example.com/catalog.xlsx", httpmock.NewBytesResponder(http.StatusOK, catalogData))
	httpmock.RegisterResponder("GET", "https://example.com/catalog2.xlsx", httpmock.NewBytesResponder(http.StatusOK, catalogData2))

	cmsm := &mockCMS{}
	ckanm := ckan.NewMock("org", nil, nil)
	s := &Services{
		CMS:         cmsm,
		Ckan:        ckanm,
		CkanOrg:     "org",
		CkanPrivate: true,
	}

	// case1: upload all files of 第2.3版
	assert.NoError(t, s.RegisterCkanResources(ctx, Item{
		ID:            "item",
		Specification: "第2.3版",
		CityGML:       "citygml",
		Catalog:       "catalog",
		All:           "all",
	}))

	pkg, err := ckanm.ShowPackage(ctx, "plateau-12210-mobara-shi-2022")
	assert.NoError(t, err)
	assert.Equal(t, "plateau-12210-mobara-shi-2022", pkg.Name)
	assert.Equal(t, "TITLE", pkg.Title)
	assert.True(t, pkg.Private)
	assert.Greater(t, len(pkg.ThumbnailURL), 100)
	assert.Equal(t, 3, len(pkg.Resources))
	assert.Equal(t, "3D Tiles, MVT（v2）", pkg.Resources[0].Name)
	assert.Equal(t, "https://example.com/all.zip", pkg.Resources[0].URL)
	assert.Equal(t, "CityGML（v2）", pkg.Resources[1].Name)
	assert.Equal(t, "https://example.com/12210_mobara-shi_2022_citygml_1_lsld.zip", pkg.Resources[1].URL)
	assert.Equal(t, "データ目録（v2）", pkg.Resources[2].Name)
	assert.Equal(t, cms.Item{ID: "item", Fields: []cms.Field{{Key: "sdk_publication", Type: "select", Value: "公開する"}}}, cmsm.item)

	// case2: upload citygml and catalog of 第1版
	assert.ErrorContains(t, s.RegisterCkanResources(ctx, Item{
		Specification: "第1版",
		CityGML:       "citygml2",
		Catalog:       "catalog2",
	}), "目録ファイルにG空間情報センター用メタデータシートがありません。")

	// case3: upload citygml and catalog of 第1版 to an existing package
	ckanm = ckan.NewMock("org", []ckan.Package{
		{
			ID:       "plateau-12210-mobara-shi-2020",
			Name:     "plateau-12210-mobara-shi-2020",
			Title:    "TITLE?",
			OwnerOrg: "org",
		},
	}, []ckan.Resource{
		{
			ID:        "aaa",
			PackageID: "plateau-12210-mobara-shi-2020",
			Name:      "CityGML（v2）",
			URL:       "hogehoge",
		},
	})
	s = &Services{
		CMS:         cmsm,
		Ckan:        ckanm,
		CkanOrg:     "org",
		CkanPrivate: false,
	}

	assert.NoError(t, s.RegisterCkanResources(ctx, Item{
		Specification: "第2版",
		CityGML:       "citygml2",
		Catalog:       "catalog2",
	}))

	pkg, err = ckanm.ShowPackage(ctx, "plateau-12210-mobara-shi-2020")
	assert.NoError(t, err)
	assert.Equal(t, "plateau-12210-mobara-shi-2020", pkg.Name)
	assert.Equal(t, "TITLE?", pkg.Title)
	assert.False(t, pkg.Private)
	assert.Equal(t, 2, len(pkg.Resources))
	assert.Equal(t, "CityGML（v2） PATCHED", pkg.Resources[0].Name)
	assert.Equal(t, "https://example.com/12210_mobara-shi_2020_citygml_1_lsld.zip", pkg.Resources[0].URL)
	assert.Equal(t, "データ目録（v2）", pkg.Resources[1].Name)
}

type mockCMS struct {
	cms.Interface
	item cms.Item
}

func (c *mockCMS) UpdateItem(ctx context.Context, itemID string, fields []cms.Field) (*cms.Item, error) {
	c.item = cms.Item{
		ID:     itemID,
		Fields: fields,
	}
	return nil, nil
}

func (*mockCMS) Asset(ctx context.Context, id string) (*cms.Asset, error) {
	if id == "catalog" {
		return &cms.Asset{
			ID:  "catalog",
			URL: "https://example.com/catalog.xlsx",
		}, nil
	}
	if id == "catalog2" {
		return &cms.Asset{
			ID:  "catalog",
			URL: "https://example.com/catalog2.xlsx",
		}, nil
	}
	if id == "citygml" {
		return &cms.Asset{
			ID:  "citygml",
			URL: "https://example.com/12210_mobara-shi_2022_citygml_1_lsld.zip",
		}, nil
	}
	if id == "citygml2" {
		return &cms.Asset{
			ID:  "citygml",
			URL: "https://example.com/12210_mobara-shi_2020_citygml_1_lsld.zip",
		}, nil
	}
	if id == "all" {
		return &cms.Asset{
			ID:  "catalog",
			URL: "https://example.com/all.zip",
		}, nil
	}
	return nil, rerror.ErrNotFound
}
