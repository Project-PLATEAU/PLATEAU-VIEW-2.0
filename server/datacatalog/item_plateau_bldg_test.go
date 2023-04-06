package datacatalog

import (
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/stretchr/testify/assert"
)

func TestBldgSetFrom(t *testing.T) {
	assert.Equal(t, &BldgSet{
		MaxLOD: &BldgSetLOD{
			LOD: 4,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod4.zip",
			},
		},
		LOD0: &BldgSetLOD{
			LOD: 0,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod0.zip",
			},
		},
		LOD1: &BldgSetLOD{
			LOD: 1,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod1.zip",
			},
		},
		LOD2: &BldgSetLOD{
			LOD: 2,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2.zip",
			},
			LowTexture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2_low_texture.zip",
			},
			NoTexture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2_no_texture.zip",
			},
		},
		LOD3: &BldgSetLOD{
			LOD: 3,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod3.zip",
			},
		},
		LOD4: &BldgSetLOD{
			LOD: 4,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod4.zip",
			},
		},
	}, BldgSetFrom([]*cms.PublicAsset{
		{
			URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod1.zip",
		},
		{
			URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2.zip",
		},
		{
			URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2_low_texture.zip",
		},
		{
			URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2_no_texture.zip",
		},
		{
			URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod3.zip",
		},
		{
			URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod0.zip",
		},
		{
			URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod4.zip",
		},
	}))
}

func TestBldgSet_Config(t *testing.T) {
	assert.Equal(t, DataCatalogItemConfig{
		Data: []DataCatalogItemConfigItem{
			{
				Name: "LOD0",
				URL:  "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod0/tileset.json",
				Type: "3dtiles",
			},
			{
				Name: "LOD1",
				URL:  "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod1/tileset.json",
				Type: "3dtiles",
			},
			{
				Name: "LOD2",
				URL:  "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2/tileset.json",
				Type: "3dtiles",
			},
			{
				Name: "LOD2（低解像度）",
				URL:  "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2_low_texture/tileset.json",
				Type: "3dtiles",
			},
			{
				Name: "LOD2（テクスチャなし）",
				URL:  "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2_no_texture/tileset.json",
				Type: "3dtiles",
			},
			{
				Name: "LOD3",
				URL:  "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod3/tileset.json",
				Type: "3dtiles",
			},
			{
				Name: "LOD4",
				URL:  "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod4/tileset.json",
				Type: "3dtiles",
			},
		},
	}, (&BldgSet{
		MaxLOD: &BldgSetLOD{
			LOD: 3,
			Texture: &cms.PublicAsset{
				URL:  "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod3.zip",
				Type: "3dtiles",
			},
		},
		LOD0: &BldgSetLOD{
			LOD: 0,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod0.zip",
			},
		},
		LOD1: &BldgSetLOD{
			LOD: 1,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod1.zip",
			},
		},
		LOD2: &BldgSetLOD{
			LOD: 2,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2.zip",
			},
			LowTexture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2_low_texture.zip",
			},
			NoTexture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod2_no_texture.zip",
			},
		},
		LOD3: &BldgSetLOD{
			LOD: 3,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod3.zip",
			},
		},
		LOD4: &BldgSetLOD{
			LOD: 4,
			Texture: &cms.PublicAsset{
				URL: "https://example.com/13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13102_chuo-ku_lod4.zip",
			},
		},
	}).Config())
}
