package sdkapi

import (
	"net/url"
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestItems_DatasetResponse(t *testing.T) {
	assert.Equal(t, &DatasetResponse{
		Data: []*DatasetPref{
			{
				ID:    "東京都",
				Title: "東京都",
				Data: []DatasetCity{
					{
						ID:           "id",
						Title:        "千代田区",
						Description:  "description",
						FeatureTypes: []string{"bldg", "tran", "frn", "veg"},
					},
				},
			},
		},
	}, Items{
		{
			ID:          "id",
			Prefecture:  "東京都",
			CityName:    "千代田区",
			Description: "description",
			CityGML: &cms.PublicAsset{
				ID:                      "citygml",
				ArchiveExtractionStatus: "done",
			},
			Bldg:           []cms.PublicAsset{{}},
			Tran:           []cms.PublicAsset{{}},
			Frn:            []cms.PublicAsset{{}},
			Veg:            []cms.PublicAsset{{}},
			MaxLOD:         &cms.PublicAsset{URL: "https://example.com/csv"},
			SDKPublication: "公開する",
		},
	}.DatasetResponse())
}

func TestMaxLODColumns_Map(t *testing.T) {
	assert.Equal(t, MaxLODMap{
		"bldg": map[string]float64{
			"1": 1,
			"2": 1,
		},
		"veg": map[string]float64{
			"1": 2,
		},
		"frn": map[string]float64{
			"2": 2,
		},
	}, MaxLODColumns{
		{Code: "1", Type: "bldg", MaxLOD: 1},
		{Code: "2", Type: "bldg", MaxLOD: 1},
		{Code: "1", Type: "veg", MaxLOD: 2},
		{Code: "2", Type: "frn", MaxLOD: 2},
	}.Map())
}

func TestMaxLODMap_Files(t *testing.T) {
	assert.Equal(t, FilesResponse{
		"bldg": []File{
			{Code: "1", URL: "https://example.com/1_bldg_xxx.gml", MaxLOD: 1},
			{Code: "2", URL: "https://example.com/2_bldg_yyy.gml", MaxLOD: 1},
		},
		"veg": []File{
			{Code: "1", URL: "https://example.com/1_veg_zzz.gml", MaxLOD: 2},
		},
		"frn": nil,
	}, MaxLODMap{
		"bldg": map[string]float64{
			"2": 1,
			"1": 1,
		},
		"veg": map[string]float64{
			"1": 2,
		},
		"frn": map[string]float64{
			"2": 2,
		},
	}.Files([]*url.URL{
		lo.Must(url.Parse("https://example.com/1_bldg_xxx.gml")),
		lo.Must(url.Parse("https://example.com/2_bldg_yyy.gml")),
		lo.Must(url.Parse("https://example.com/1_veg_zzz.gml")),
	}))
}

func TestItemsFromIntegration(t *testing.T) {
	cmsitems := []cms.Item{
		{
			ID: "xxx",
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
						"archiveExtractionStatus": "done",
						"contentType":             "application/octet-stream",
						"createdAt":               "2023-03-01T00:00:00.00Z",
						"file": map[string]any{
							"contentType": "application/octet-stream",
							"name":        "c.zip",
							"path":        "/c.zip",
							"size":        1000,
						},
						"id":          "assetc",
						"name":        "b.zip",
						"previewType": "geo",
						"projectId":   "prj",
						"totalSize":   1000,
						"url":         "https://example.com/c.zip",
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
							"archiveExtractionStatus": "done",
							"contentType":             "application/octet-stream",
							"createdAt":               "2023-03-01T00:00:00.00Z",
							"file": map[string]any{
								"contentType": "application/octet-stream",
								"name":        "b.zip",
								"path":        "/b.zip",
								"size":        1000,
							},
							"id":          "asset",
							"name":        "b.zip",
							"previewType": "geo",
							"projectId":   "prj",
							"totalSize":   1000,
							"url":         "https://example.com/b.zip",
						},
					},
				},
				{
					Key:   "max_lod",
					Value: nil,
				},
				{
					Key:   "sdk_publication",
					Value: "公開する",
				},
			},
		},
	}

	items := ItemsFromIntegration(cmsitems)
	assert.Equal(t, Items{
		{
			ID:          "xxx",
			Prefecture:  "pref",
			CityName:    "city",
			Description: "desc",
			CityGML: &cms.PublicAsset{
				Type:                    "asset",
				ID:                      "assetc",
				URL:                     "https://example.com/c.zip",
				ContentType:             "application/octet-stream",
				ArchiveExtractionStatus: "done",
			},
			MaxLOD: nil,
			Bldg: []cms.PublicAsset{
				{
					Type:                    "asset",
					ID:                      "asset",
					URL:                     "https://example.com/b.zip",
					ContentType:             "application/octet-stream",
					ArchiveExtractionStatus: "done",
				},
			},
			Tran:           []cms.PublicAsset{},
			Frn:            []cms.PublicAsset{},
			Veg:            []cms.PublicAsset{},
			SDKPublication: "公開する",
		},
	}, items)

}

func TestCityCode(t *testing.T) {
	assert.Equal(t, 123, cityCode(&cms.PublicAsset{
		URL: "https://example.com/aaa/123_aaa.zip",
	}))
	assert.Equal(t, 0, cityCode(&cms.PublicAsset{
		URL: "https://example.com/aaa/aaa_aaa.zip",
	}))
	assert.Equal(t, 0, cityCode(&cms.PublicAsset{
		URL: "",
	}))
	assert.Equal(t, 0, cityCode(nil))
}
