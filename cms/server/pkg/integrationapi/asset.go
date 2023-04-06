package integrationapi

import (
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/samber/lo"
)

func NewAsset(a *asset.Asset, f *asset.File, url string, all bool) *Asset {
	if a == nil {
		return nil
	}

	var ct, n *string
	if fct := f.ContentType(); fct != "" {
		ct = lo.ToPtr(fct)
	}
	if fn := f.Name(); fn != "" {
		n = lo.ToPtr(fn)
	}

	return &Asset{
		Id:                      a.ID(),
		ContentType:             ct,
		CreatedAt:               a.CreatedAt(),
		Name:                    n,
		PreviewType:             ToPreviewType(a.PreviewType()),
		ProjectId:               a.Project(),
		TotalSize:               lo.ToPtr(float32(a.Size())),
		Url:                     url,
		File:                    ToAssetFile(f, all),
		ArchiveExtractionStatus: ToAssetArchiveExtractionStatus(a.ArchiveExtractionStatus()),
	}
}

func ToAssetArchiveExtractionStatus(s *asset.ArchiveExtractionStatus) *AssetArchiveExtractionStatus {
	if s == nil {
		return nil
	}
	ss := ""
	switch *s {
	case asset.ArchiveExtractionStatusDone:
		ss = "done"
	case asset.ArchiveExtractionStatusFailed:
		ss = "failed"
	case asset.ArchiveExtractionStatusInProgress:
		ss = "in_progress"
	case asset.ArchiveExtractionStatusPending:
		ss = "pending"
	default:
		return nil
	}
	return lo.ToPtr(AssetArchiveExtractionStatus(ss))
}

func ToAssetFile(f *asset.File, all bool) *File {
	if f == nil {
		return nil
	}

	var children *[]File
	if all {
		children = lo.ToPtr(lo.FilterMap(f.Children(), func(c *asset.File, _ int) (File, bool) {
			f := ToAssetFile(c, true)
			if f == nil {
				return File{}, false
			}
			return *f, true
		}))
		if *children != nil && len(*children) == 0 {
			children = nil
		}
	}

	return &File{
		Name:        lo.ToPtr(f.Name()),
		ContentType: lo.ToPtr(f.ContentType()),
		Size:        lo.ToPtr(float32(f.Size())),
		Path:        lo.ToPtr(f.Path()),
		Children:    children,
	}
}

func ToPreviewType(pt *asset.PreviewType) *AssetPreviewType {
	if pt == nil {
		return lo.ToPtr(Unknown)
	}
	switch *pt {
	case asset.PreviewTypeGeo:
		return lo.ToPtr(Geo)
	case asset.PreviewTypeGeo3dTiles:
		return lo.ToPtr(Geo3dTiles)
	case asset.PreviewTypeGeoMvt:
		return lo.ToPtr(GeoMvt)
	case asset.PreviewTypeModel3d:
		return lo.ToPtr(Model3d)
	case asset.PreviewTypeImage:
		return lo.ToPtr(Image)
	case asset.PreviewTypeImageSvg:
		return lo.ToPtr(ImageSvg)
	case asset.PreviewTypeUnknown:
		return lo.ToPtr(Unknown)
	default:
		return lo.ToPtr(Unknown)
	}
}
