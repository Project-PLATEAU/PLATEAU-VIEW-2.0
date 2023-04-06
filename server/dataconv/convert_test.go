package dataconv

import (
	"encoding/json"
	"os"
	"testing"

	geojson "github.com/paulmach/go.geojson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const borderName = "11238_hasuda-shi_border"

const border = `{
	"type" : "FeatureCollection",
	"name" : "11238_hasuda-shi_border",
	"features" : [
		{
			"type" : "Feature",
			"geometry" : {
				"type" : "MultiLineString",
				"coordinates" : [
					[
						[ 139.6050448421, 36.0366165201 ],
						[ 139.6048803941, 36.0366387451 ],
						[ 139.604250394, 36.0367901312 ],
						[ 139.6050448421, 36.0366165201 ]
					]
				]
			},
			"properties" : {
				"prefecture_code" : "11",
				"prefecture_name" : "埼玉県",
				"city_code" : "11238",
				"city_name" : "蓮田市"
			}
		}
	]
}`

var expectedBorder = `[
	{
    "id": "document",
		"name": "11238_hasuda-shi_border",
    "version": "1.0"
  },
  {
    "id": "11238_hasuda-shi_border_1_1",
    "wall": {
      "material": {
        "image": {
					"image": "` + wallImageDataURL + `",
          "repeat": true,
          "transparent": true
        }
      },
      "positions": {
        "cartographicDegrees": [
					139.6050448421,
					36.0366165201,
					100,
					139.6048803941,
					36.0366387451,
					100,
					139.604250394,
					36.0367901312,
					100,
					139.6050448421,
					36.0366165201,
					100
				]
			}
		},
		"properties" : {
			"prefecture_code" : "11",
			"prefecture_name" : "埼玉県",
			"city_code" : "11238",
			"city_name" : "蓮田市"
		}
	}
]`

func TestConvertBorder(t *testing.T) {
	var fc *geojson.FeatureCollection
	assert.NoError(t, json.Unmarshal([]byte(border), &fc))

	res, err := ConvertBorder(fc, borderName)
	assert.NoError(t, err)

	var expectedBorderJSON any
	assert.NoError(t, json.Unmarshal([]byte(expectedBorder), &expectedBorderJSON))

	assert.Equal(t, expectedBorderJSON, res)
}

func TestGenerateLandmarkImage(t *testing.T) {
	image, err := GenerateLandmarkImage("日本カメラ博物館")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile("test.png", image, 0666))
}

func TestProcessProperties(t *testing.T) {
	var m map[string]any
	_ = json.Unmarshal([]byte(`{"名称":"a","高さ":null}`), &m)
	assert.Equal(t, map[string]any{"名称": "a"}, processProperties(m))
}
