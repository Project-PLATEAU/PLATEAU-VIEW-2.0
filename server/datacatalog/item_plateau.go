package datacatalog

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/eukarya-inc/jpareacode"
	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/spkg/bom"
)

//go:embed urf.csv
var urfFeatureTypesData []byte
var urfFeatureTypes map[string]string

func init() {
	r := csv.NewReader(bom.NewReader(bytes.NewReader(urfFeatureTypesData)))
	d := lo.Must(r.ReadAll())
	urfFeatureTypes = lo.SliceToMap(d[1:], func(c []string) (string, string) {
		return c[0], c[1]
	})
}

type PlateauItem struct {
	ID              string             `json:"id"`
	Prefecture      string             `json:"prefecture"`
	CityName        string             `json:"city_name"`
	Specification   string             `json:"specification"`
	CityGML         *cms.PublicAsset   `json:"citygml"`
	DescriptionBldg string             `json:"description_bldg"`
	DescriptionTran string             `json:"description_tran"`
	DescriptionFrn  string             `json:"description_frn"`
	DescriptionVeg  string             `json:"description_veg"`
	DescriptionLuse string             `json:"description_luse"`
	DescriptionLsld string             `json:"description_lsld"`
	DescriptionUrf  []string           `json:"description_urf"`
	DescriptionFld  []string           `json:"description_fld"`
	DescriptionHtd  []string           `json:"description_htd"`
	DescriptionIfld []string           `json:"description_ifld"`
	DescriptionTnm  []string           `json:"description_tnm"`
	DescriptionBrid string             `json:"description_brid"`
	DescriptionRail string             `json:"description_rail"`
	DescriptionGen  []string           `json:"description_gen"`
	Bldg            []*cms.PublicAsset `json:"bldg"`
	Tran            []*cms.PublicAsset `json:"tran"`
	Frn             []*cms.PublicAsset `json:"frn"`
	Veg             []*cms.PublicAsset `json:"veg"`
	Luse            []*cms.PublicAsset `json:"luse"`
	Lsld            []*cms.PublicAsset `json:"lsld"`
	Urf             []*cms.PublicAsset `json:"urf"`
	Fld             []*cms.PublicAsset `json:"fld"`
	Htd             []*cms.PublicAsset `json:"htd"`
	Ifld            []*cms.PublicAsset `json:"ifld"`
	Tnm             []*cms.PublicAsset `json:"tnm"`
	Brid            []*cms.PublicAsset `json:"brid"`
	Rail            []*cms.PublicAsset `json:"rail"`
	Gen             []*cms.PublicAsset `json:"gen"`
	Dictionary      *cms.PublicAsset   `json:"dictionary"`
	Dic             string             `json:"dic"`
	SearchIndex     []*cms.PublicAsset `json:"search_index"`
	OpenDataURL     string             `json:"opendata_url"`
}

func (i PlateauItem) FrnItem(c PlateauIntermediateItem) *DataCatalogItem {
	if len(i.Frn) == 0 {
		return nil
	}

	a := i.Frn[0]
	return c.DataCatalogItem("都市設備モデル", AssetNameFrom(a.URL), a.URL, i.DescriptionFrn, nil, false, "")
}

func (i PlateauItem) VegItem(c PlateauIntermediateItem) *DataCatalogItem {
	if len(i.Veg) == 0 {
		return nil
	}

	a := i.Veg[0]
	return c.DataCatalogItem("植生モデル", AssetNameFrom(a.URL), a.URL, i.DescriptionVeg, nil, false, "")
}

func (i PlateauItem) LuseItem(c PlateauIntermediateItem) *DataCatalogItem {
	if i.Luse == nil {
		return nil
	}

	a := i.Luse[0]
	return c.DataCatalogItem("土地利用モデル", AssetNameFrom(a.URL), a.URL, i.DescriptionLuse, []string{"luse"}, false, "")
}

func (i PlateauItem) LsldItem(c PlateauIntermediateItem) *DataCatalogItem {
	if i.Lsld == nil {
		return nil
	}

	a := i.Lsld[0]
	return c.DataCatalogItem("土砂災害警戒区域モデル", AssetNameFrom(a.URL), a.URL, i.DescriptionLsld, []string{"lsld"}, false, "")
}

func (i PlateauItem) UrfItems(c PlateauIntermediateItem) []*DataCatalogItem {
	if len(i.Urf) == 0 {
		return nil
	}

	return lo.Map(i.Urf, func(a *cms.PublicAsset, _ int) *DataCatalogItem {
		an := AssetNameFrom(a.URL)

		name, desc := descFromAsset(a, i.DescriptionUrf)
		return c.DataCatalogItem("都市計画決定情報モデル", an, a.URL, desc, urfLayers(an.UrfFeatureType), false, name)
	})
}

