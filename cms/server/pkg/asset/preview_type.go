package asset

import (
	"strings"

	"github.com/samber/lo"
)

type PreviewType string

func (p PreviewType) Ref() *PreviewType {
	return &p
}

const (
	PreviewTypeImage      PreviewType = "image"
	PreviewTypeImageSvg   PreviewType = "image_svg"
	PreviewTypeGeo        PreviewType = "geo"
	PreviewTypeGeo3dTiles PreviewType = "geo_3d_tiles"
	PreviewTypeGeoMvt     PreviewType = "geo_mvt"
	PreviewTypeModel3d    PreviewType = "model_3d"
	PreviewTypeUnknown    PreviewType = "unknown"
)

func PreviewTypeFrom(p string) (PreviewType, bool) {
	pp := strings.ToLower(p)
	switch PreviewType(pp) {
	case PreviewTypeImage:
		return PreviewTypeImage, true
	case PreviewTypeImageSvg:
		return PreviewTypeImageSvg, true
	case PreviewTypeGeo:
		return PreviewTypeGeo, true
	case PreviewTypeGeo3dTiles:
		return PreviewTypeGeo3dTiles, true
	case PreviewTypeGeoMvt:
		return PreviewTypeGeoMvt, true
	case PreviewTypeModel3d:
		return PreviewTypeModel3d, true
	case PreviewTypeUnknown:
		return PreviewTypeUnknown, true
	default:
		return PreviewType(""), false
	}
}

func PreviewTypeFromRef(p *string) *PreviewType {
	if p == nil {
		return nil
	}

	pp, ok := PreviewTypeFrom(*p)
	if !ok {
		return nil
	}
	return &pp
}

func PreviewTypeFromContentType(c string) *PreviewType {
	if strings.HasPrefix(c, "image/") {
		if strings.HasPrefix(c, "image/svg") {
			return lo.ToPtr(PreviewTypeImageSvg)
		}
		return lo.ToPtr(PreviewTypeImage)
	}
	return lo.ToPtr(PreviewTypeUnknown)
}

func (p PreviewType) String() string {
	return string(p)
}

func (p *PreviewType) StringRef() *string {
	if p == nil {
		return nil
	}
	p2 := string(*p)
	return &p2
}
