package geospatialjp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestCatalogFile(t *testing.T) {
	f, err := os.Open("testdata/xxxxx_xxx_catalog.xlsx")
	assert.NoError(t, err)

	thumbnail, err := os.ReadFile("testdata/test.jpg")
	assert.NoError(t, err)

	defer f.Close()
	xf, err := excelize.OpenReader(f)
	assert.NoError(t, err)

	c := NewCatalogFile(xf)

	cc, err := c.Parse()
	assert.NoError(t, err)
	assert.Equal(t, &Catalog{
		Title:             "TITLE",
		URL:               "URL",
		Notes:             "DESC\nDesc",
		Tags:              []string{"A", "B", "C", "D"},
		License:           "LICENSE",
		Organization:      "ORGANIZATION",
		Public:            "パブリック",
		Source:            "https://example.com",
		Version:           "1",
		Author:            "AAA",
		AuthorEmail:       "example@example.com",
		Maintainer:        "BBB",
		MaintainerEmail:   "example2@example.com",
		Spatial:           "",
		Quality:           "Quality",
		Restriction:       "Constraints",
		RegisteredDate:    "",
		Charge:            "A",
		Emergency:         "B",
		Area:              "Xxx",
		Fee:               "PRICE",
		LicenseAgreement:  "LICENSE AGREEMENT",
		CustomFields:      nil,
		Thumbnail:         thumbnail,
		ThumbnailFileName: "image1.jpeg",
	}, cc)

	// delete sheet
	assert.NoError(t, c.MustDeleteSheet())
	buf, err := c.File().WriteToBuffer()
	assert.NoError(t, err)

	xf2, err := excelize.OpenReader(buf)
	assert.NoError(t, err)
	c2 := NewCatalogFile(xf2)

	sheet := c2.getSheet()
	assert.Empty(t, sheet)
}

func TestMinXPos(t *testing.T) {
	assert.Equal(t, "A2", minXPos([]string{"C10", "A2", "B1"}))
}
