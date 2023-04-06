package datacatalog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetNameFrom(t *testing.T) {
	assert.Equal(t, AssetName{
		CityCode: "13229",
		CityEn:   "nishitokyo-shi",
		Year:     "2022",
		Format:   "citygml",
		Op:       "1_op",
		Ext:      ".zip",
	}, AssetNameFrom("https://example.com/12345/13229_nishitokyo-shi_2022_citygml_1_op.zip"))

	assert.Equal(t, AssetName{
		CityCode: "22100",
		CityEn:   "shizuoka-shi",
		Year:     "2022",
		Format:   "citygml",
		Op:       "1_op",
		Ext:      ".zip",
	}, AssetNameFrom("22100_shizuoka-shi_2022_citygml_1_op.zip"))

	assert.Equal(t, AssetName{
		CityCode: "22100",
		CityEn:   "shizuoka-shi",
		Year:     "2022",
		Format:   "3dtiles",
		Op:       "op_1",
		Feature:  "bldg",
		WardCode: "22101",
		WardEn:   "aoi-ku",
		LOD:      "1",
		Ext:      ".zip",
	}, AssetNameFrom("22100_shizuoka-shi_2022_3dtiles_op_1_bldg_22101_aoi-ku_lod1.zip"))

	assert.Equal(t, AssetName{
		CityCode:  "13100",
		CityEn:    "tokyo23-ku",
		Year:      "2022",
		Format:    "3dtiles",
		WardCode:  "13101",
		WardEn:    "chiyoda-ku",
		Feature:   "bldg",
		LOD:       "2",
		NoTexture: true,
		Op:        "1_1_op",
		Ext:       ".zip",
	}, AssetNameFrom("https://example.com/12345/13100_tokyo23-ku_2022_3dtiles%20_1_1_op_bldg_13101_chiyoda-ku_lod2_no_texture.zip"))

	assert.Equal(t, AssetName{
		CityCode: "22325",
		CityEn:   "kannami-cho",
		Year:     "2022",
		Format:   "3dtiles",
		Op:       "op",
		Feature:  "bldg",
		Ex:       "",
		LOD:      "1",
		Ext:      ".zip",
	}, AssetNameFrom("22325_kannami-cho_2022_3dtiles_op_bldg_lod1.zip"))

	assert.Equal(t, AssetName{
		CityCode: "22221",
		CityEn:   "kosai-shi",
		Year:     "2022",
		Format:   "3dtiles",
		Op:       "op",
		NoDEM:    true,
		Feature:  "bldg",
		LOD:      "1",
		Ext:      ".zip",
	}, AssetNameFrom("22221_kosai-shi_2022_3dtiles_op_nodem_bldg_lod1.zip"))

	assert.Equal(t, AssetName{
		CityCode:    "13229",
		CityEn:      "nishitokyo-shi",
		Year:        "2022",
		Format:      "3dtiles",
		Op:          "1_op2",
		Feature:     "fld",
		Ext:         ".zip",
		FldCategory: "pref",
		FldName:     "shakujiigawa-shirakogawa_op",
	}, AssetNameFrom("13229_nishitokyo-shi_2022_3dtiles_1_op2_fld_pref_shakujiigawa-shirakogawa_op.zip"))

	assert.Equal(t, AssetName{
		CityCode:  "13100",
		CityEn:    "tokyo23-ku",
		Year:      "2020",
		Format:    "3dtiles",
		Op:        "4_2_op",
		Feature:   "bldg",
		Ext:       ".zip",
		WardCode:  "13109",
		WardEn:    "shinagawa-ku",
		LOD:       "2",
		NoTexture: true,
	}, AssetNameFrom("13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13109_shinagawa-ku_lod2_no_texture.zip"))

	assert.Equal(t, AssetName{
		CityCode:       "23212",
		CityEn:         "anjo-shi",
		Year:           "2020",
		Format:         "mvt",
		Op:             "4_op",
		Feature:        "urf",
		Ex:             "",
		Ext:            ".zip",
		UrfFeatureType: "UseDistrict",
	}, AssetNameFrom("23212_anjo-shi_2020_mvt_4_op_urf_UseDistrict.zip"))

	assert.Equal(t, AssetName{
		CityCode:    "13100",
		CityEn:      "tokyo23-ku",
		Year:        "2020",
		Format:      "3dtiles",
		Op:          "4_2_op",
		Feature:     "fld",
		Ex:          "",
		Ext:         ".zip",
		FldCategory: "natl",
		FldName:     "tmagawa_tamagawa-asakawa-etc_op",
	}, AssetNameFrom("13100_tokyo23-ku_2020_3dtiles_4_2_op_fld_natl_tmagawa_tamagawa-asakawa-etc_op.zip"))

	assert.Equal(t, AssetName{
		CityCode: "14130",
		CityEn:   "kawasaki-shi",
		Year:     "2022",
		Format:   "3dtiles",
		Op:       "1_op",
		Feature:  "frn",
		Ext:      ".zip",
	}, AssetNameFrom("14130_kawasaki-shi_2022_3dtiles_1_op_frn.zip"))

	assert.Equal(t, AssetName{
		CityCode: "43204",
		CityEn:   "arao-shi",
		Year:     "2020",
		Format:   "mvt",
		Op:       "5_op",
		Feature:  "tran",
		LOD:      "3",
		Ext:      ".zip",
	}, AssetNameFrom("https://example.com/43204_arao-shi_2020_mvt_5_op_tran_lod3.zip"))

	assert.Equal(t, AssetName{
		CityCode: "12210",
		CityEn:   "mobara-shi",
		Year:     "2022",
		Format:   "3dtiles",
		Op:       "1_op",
		Feature:  "tnm",
		FldName:  "12_1",
		Ext:      ".zip",
	}, AssetNameFrom("12210_mobara-shi_2022_3dtiles_1_op_tnm_12_1.zip"))

	assert.Equal(t, AssetName{
		CityCode: "14100",
		CityEn:   "yokohama-shi",
		Year:     "2022",
		Format:   "3dtiles",
		Op:       "1_op",
		Feature:  "brid",
		Ext:      ".zip",
	}, AssetNameFrom("14100_yokohama-shi_2022_3dtiles_1_op_brid.zip"))

	assert.Equal(t, AssetName{
		CityCode: "07212",
		CityEn:   "minamisouma-shi",
		Year:     "2022",
		Format:   "mvt",
		Op:       "1_op",
		Feature:  "rail",
		LOD:      "1",
		Ext:      ".zip",
	}, AssetNameFrom("07212_minamisouma-shi_2022_mvt_1_op_rail_lod1.zip"))

	assert.Equal(t, AssetName{
		CityCode: "20217",
		CityEn:   "saku-shi",
		Year:     "2022",
		Format:   "mvt",
		Op:       "1_op",
		Feature:  "gen",
		GenName:  "development_guidance_area",
		Ext:      ".zip",
	}, AssetNameFrom("20217_saku-shi_2022_mvt_1_op_gen_development_guidance_area.zip"))

	assert.Equal(t, AssetName{
		CityCode: "20217",
		CityEn:   "saku-shi",
		Year:     "2022",
		Format:   "mvt",
		Op:       "1_op",
		Feature:  "gen",
		GenName:  "development_guidance_area",
		LOD:      "1",
		Ext:      ".zip",
	}, AssetNameFrom("20217_saku-shi_2022_mvt_1_op_gen_development_guidance_area_lod1.zip"))

	assert.Equal(t, AssetName{
		CityCode: "22210",
		CityEn:   "fuji-shi",
		Year:     "2022",
		Format:   "3dtiles",
		Op:       "op",
		Feature:  "htd",
		FldName:  "22_1",
		Ext:      ".zip",
	}, AssetNameFrom("22210_fuji-shi_2022_3dtiles_op_htd_22_1.zip"))
}

