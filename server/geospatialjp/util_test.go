package geospatialjp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatalogFinalFileName(t *testing.T) {
	// TODO
}

func TestResources(t *testing.T) {
	// TODO
}

func TestExtractCityName(t *testing.T) {
	code, cityName, year, err := extractCityName("https://example.com/12210_mobara-shi_2022_citygml_1_lsld.zip")
	assert.NoError(t, err)
	assert.Equal(t, "12210", code)
	assert.Equal(t, "mobara-shi", cityName)
	assert.Equal(t, 2022, year)

	code, cityName, year, err = extractCityName("30422_taiji-cho_2021_citygml_2_op.zip")
	assert.NoError(t, err)
	assert.Equal(t, "30422", code)
	assert.Equal(t, "taiji-cho", cityName)
	assert.Equal(t, 2021, year)

	code, cityName, year, err = extractCityName("aaa")
	assert.EqualError(t, err, "invalid file name")
	assert.Empty(t, code)
	assert.Empty(t, cityName)
	assert.Equal(t, 0, year)
}

func TestSuffixFromSpec(t *testing.T) {
	assert.Equal(t, "（v2）", suffixFromSpec(2.3))
	assert.Equal(t, "", suffixFromSpec(1))
	assert.Equal(t, "", suffixFromSpec(0))
}

func TestDatasetName(t *testing.T) {
	assert.Equal(t, "plateau-11111-aaaa-2022", datasetName("11111", "aaaa", 2022))
	assert.Equal(t, "plateau-tokyo23ku", datasetName("11111", "tokyo23ku", 2020))
	assert.Equal(t, "plateau-tokyo23ku", datasetName("11111", "tokyo23-ku", 2020))
	assert.Equal(t, "plateau-tokyo23ku", datasetName("11111", "tokyo-23ku", 2020))
	assert.Equal(t, "plateau-tokyo23ku-2021", datasetName("11111", "tokyo23ku", 2021))
	assert.Equal(t, "plateau-tokyo23ku-2022", datasetName("11111", "tokyo23ku", 2022))
}