func (i PlateauItem) HtdItems(c PlateauIntermediateItem) []*DataCatalogItem {
	if len(i.Htd) == 0 {
		return nil
	}

	return lo.Map(i.Htd, func(a *cms.PublicAsset, _ int) *DataCatalogItem {
		an := AssetNameFrom(a.URL)

		name, desc := descFromAsset(a, i.DescriptionHtd)
		dci := c.DataCatalogItem("高潮浸水想定区域モデル", an, a.URL, desc, nil, false, name)

		if dci != nil {
			dci.Name = htdTnmIfldName(name, i.CityName, an.FldName, c.Dic.Htd(an.FldName))
		}
		return dci
	})
}

func (i PlateauItem) IfldItems(c PlateauIntermediateItem) []*DataCatalogItem {
	if len(i.Ifld) == 0 {
		return nil
	}

	return lo.Map(i.Ifld, func(a *cms.PublicAsset, _ int) *DataCatalogItem {
		an := AssetNameFrom(a.URL)

		name, desc := descFromAsset(a, i.DescriptionIfld)
		dci := c.DataCatalogItem("内水浸水想定区域モデル", an, a.URL, desc, nil, false, name)

		if dci != nil {
			dci.Name = htdTnmIfldName(name, i.CityName, an.FldName, c.Dic.Ifld(an.FldName))
		}
		return dci
	})
}

func (i PlateauItem) TnmItems(c PlateauIntermediateItem) []*DataCatalogItem {
	if len(i.Tnm) == 0 {
		return nil
	}

	return lo.Map(i.Tnm, func(a *cms.PublicAsset, _ int) *DataCatalogItem {
		an := AssetNameFrom(a.URL)

		name, desc := descFromAsset(a, i.DescriptionTnm)
		dci := c.DataCatalogItem("津波浸水想定区域モデル", an, a.URL, desc, nil, false, name)

		if dci != nil {
			dci.Name = htdTnmIfldName(name, i.CityName, an.FldName, c.Dic.Tnm(an.FldName))
		}
		return dci
	})
}

func (i PlateauItem) DataCatalogItems() []DataCatalogItem {
	c := i.IntermediateItem()
	if c.ID == "" {
		return nil
	}

	return util.DerefSlice(lo.Filter(
		append(append(append(append(append(append(append(append(append(
			i.BldgItems(c),
			i.TranItem(c),
			i.FrnItem(c),
			i.VegItem(c),
			i.LuseItem(c),
			i.LsldItem(c)),
			i.UrfItems(c)...),
			i.FldItems(c)...),
			i.TnmItems(c)...),
			i.HtdItems(c)...),
			i.IfldItems(c)...),
			i.BridItem(c)),
			i.RailItem(c)),
			i.GenItems(c)...,
		),
		func(i *DataCatalogItem, _ int) bool {
			return i != nil
		},
	))
}

func (i PlateauItem) IntermediateItem() PlateauIntermediateItem {
	au := ""
	if i.CityGML != nil {
		au = i.CityGML.URL
	} else if len(i.Bldg) > 0 {
		au = i.Bldg[0].URL
	}

	if au == "" {
		return PlateauIntermediateItem{}
	}

	an := AssetNameFrom(au)
	dic := Dic{}
	_ = json.Unmarshal(bom.Clean([]byte(i.Dic)), &dic)

	return PlateauIntermediateItem{
		ID:          i.ID,
		Prefecture:  i.Prefecture,
		City:        i.CityName,
		CityEn:      an.CityEn,
		CityCode:    an.CityCode,
		Dic:         dic,
		OpenDataURL: i.OpenDataURL,
	}
}

type PlateauIntermediateItem struct {
	ID          string
	Prefecture  string
	City        string
	CityEn      string
	CityCode    string
	Dic         Dic
	OpenDataURL string
}

func (i *PlateauIntermediateItem) DataCatalogItem(t string, an AssetName, assetURL, desc string, layers []string, firstWard bool, nameOverride string) *DataCatalogItem {
	if i == nil {
		return nil
	}

	id := i.id(an)
	if id == "" {
		return nil
	}

	wardName := i.Dic.WardName(an.WardCode)
	if wardName == "" && an.WardCode != "" {
		wardName = an.WardEn
	}

	cityOrWardName := i.City
	if wardName != "" {
		cityOrWardName = wardName
	}

	name, t2, t2en := itemName(t, cityOrWardName, nameOverride, an)
	y, _ := strconv.Atoi(an.Year)
	pref, prefCode := normalizePref(i.Prefecture)

	var itemID string
	if an.Feature == "bldg" && (an.WardCode == "" || firstWard) {
		itemID = i.ID
	}

	return &DataCatalogItem{
		ID:          id,
		ItemID:      itemID,
		Type:        t,
		TypeEn:      an.Feature,
		Type2:       t2,
		Type2En:     t2en,
		Name:        name,
		Pref:        pref,
		PrefCode:    jpareacode.FormatPrefectureCode(prefCode),
		City:        i.City,
		CityEn:      i.CityEn,
		CityCode:    cityCode(i.CityCode, i.City, prefCode),
		Ward:        wardName,
		WardEn:      an.WardEn,
		WardCode:    cityCode(an.WardCode, wardName, prefCode),
		Description: desc,
		URL:         assetURLFromFormat(assetURL, an.Format),
		Format:      an.Format,
		Year:        y,
		Layers:      layers,
		OpenDataURL: i.OpenDataURL,
	}
}

