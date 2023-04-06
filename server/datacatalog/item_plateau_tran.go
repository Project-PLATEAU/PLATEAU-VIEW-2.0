package datacatalog

import (
	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/samber/lo"
)

func (i PlateauItem) TranItem(c PlateauIntermediateItem) *DataCatalogItem {
	set := TranSetFrom(i.Tran)
	if set.MaxLOD == nil {
		return nil
	}

	item := c.DataCatalogItem(
		"道路モデル",
		AssetNameFrom(set.MaxLOD.URL),
		set.MaxLOD.URL,
		i.DescriptionTran,
		tranLayers(set.MaxLODN),
		false,
		"",
	)

	item.Config = set.Config()

	return item
}

type TranSet struct {
	MaxLODN int
	MaxLOD  *cms.PublicAsset
	LOD0    *cms.PublicAsset
	LOD1    *cms.PublicAsset
	LOD2    *cms.PublicAsset
	LOD3    *cms.PublicAsset
}

func TranSetFrom(a []*cms.PublicAsset) TranSet {
	if len(a) == 0 {
		return TranSet{}
	}

	lods, maxLOD := assetWithLODFromList(a)
	if maxLOD == 0 {
		return TranSet{}
	}

	return TranSet{
		MaxLODN: maxLOD,
		MaxLOD:  tranSetLODFrom(lods, maxLOD),
		LOD0:    tranSetLODFrom(lods, 0),
		LOD1:    tranSetLODFrom(lods, 1),
		LOD2:    tranSetLODFrom(lods, 2),
		LOD3:    tranSetLODFrom(lods, 3),
	}
}

func tranSetLODFrom(assets []assetWithLOD, lod int) *cms.PublicAsset {
	tex, _ := lo.Find(assets, func(a assetWithLOD) bool {
		return a.LOD == lod
	})
	return tex.A
}

func (s TranSet) Config() (c DataCatalogItemConfig) {
	if s.LOD0 != nil {
		// mvt
		c.Data = append(c.Data, DataCatalogItemConfigItem{
			Name:   "LOD0",
			URL:    assetURLFromFormat(s.LOD0.URL, "mvt"),
			Type:   "mvt",
			Layers: tranLayers(0),
		})
	}

	if s.LOD1 != nil {
		// mvt
		c.Data = append(c.Data, DataCatalogItemConfigItem{
			Name:   "LOD1",
			URL:    assetURLFromFormat(s.LOD1.URL, "mvt"),
			Type:   "mvt",
			Layers: tranLayers(1),
		})
	}

	if s.LOD2 != nil {
		// mvt
		c.Data = append(c.Data, DataCatalogItemConfigItem{
			Name:   "LOD2",
			URL:    assetURLFromFormat(s.LOD2.URL, "mvt"),
			Type:   "mvt",
			Layers: tranLayers(2),
		})
	}

	if s.LOD3 != nil {
		// 3dtiles
		c.Data = append(c.Data, DataCatalogItemConfigItem{
			Name:   "LOD3",
			URL:    assetURLFromFormat(s.LOD3.URL, "3dtiles"),
			Type:   "3dtiles",
			Layers: tranLayers(3),
		})
	}

	return
}

func tranLayers(lod int) []string {
	if lod <= 1 {
		return []string{"Road"}
	}

	if lod == 2 {
		return []string{"TrafficArea", "AuxiliaryTrafficArea"}
	}

	return nil
}
