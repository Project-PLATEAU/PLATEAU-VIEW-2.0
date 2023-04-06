package datacatalog

import (
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/stretchr/testify/assert"
)

func TestDescFromAsset(t *testing.T) {
	name, desc := descFromAsset(&cms.PublicAsset{
		URL: "https://example.com/aaa.zip",
	}, []string{
		"bbb.zip\n\nBBB",
		"aaa.zip\n\nAAA",
		"CCC",
	})
	assert.Equal(t, "", name)
	assert.Equal(t, "AAA", desc)

	name, desc = descFromAsset(nil, []string{
		"bbb.zip\n\nBBB",
		"aaa.zip\n\nAAA",
		"CCC",
	})
	assert.Equal(t, "", name)
	assert.Equal(t, "", desc)

	name, desc = descFromAsset(&cms.PublicAsset{
		URL: "https://example.com/aaa.zip",
	}, []string{
		"CCC",
	})
	assert.Equal(t, "", name)
	assert.Equal(t, "", desc)

	name, desc = descFromAsset(&cms.PublicAsset{
		URL: "https://example.com/aaa.zip",
	}, []string{
		"aaa.zip\n@name: CCC\n\naaaa\nbbbb",
	})
	assert.Equal(t, "CCC", name)
	assert.Equal(t, "aaaa\nbbbb", desc)

	name, desc = descFromAsset(&cms.PublicAsset{
		URL: "https://example.com/aaa.zip",
	}, []string{
		"aaa.zip\n\n@name: CCC\naaaa\nbbbb",
	})
	assert.Equal(t, "CCC", name)
	assert.Equal(t, "aaaa\nbbbb", desc)

	name, desc = descFromAsset(&cms.PublicAsset{
		URL: "https://example.com/aaa.zip",
	}, []string{
		"aaa.zip\n@name:CCC",
	})
	assert.Equal(t, "CCC", name)
	assert.Equal(t, "", desc)
}

func TestItemName(t *testing.T) {
	name, t2, t2en := itemName("建築物モデル", "xxx市", "", AssetName{
		Feature: "bldg",
	})
	assert.Equal(t, "建築物モデル（xxx市）", name)
	assert.Equal(t, "", t2)
	assert.Equal(t, "", t2en)

	name, t2, t2en = itemName("建築物モデル", "xxx市", "AAA", AssetName{
		Feature: "bldg",
	})
	assert.Equal(t, "AAA（xxx市）", name)
	assert.Equal(t, "", t2)
	assert.Equal(t, "", t2en)

	name, t2, t2en = itemName("都市計画決定情報モデル", "xxx市", "", AssetName{
		Feature:        "urf",
		UrfFeatureType: "DistrictPlan",
	})
	assert.Equal(t, "地区計画モデル（xxx市）", name)
	assert.Equal(t, "地区計画", t2)
	assert.Equal(t, "DistrictPlan", t2en)

	name, t2, t2en = itemName("都市計画決定情報モデル", "xxx市", "AAAA", AssetName{
		Feature:        "urf",
		UrfFeatureType: "DistrictPlan",
	})
	assert.Equal(t, "AAAA（xxx市）", name)
	assert.Equal(t, "地区計画", t2)
	assert.Equal(t, "DistrictPlan", t2en)

	name, t2, t2en = itemName("都市計画決定情報モデル", "xxx市", "", AssetName{
		Feature:        "urf",
		UrfFeatureType: "XXX",
	})
	assert.Equal(t, "XXX（xxx市）", name)
	assert.Equal(t, "XXX", t2)
	assert.Equal(t, "XXX", t2en)
}
