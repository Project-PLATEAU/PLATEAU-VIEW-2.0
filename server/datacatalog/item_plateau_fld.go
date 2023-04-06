package datacatalog

import (
	"fmt"
	"sort"
	"strings"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/samber/lo"
)

const fldModelName = "洪水浸水想定区域モデル"

type river struct {
	a   *cms.PublicAsset
	an  AssetName
	dic *DicEntry
	i   int
}

func (i PlateauItem) FldItems(c PlateauIntermediateItem) []*DataCatalogItem {
	if len(i.Fld) == 0 {
		return nil
	}

	rivers := lo.Map(i.Fld, func(a *cms.PublicAsset, i int) river {
		an := AssetNameFrom(a.URL)
		return river{
			a:   a,
			an:  an,
			dic: c.Dic.Fld(an.FldName, an.FldCategory),
			i:   i,
		}
	})

	riverGroups := lo.GroupBy(rivers, func(r river) string {
		if r.dic == nil {
			return ""
		}
		key := fmt.Sprintf("%s_%s", r.dic.Description, r.dic.Admin)
		return key
	})

	type entry struct {
		i    int
		item *DataCatalogItem
	}

	entries := lo.MapToSlice(riverGroups, func(key string, rivers []river) entry {
		if len(rivers) == 0 {
			return entry{}
		}

		sortRivers(rivers)

		r := rivers[0]
		name, desc := descFromAsset(r.a, i.DescriptionFld)
		dci := c.DataCatalogItem(fldModelName, r.an, r.a.URL, desc, nil, false, name)

		if dci != nil {
			dci.Name = fldName(fldModelName, i.CityName, r.an.FldName, r.dic)
			dci.Config = DataCatalogItemConfig{
				Data: lo.Map(rivers, func(rr river, _ int) DataCatalogItemConfigItem {
					name := dci.Name
					if rr.dic != nil {
						name = rr.dic.Scale
					}

					return DataCatalogItemConfigItem{
						Name: name,
						URL:  assetURLFromFormat(rr.a.URL, rr.an.Format),
						Type: rr.an.Format,
					}
				}),
			}
		}

		return entry{i: r.i, item: dci}
	})

	sort.Slice(entries, func(a, b int) bool {
		return entries[a].i < entries[b].i
	})

	return lo.FilterMap(entries, func(e entry, _ int) (*DataCatalogItem, bool) {
		if e.item == nil {
			return nil, false
		}
		return e.item, true
	})
}

func fldName(t, cityName, raw string, e *DicEntry) string {
	if e == nil {
		return raw
	}
	return fmt.Sprintf("%s %s（%s管理区間）（%s）", t, e.Description, e.Admin, cityName)
}

func sortRivers(rivers []river) {
	sort.Slice(rivers, func(a, b int) bool {
		if rivers[a].dic == nil {
			return false
		}
		if rivers[b].dic == nil {
			return true
		}
		s1, s2 := rivers[a].dic.Scale, rivers[b].dic.Scale
		if s1 == keikakukibo && s2 == souteisaidaikibo {
			return true
		}
		if s1 == souteisaidaikibo && s2 == keikakukibo {
			return false
		}
		return strings.Compare(rivers[a].dic.Scale, rivers[b].dic.Scale) < 0
	})
}

const keikakukibo = "計画規模"
const souteisaidaikibo = "想定最大規模"
