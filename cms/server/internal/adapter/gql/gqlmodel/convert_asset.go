package gqlmodel

import (
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
)

func ToAsset(a *asset.Asset, urlResolver func(a *asset.Asset) string) *Asset {
	if a == nil {
		return nil
	}

	var url string
	if urlResolver != nil {
		url = urlResolver(a)
	}

	var createdBy ID
	var createdByType OperatorType
	if a.User() != nil {
		createdBy = IDFrom(*a.User())
		createdByType = OperatorTypeUser
	}
	if a.Integration() != nil {
		createdBy = IDFrom(*a.Integration())
		createdByType = OperatorTypeIntegration
	}

	return &Asset{
		ID:                      IDFrom(a.ID()),
		ProjectID:               IDFrom(a.Project()),
		CreatedAt:               a.CreatedAt(),
		CreatedByID:             createdBy,
		CreatedByType:           createdByType,
		PreviewType:             ToPreviewType(a.PreviewType()),
		UUID:                    a.UUID(),
		URL:                     url,
		ThreadID:                IDFrom(a.Thread()),
		ArchiveExtractionStatus: ToArchiveExtractionStatus(a.ArchiveExtractionStatus()),
		Size:                    int64(a.Size()),
	}
}

func FromPreviewType(p *PreviewType) *asset.PreviewType {
	if p == nil {
		return nil
	}

	var p2 asset.PreviewType
	switch *p {
	case PreviewTypeImage:
		p2 = asset.PreviewTypeImage
	case PreviewTypeImageSVG:
		p2 = asset.PreviewTypeImageSvg
	case PreviewTypeGeo:
		p2 = asset.PreviewTypeGeo
	case PreviewTypeGeo3dTiles:
		p2 = asset.PreviewTypeGeo3dTiles
	case PreviewTypeGeoMvt:
		p2 = asset.PreviewTypeGeoMvt
	case PreviewTypeModel3d:
		p2 = asset.PreviewTypeModel3d
	case PreviewTypeUnknown:
		p2 = asset.PreviewTypeUnknown
	default:
		return nil
	}

	return &p2
}

func ToPreviewType(p *asset.PreviewType) *PreviewType {
	if p == nil {
		return nil
	}

	var p2 PreviewType
	switch *p {
	case asset.PreviewTypeImage:
		p2 = PreviewTypeImage
	case asset.PreviewTypeImageSvg:
		p2 = PreviewTypeImageSVG
	case asset.PreviewTypeGeo:
		p2 = PreviewTypeGeo
	case asset.PreviewTypeGeo3dTiles:
		p2 = PreviewTypeGeo3dTiles
	case asset.PreviewTypeGeoMvt:
		p2 = PreviewTypeGeoMvt
	case asset.PreviewTypeModel3d:
		p2 = PreviewTypeModel3d
	case asset.PreviewTypeUnknown:
		p2 = PreviewTypeUnknown
	default:
		return nil
	}

	return &p2
}

func ToArchiveExtractionStatus(s *asset.ArchiveExtractionStatus) *ArchiveExtractionStatus {
	if s == nil {
		return nil
	}

	var s2 ArchiveExtractionStatus
	switch *s {
	case asset.ArchiveExtractionStatusSkipped:
		s2 = ArchiveExtractionStatusSkipped
	case asset.ArchiveExtractionStatusPending:
		s2 = ArchiveExtractionStatusPending
	case asset.ArchiveExtractionStatusInProgress:
		s2 = ArchiveExtractionStatusInProgress
	case asset.ArchiveExtractionStatusDone:
		s2 = ArchiveExtractionStatusDone
	case asset.ArchiveExtractionStatusFailed:
		s2 = ArchiveExtractionStatusFailed
	default:
		return nil
	}

	return &s2
}

func ToAssetFile(a *asset.File) *AssetFile {
	if a == nil {
		return nil
	}

	return &AssetFile{
		Name:        a.Name(),
		Size:        int64(a.Size()),
		ContentType: lo.ToPtr(a.ContentType()),
		Path:        a.Path(),
		Children:    lo.Map(a.Children(), func(c *asset.File, _ int) *AssetFile { return ToAssetFile(c) }),
	}
}

func (s *AssetSort) Into() *usecasex.Sort {
	if s == nil {
		return nil
	}
	key := ""
	switch s.SortBy {
	case AssetSortTypeDate:
		key = "createdat"
	case AssetSortTypeName:
		key = "filename"
	case AssetSortTypeSize:
		key = "size"
	}
	if key == "" {
		return nil
	}
	return &usecasex.Sort{
		Key:      key,
		Reverted: s.Direction != nil && *s.Direction == SortDirectionDesc,
	}
}
