package datacatalog

import (
	"fmt"
	"sort"
	"strings"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/samber/lo"
)

const bridModelName = "橋梁モデル"
const railModelName = "鉄道モデル"

func (i PlateauItem) BridItem(c PlateauIntermediateItem) *DataCatalogItem {
	if len(i.Brid) == 0 {
		return nil
	}

	data := lo.Map(i.Brid, func(a *cms.PublicAsset, j int) DataCatalogItemConfigItem {
		an := AssetNameFrom(a.URL)
		name := ""
		if an.LOD != "" {
			name = fmt.Sprintf("LOD%s", an.LOD)
		} else if len(i.Brid) == 1 {
			name = bridModelName
		} else {
			name = fmt.Sprintf("%s%d", bridModelName, j+1)
		}

		return DataCatalogItemConfigItem{
			Name: name,
			URL:  assetURLFromFormat(a.URL, an.Format),
			Type: an.Format,
		}
	})
	sort.Slice(data, func(a, b int) bool {
		return strings.Compare(data[a].Name, data[b].Name) < 0
	})

	an := AssetNameFrom(i.Brid[0].URL)
	dci := c.DataCatalogItem(bridModelName, an, i.Brid[0].URL, i.DescriptionBrid, nil, false, "")
	if dci != nil {
		dci.Config = DataCatalogItemConfig{
			Data: data,
		}
	}

	return dci
}

func (i PlateauItem) RailItem(c PlateauIntermediateItem) *DataCatalogItem {
	if len(i.Rail) == 0 {
		return nil
	}

	data := lo.Map(i.Rail, func(a *cms.PublicAsset, j int) DataCatalogItemConfigItem {
		an := AssetNameFrom(a.URL)
		name := ""
		if an.LOD != "" {
			name = fmt.Sprintf("LOD%s", an.LOD)
		} else if len(i.Rail) == 1 {
			name = railModelName
		} else {
			name = fmt.Sprintf("%s%d", railModelName, j+1)
		}

		return DataCatalogItemConfigItem{
			Name:   name,
			URL:    assetURLFromFormat(a.URL, an.Format),
			Type:   an.Format,
			Layers: []string{"rail"},
		}
	})
	sort.Slice(data, func(a, b int) bool {
		return strings.Compare(data[a].Name, data[b].Name) < 0
	})

	an := AssetNameFrom(i.Rail[0].URL)
	dci := c.DataCatalogItem(railModelName, an, i.Rail[0].URL, i.DescriptionRail, []string{"rail"}, false, "")
	if dci != nil {
		dci.Config = DataCatalogItemConfig{
			Data: data,
		}
	}

	return dci
}