func (i *PlateauIntermediateItem) id(an AssetName) string {
	return strings.Join(lo.Filter([]string{
		i.CityCode,
		i.CityEn,
		an.WardCode,
		an.WardEn,
		an.Feature,
		an.UrfFeatureType,
		an.FldNameAndCategory(),
		an.GenName,
	}, func(s string, _ int) bool { return s != "" }), "_")
}

func itemName(t, cityOrWardName, nameOverride string, an AssetName) (name, t2, t2en string) {
	if an.Feature == "urf" {
		t2 = an.UrfFeatureType
		t2en = an.UrfFeatureType

		if urfName := urfFeatureTypes[an.UrfFeatureType]; urfName != "" {
			t2 = urfName
			if nameOverride == "" {
				name = fmt.Sprintf("%sモデル", urfName)
			}
		} else {
			name = an.UrfFeatureType
		}
	}

	if name == "" {
		if nameOverride != "" {
			name = nameOverride
		} else {
			name = t
		}
	}

	name += fmt.Sprintf("（%s）", cityOrWardName)
	return
}

func assetsByWards(a []*cms.PublicAsset) map[string][]*cms.PublicAsset {
	if len(a) == 0 {
		return nil
	}

	r := map[string][]*cms.PublicAsset{}
	for _, a := range a {
		if a == nil {
			continue
		}

		an := AssetNameFrom(a.URL)
		k := an.WardCode
		if _, ok := r[k]; !ok {
			r[k] = []*cms.PublicAsset{a}
		} else {
			r[k] = append(r[k], a)
		}
	}
	return r
}

var reName = regexp.MustCompile(`^@name:\s*(.+)(?:$|\n)`)

func descFromAsset(a *cms.PublicAsset, descs []string) (string, string) {
	if a == nil || len(descs) == 0 {
		return "", ""
	}

	fn := strings.TrimSuffix(path.Base(a.URL), path.Ext(a.URL))
	for _, desc := range descs {
		b, a, ok := strings.Cut(desc, "\n")
		if ok && strings.Contains(b, fn) {
			return nameFromDescription(strings.TrimSpace(a))
		}
	}

	return "", ""
}

func nameFromDescription(d string) (string, string) {
	if m := reName.FindStringSubmatch(d); len(m) > 0 {
		name := m[1]
		_, n, _ := strings.Cut(d, "\n")
		return name, strings.TrimSpace(n)
	}

	return "", d
}

type DataCatalogItemConfig struct {
	Data []DataCatalogItemConfigItem `json:"data,omitempty"`
}

type DataCatalogItemConfigItem struct {
	Name   string   `json:"name"`
	URL    string   `json:"url"`
	Type   string   `json:"type"`
	Layers []string `json:"layer,omitempty"`
}

type assetWithLOD struct {
	A   *cms.PublicAsset
	F   AssetName
	LOD int
}

func assetWithLODFromList(a []*cms.PublicAsset) ([]assetWithLOD, int) {
	maxLOD := 0
	return lo.FilterMap(a, func(a *cms.PublicAsset, _ int) (assetWithLOD, bool) {
		l := assetWithLODFrom(a)
		if l != nil && maxLOD < l.LOD {
			maxLOD = l.LOD
		}
		return *l, l != nil
	}), maxLOD
}

func assetWithLODFrom(a *cms.PublicAsset) *assetWithLOD {
	if a == nil {
		return nil
	}
	f := AssetNameFrom(a.URL)
	l, _ := strconv.Atoi(f.LOD)
	return &assetWithLOD{A: a, LOD: l, F: f}
}

func htdTnmIfldName(t, cityName, raw string, e *DicEntry) string {
	if e == nil {
		return raw
	}
	return fmt.Sprintf("%s %s（%s）", t, e.Description, cityName)
}

func urfLayers(ty string) []string {
	if ty == "WaterWay" {
		ty = "Waterway"
	}
	return []string{ty}
}
