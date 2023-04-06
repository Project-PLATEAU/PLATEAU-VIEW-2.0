package datacatalog

import (
	"fmt"
	"sort"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/samber/lo"
)

const genModelName = "汎用都市オブジェクトモデル"

type gen struct {
	a  *cms.PublicAsset
	an AssetName
	i  int
}

func (i PlateauItem) GenItems(c PlateauIntermediateItem) []*DataCatalogItem {
	assets := i.Gen
	if len(assets) == 0 {
		return nil
	}

	gens := lo.Map(assets, func(a *cms.PublicAsset, i int) gen {
		an := AssetNameFrom(a.URL)
		return gen{
			a:  a,
			an: an,
			i:  i,
		}
	})

	genGroups := lo.GroupBy(gens, func(r gen) string {
		return r.an.GenName
	})

	type entry struct {
		i    int
		item *DataCatalogItem
	}

	entries := lo.MapToSlice(genGroups, func(key string, gens []gen) entry {
		if len(gens) == 0 {
			return entry{}
		}

		f := gens[0]
		var layers []string
		if f.an.Format == "mvt" {
			layers = append(layers, f.an.GenName)
		}

		name, desc := descFromAsset(f.a, i.DescriptionGen)
		if name == "" {
			name = f.an.GenName
		}

		dci := c.DataCatalogItem(genModelName, f.an, f.a.URL, desc, layers, false, name)

		if dci != nil {
			dci.Config = DataCatalogItemConfig{
				Data: lo.Map(gens, func(g gen, i int) DataCatalogItemConfigItem {
					var layers []string
					if g.an.Format == "mvt" {
						layers = append(layers, g.an.GenName)
					}

					dname := name
					if len(gens) > 1 {
						if g.an.LOD != "" {
							dname = fmt.Sprintf("LOD%s", g.an.LOD)
						} else {
							dname = fmt.Sprintf("%s %d", name, i+1)
						}
					}

					return DataCatalogItemConfigItem{
						Name:   dname,
						URL:    assetURLFromFormat(g.a.URL, g.an.Format),
						Type:   g.an.Format,
						Layers: layers,
					}
				}),
			}
		}

		return entry{i: f.i, item: dci}
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