func TestAssetName_String(t *testing.T) {
	assert.Equal(t, "13229_nishitokyo-shi_2022_citygml_1_op.zip", AssetName{
		CityCode: "13229",
		CityEn:   "nishitokyo-shi",
		Year:     "2022",
		Format:   "citygml",
		Op:       "1_op",
		Ext:      ".zip",
	}.String())

	assert.Equal(t, "13229_nishitokyo-shi_2022_3dtiles_1_op2_fld_pref_shakujiigawa-shirakogawa_op.zip", AssetName{
		CityCode:    "13229",
		CityEn:      "nishitokyo-shi",
		Year:        "2022",
		Format:      "3dtiles",
		Op:          "1_op2",
		Feature:     "fld",
		FldCategory: "pref",
		FldName:     "shakujiigawa-shirakogawa_op",
		Ext:         ".zip",
	}.String())

	assert.Equal(t, "13100_tokyo23-ku_2020_3dtiles_4_2_op_bldg_13109_shinagawa-ku_lod2_no_texture.zip", AssetName{
		CityCode:  "13100",
		CityEn:    "tokyo23-ku",
		Year:      "2020",
		Format:    "3dtiles",
		Op:        "4_2_op",
		Feature:   "bldg",
		Ext:       ".zip",
		WardCode:  "13109",
		WardEn:    "shinagawa-ku",
		LOD:       "2",
		NoTexture: true,
	}.String())

	assert.Equal(t, "23212_anjo-shi_2020_mvt_4_op_urf_UseDistrict.zip", AssetName{
		CityCode:       "23212",
		CityEn:         "anjo-shi",
		Year:           "2020",
		Format:         "mvt",
		Op:             "4_op",
		Feature:        "urf",
		Ext:            ".zip",
		UrfFeatureType: "UseDistrict",
	}.String())

	assert.Equal(t, "14130_kawasaki-shi_2022_3dtiles_1_op_frn.zip", AssetName{
		CityCode: "14130",
		CityEn:   "kawasaki-shi",
		Year:     "2022",
		Format:   "3dtiles",
		Op:       "1_op",
		Feature:  "frn",
		Ext:      ".zip",
	}.String())
}
