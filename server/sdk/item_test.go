package sdk

import (
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/stretchr/testify/assert"
)

var item = Item{
	ID:             "xxx",
	CityGML:        "citygml_assetid",
	MaxLOD:         "maxlod_assetid",
	MaxLODStatus:   "未実行",
	SDKPublication: "公開する",
}

var cmsitem = cms.Item{
	ID: "xxx",
	Fields: []cms.Field{
		{Key: "citygml", Type: "asset", Value: "citygml_assetid"},
		{Key: "max_lod", Type: "asset", Value: "maxlod_assetid"},
		{Key: "max_lod_status", Type: "select", Value: "未実行"},
		{Key: "sdk_publication", Type: "select", Value: "公開する"},
	},
}

func TestItem(t *testing.T) {
	assert.Equal(t, item, ItemFrom(cmsitem))
	assert.Equal(t, Item{}, ItemFrom(cms.Item{}))
	assert.Equal(t, cmsitem.Fields, item.Fields())
	assert.Equal(t, []cms.Field(nil), Item{}.Fields())
}
