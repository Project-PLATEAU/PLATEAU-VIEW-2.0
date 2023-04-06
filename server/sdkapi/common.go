package sdkapi

import (
	"fmt"
	"net/url"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/eukarya-inc/jpareacode"
	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/mitchellh/mapstructure"
	"github.com/reearth/reearthx/log"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

const modelKey = "plateau"
const tokyo = "東京都"

type Config struct {
	CMSBaseURL   string
	CMSToken     string
	Project      string
	Model        string
	Token        string
	DisableCache bool
	CacheTTL     int
}

func (c *Config) Default() {
	if c.Model == "" {
		c.Model = modelKey
	}
}

type DatasetResponse struct {
	Data []*DatasetPref `json:"data"`
}

type DatasetPref struct {
	ID    string        `json:"id"`
	Title string        `json:"title"`
	Data  []DatasetCity `json:"data"`
}

type DatasetCity struct {
	ID           string   `json:"id"`
	CityCode     int      `json:"-"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	FeatureTypes []string `json:"featureTypes"`
}

type FilesResponse map[string][]File

type File struct {
	Code   string  `json:"code"`
	URL    string  `json:"url"`
	MaxLOD float64 `json:"maxLod"`
}

type Items []Item

func (i Items) DatasetResponse() (r *DatasetResponse) {
	warning := []string{}
	r = &DatasetResponse{}
	prefs := []*DatasetPref{}
	prefm := map[string]*DatasetPref{}
	for _, i := range i {
		invalid := false
		if !i.IsPublic() {
			warning = append(warning, fmt.Sprintf("%s:not_published", i.CityName))
			invalid = true
		}

		if i.CityGML == nil || i.CityGML.ID == "" {
			warning = append(warning, fmt.Sprintf("%s:no_citygml", i.CityName))
			invalid = true
		}

		if i.CityGML != nil && !i.CityGML.IsExtractionDone() {
			warning = append(warning, fmt.Sprintf("%s:invalid_citygml", i.CityName))
			invalid = true
		}

		if i.MaxLOD == nil || i.MaxLOD.URL == "" {
			warning = append(warning, fmt.Sprintf("%s:no_maxlod", i.CityName))
			invalid = true
		}

		ft := i.FeatureTypes()
		if len(ft) == 0 {
			warning = append(warning, fmt.Sprintf("%s:no_features", i.CityName))
			invalid = true
		}

		if invalid {
			continue
		}

		if _, ok := prefm[i.Prefecture]; !ok {
			pd := &DatasetPref{
				ID:    i.Prefecture,
				Title: i.Prefecture,
			}
			prefs = append(prefs, pd)
			prefm[i.Prefecture] = prefs[len(prefs)-1]
		}

		d := DatasetCity{
			ID:           i.ID,
			CityCode:     i.CityCode(),
			Title:        i.CityName,
			Description:  i.Description,
			FeatureTypes: ft,
		}
		pd := prefm[i.Prefecture]
		pd.Data = append(pd.Data, d)
	}

	// sort
	sort.Slice(prefs, func(a, b int) bool {
		at, bt := prefs[a].Title, prefs[b].Title
		ac, bc := 0, 0
		if at != tokyo {
			ac = jpareacode.PrefectureCodeInt(at)
		}
		if bt != tokyo {
			bc = jpareacode.PrefectureCodeInt(bt)
		}
		return ac < bc
	})

	for _, p := range prefs {
		sort.Slice(p.Data, func(a, b int) bool {
			return p.Data[a].CityCode < p.Data[b].CityCode
		})
	}

	r.Data = prefs

	if len(warning) > 0 {
		log.Warnf("sdk: dataset warn: %s", strings.Join(warning, ", "))
	}

	return
}

type Item struct {
	ID             string            `json:"id"`
	Prefecture     string            `json:"prefecture"`
	CityName       string            `json:"city_name"`
	CityGML        *cms.PublicAsset  `json:"citygml"`
	Description    string            `json:"description_bldg"`
	MaxLOD         *cms.PublicAsset  `json:"max_lod"`
	Bldg           []cms.PublicAsset `json:"bldg"`
	Tran           []cms.PublicAsset `json:"tran"`
	Frn            []cms.PublicAsset `json:"frn"`
	Veg            []cms.PublicAsset `json:"veg"`
	SDKPublication string            `json:"sdk_publication"`
}

func (i Item) IsPublic() bool {
	return i.SDKPublication == "公開する"
}

func (i Item) CityCode() int {
	return cityCode(i.CityGML)
}

func (i Item) FeatureTypes() (t []string) {
	if len(i.Bldg) > 0 {
		t = append(t, "bldg")
	}
	if len(i.Tran) > 0 {
		t = append(t, "tran")
	}
	if len(i.Frn) > 0 {
		t = append(t, "frn")
	}
	if len(i.Veg) > 0 {
		t = append(t, "veg")
	}
	return
}

type MaxLODColumns []MaxLODColumn

type MaxLODColumn struct {
	Code   string  `json:"code"`
	Type   string  `json:"type"`
	MaxLOD float64 `json:"maxLod"`
}

type MaxLODMap map[string]map[string]float64

func (mc MaxLODColumns) Map() MaxLODMap {
	m := MaxLODMap{}

	for _, c := range mc {
		if _, ok := m[c.Type]; !ok {
			m[c.Type] = map[string]float64{}
		}
		t := m[c.Type]
		t[c.Code] = c.MaxLOD
	}

	return m
}

func (mm MaxLODMap) Files(urls []*url.URL) (r FilesResponse) {
	r = FilesResponse{}
	for ty, m := range mm {
		if _, ok := r[ty]; !ok {
			r[ty] = ([]File)(nil)
		}
		for code, maxlod := range m {
			prefix := fmt.Sprintf("%s_%s_", code, ty)
			u, ok := lo.Find(urls, func(u *url.URL) bool {
				return strings.HasPrefix(path.Base(u.Path), prefix) && path.Ext(u.Path) == ".gml"
			})
			if ok {
				r[ty] = append(r[ty], File{
					Code:   code,
					URL:    u.String(),
					MaxLOD: maxlod,
				})
			}
		}
		slices.SortFunc(r[ty], func(i, j File) bool {
			return i.Code < j.Code
		})
	}
	return
}

type IItem struct {
	ID             string `json:"id" cms:"id,text"`
	Prefecture     string `json:"prefecture" cms:"prefecture,text"`
	CityName       string `json:"city_name" cms:"city_name,text"`
	CityGML        any    `json:"citygml" cms:"citygml,asset"`
	Description    string `json:"description_bldg" cms:"description_bldg,textarea"`
	MaxLOD         any    `json:"max_lod" cms:"max_lod,asset"`
	Bldg           []any  `json:"bldg" cms:"bldg,asset"`
	Tran           []any  `json:"tran" cms:"tran,asset"`
	Frn            []any  `json:"frn" cms:"frn,asset"`
	Veg            []any  `json:"veg" cms:"veg,asset"`
	SDKPublication string `json:"sdk_publication" cms:"sdk_publication,select"`
}

func (i IItem) Item() Item {
	return Item{
		ID:             i.ID,
		Prefecture:     i.Prefecture,
		CityName:       i.CityName,
		CityGML:        integrationAssetToAsset(i.CityGML).ToPublic(),
		Description:    i.Description,
		MaxLOD:         integrationAssetToAsset(i.MaxLOD).ToPublic(),
		Bldg:           assetsToPublic(integrationAssetToAssets(i.Bldg)),
		Tran:           assetsToPublic(integrationAssetToAssets(i.Tran)),
		Frn:            assetsToPublic(integrationAssetToAssets(i.Frn)),
		Veg:            assetsToPublic(integrationAssetToAssets(i.Veg)),
		SDKPublication: i.SDKPublication,
	}
}

func cityCode(a *cms.PublicAsset) int {
	if a == nil || a.URL == "" {
		return 0
	}

	u, err := url.Parse(a.URL)
	if err != nil {
		return 0
	}

	code, _, ok := strings.Cut(path.Base(u.Path), "_")
	if !ok {
		return 0
	}

	c, _ := strconv.Atoi(code)
	return c
}

func ItemsFromIntegration(items []cms.Item) Items {
	return lo.FilterMap(items, func(i cms.Item, _ int) (Item, bool) {
		item := ItemFromIntegration(&i)
		return item, item.IsPublic()
	})
}

func ItemFromIntegration(ci *cms.Item) Item {
	i := IItem{}
	ci.Unmarshal(&i)
	return i.Item()
}

func assetsToPublic(a []cms.Asset) []cms.PublicAsset {
	return lo.FilterMap(a, func(a cms.Asset, _ int) (cms.PublicAsset, bool) {
		p := a.ToPublic()
		if p == nil {
			return cms.PublicAsset{}, false
		}
		return *p, true
	})
}

func integrationAssetToAssets(a []any) []cms.Asset {
	return lo.FilterMap(a, func(a any, _ int) (cms.Asset, bool) {
		pa := integrationAssetToAsset(a)
		if pa == nil {
			return cms.Asset{}, false
		}
		return *pa, true
	})
}

func integrationAssetToAsset(a any) *cms.Asset {
	if a == nil {
		return nil
	}

	m, ok := a.(map[string]any)
	if !ok {
		return nil
	}

	pa := &cms.Asset{}
	if err := mapstructure.Decode(m, pa); err != nil {
		return nil
	}
	return pa
}
