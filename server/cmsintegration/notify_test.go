package cmsintegration

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestItemFromUploadResult(t *testing.T) {
	assert.Equal(t, Item{
		All:        "all",
		Dictionary: "dic",
		Bldg:       []string{"bldg"},
		Tran:       []string{"tran"},
		Frn:        []string{"frn"},
		Veg:        []string{"veg"},
		Luse:       []string{"luse"},
		Lsld:       []string{"lsld"},
		Urf:        []string{"urf"},
		Fld:        []string{"fld"},
		Htd:        []string{"htd"},
		Tnm:        []string{"tnm"},
		Ifld:       []string{"ifld"},
	}, itemFromUploadResult(map[string][]string{
		"all":        {"all"},
		"dictionary": {"dic"},
		"bldg":       {"bldg"},
		"tran":       {"tran"},
		"frn":        {"frn"},
		"veg":        {"veg"},
		"luse":       {"luse"},
		"lsld":       {"lsld"},
		"urf":        {"urf"},
		"fld":        {"fld"},
		"htd":        {"htd"},
		"tnm":        {"tnm"},
		"ifld":       {"ifld"},
	}))
}

func TestReadDic(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder("GET", "https://example.com", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"hello": []any{"world"},
	})))

	d, err := readDic(context.Background(), "https://example.com")
	assert.NoError(t, err)
	assert.Equal(t, `{"hello":["world"]}`, d)
}
