package datacatalog

import (
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/stretchr/testify/assert"
)

func TestTranSetFrom(t *testing.T) {
	a := TranSetFrom([]*cms.PublicAsset{
		{URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod0.zip"},
		{URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod1.zip"},
		{URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod2.zip"},
		{URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod3.zip"},
	})

	assert.Equal(t, TranSet{
		MaxLODN: 3,
		MaxLOD: &cms.PublicAsset{
			URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod3.zip",
		},
		LOD0: &cms.PublicAsset{
			URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod0.zip",
		},
		LOD1: &cms.PublicAsset{
			URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod1.zip",
		},
		LOD2: &cms.PublicAsset{
			URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod2.zip",
		},
		LOD3: &cms.PublicAsset{
			URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod3.zip",
		},
	}, a)
}

func TestTranSet_Config(t *testing.T) {
	assert.Equal(
		t,
		DataCatalogItemConfig{
			Data: []DataCatalogItemConfigItem{
				{
					Name:   "LOD0",
					URL:    "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod0/{z}/{x}/{y}.mvt",
					Type:   "mvt",
					Layers: []string{"Road"},
				},
				{
					Name:   "LOD1",
					URL:    "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod1/{z}/{x}/{y}.mvt",
					Type:   "mvt",
					Layers: []string{"Road"},
				},
				{
					Name:   "LOD2",
					URL:    "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod2/{z}/{x}/{y}.mvt",
					Type:   "mvt",
					Layers: []string{"TrafficArea", "AuxiliaryTrafficArea"},
				},
				{
					Name: "LOD3",
					URL:  "https://example.com/43204_arao-shi_2020_3dtiles_5_op_tran_lod3/tileset.json",
					Type: "3dtiles",
				},
			},
		},
		TranSet{
			MaxLODN: 3,
			MaxLOD: &cms.PublicAsset{
				URL: "https://example.com/43204_arao-shi_2020_3dtiles_5_op_tran_lod3.zip",
			},
			LOD0: &cms.PublicAsset{
				URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod0.zip",
			},
			LOD1: &cms.PublicAsset{
				URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod1.zip",
			},
			LOD2: &cms.PublicAsset{
				URL: "https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod2.zip",
			},
			LOD3: &cms.PublicAsset{
				URL: "https://example.com/43204_arao-shi_2020_3dtiles_5_op_tran_lod3.zip",
			},
		}.Config(),
	)
}
