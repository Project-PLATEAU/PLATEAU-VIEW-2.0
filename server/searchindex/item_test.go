package searchindex

import (
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/stretchr/testify/assert"
)

var item = Item{
	ID:                "xxx",
	Bldg:              []string{"bldg_assetid", "bldg_assetid2"},
	SearchIndex:       []string{"searchindex_assetid"},
	SearchIndexStatus: "エラー",
}

var cmsitem = cms.Item{
	ID: "xxx",
	Fields: []cms.Field{
		{Key: "bldg", Type: "asset", Value: []string{"bldg_assetid", "bldg_assetid2"}},
		{Key: "search_index", Type: "asset", Value: []string{"searchindex_assetid"}},
		{Key: "search_index_status", Type: "select", Value: "エラー"},
	},
}

func TestItem(t *testing.T) {
	assert.Equal(t, item, ItemFrom(cmsitem))
	assert.Equal(t, Item{}, ItemFrom(cms.Item{}))
	assert.Equal(t, cmsitem.Fields, item.Fields())
	assert.Equal(t, []cms.Field(nil), Item{}.Fields())
}
