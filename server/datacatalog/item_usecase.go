package datacatalog

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/eukarya-inc/jpareacode"
	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

const folder = "フォルダ"
const folderEn = "folder"

var usecaseTypes = map[string]string{
	"公園":     "park",
	"避難施設":   "shelter",
	"鉄道":     "railway",
	"鉄道駅":    "station",
	"行政界":    "border",
	"ランドマーク": "landmark",
	"緊急輸送道路": "emergency_route",
	"ユースケース": "usecase",
}

type UsecaseItem struct {
	ID          string           `json:"id,omitempty"`
	Name        string           `json:"name,omitempty"`
	Type        string           `json:"type,omitempty"`
	Prefecture  string           `json:"prefecture,omitempty"`
	CityName    string           `json:"city_name,omitempty"`
	WardName    string           `json:"ward_name,omitempty"`
	OpenDataURL string           `json:"opendata_url,omitempty"`
	Description string           `json:"description,omitempty"`
	Year        string           `json:"year,omitempty"`
	Data        *cms.PublicAsset `json:"data,omitempty"`
	DataFormat  string           `json:"data_format,omitempty"`
	DataURL     string           `json:"data_url,omitempty"`
	DataLayers  string           `json:"data_layer,omitempty"`
	Config      string           `json:"config,omitempty"`
	Order       *int             `json:"order,omitempty"`
}

var reReiwa = regexp.MustCompile(`令和([0-9]+?)年度`)

func (i UsecaseItem) DataCatalogs() []DataCatalogItem {
	pref, prefCodeInt := normalizePref(i.Prefecture)
	prefCode := jpareacode.FormatPrefectureCode(prefCodeInt)

	var city, ward string
	if i.WardName != "" {
		city = i.CityName
		ward = i.WardName
	} else {
		city, ward, _ = strings.Cut(i.CityName, "/")
	}

	cCode := cityCode("", city, prefCodeInt)
	wCode := cityCode("", ward, prefCodeInt)

	if i.DataFormat == folder {
		return []DataCatalogItem{{
			ID:          i.ID,
			Name:        i.Name,
			Type:        folder,
			TypeEn:      folderEn,
			Pref:        pref,
			PrefCode:    prefCode,
			City:        city,
			CityCode:    cCode,
			Ward:        ward,
			WardCode:    wCode,
			Description: i.Description,
		}}
	}

	var c any
	_ = json.Unmarshal([]byte(i.Config), &c)

	u := ""
	if i.Data != nil && i.Data.URL != "" {
		u = i.Data.URL
	}
	if u == "" {
		u = i.DataURL
	}

	f := formatTypeEn(i.DataFormat)

	y := 0
	if ym := reReiwa.FindStringSubmatch(i.Year); len(ym) > 1 {
		yy, _ := strconv.Atoi(ym[1])
		if yy > 0 {
			y = yy + 2018
		}
	}

	t := i.Type
	if t != "" && t != "ユースケース" && !strings.HasSuffix(t, "情報") {
		t += "情報"
	}

	var layers []string
	if i.DataLayers != "" {
		layers = lo.Filter(util.Map(strings.Split(i.DataLayers, ","), strings.TrimSpace), func(s string, _ int) bool { return s != "" })
	}

	name := i.Name
	if t != "" && t != "ユースケース" {
		cityOrWard := city
		if ward != "" {
			cityOrWard = ward
		}
		name = fmt.Sprintf("%s（%s）", t, cityOrWard)
	}

	return []DataCatalogItem{{
		ID:          i.ID,
		Name:        name,
		Type:        t,
		TypeEn:      usecaseTypes[i.Type],
		Pref:        pref,
		PrefCode:    prefCode,
		City:        city,
		CityCode:    cCode,
		Ward:        ward,
		WardCode:    wCode,
		Format:      f,
		URL:         assetURLFromFormat(u, f),
		Description: i.Description,
		Config:      c,
		Layers:      layers,
		Year:        y,
		OpenDataURL: i.OpenDataURL,
		Order:       i.Order,
	}}
}

func formatTypeEn(f string) string {
	if f == "3D Tiles" {
		return "3dtiles"
	}
	return strings.ToLower(f)
}
