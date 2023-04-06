package geospatialjp

import (
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/stretchr/testify/assert"
)

var item = Item{
	ID:                  "xxx",
	Prefecture:          "東京都",
	CityName:            "千代田区",
	CityGML:             "citygml_assetid",
	CityGMLGeoSpatialJP: "citygml_geospatialjp_assetid",
	Catalog:             "catalog_assetid",
	All:                 "all_assetid",
	ConversionStatus:    "実行中",
	CatalogStatus:       "完了",
	SDKPublication:      "公開する",
}

var cmsitem = cms.Item{
	ID: "xxx",
	Fields: []cms.Field{
		{Key: "prefecture", Type: "select", Value: "東京都"},
		{Key: "city_name", Type: "text", Value: "千代田区"},
		{Key: "citygml", Type: "asset", Value: "citygml_assetid"},
		{Key: "citygml_geospatialjp", Type: "asset", Value: "citygml_geospatialjp_assetid"},
		{Key: "catalog", Type: "asset", Value: "catalog_assetid"},
		{Key: "all", Type: "asset", Value: "all_assetid"},
		{Key: "conversion_status", Type: "select", Value: "実行中"},
		{Key: "catalog_status", Type: "select", Value: "完了"},
		{Key: "sdk_publication", Type: "select", Value: "公開する"},
	},
}

func TestItem(t *testing.T) {
	assert.Equal(t, item, ItemFrom(cmsitem))
	assert.Equal(t, Item{}, ItemFrom(cms.Item{}))
	assert.Equal(t, cmsitem.Fields, item.Fields())
	assert.Equal(t, []cms.Field(nil), Item{}.Fields())
}

func TestItem_SpecVersion(t *testing.T) {
	assert.Equal(t, 2.3, Item{Specification: "第2.3版"}.SpecVersion())
	assert.Equal(t, 2.0, Item{Specification: "第2版"}.SpecVersion())
	assert.Equal(t, 1.0, Item{Specification: "第1版"}.SpecVersion())
	assert.Equal(t, 0.0, Item{Specification: ""}.SpecVersion())
}
