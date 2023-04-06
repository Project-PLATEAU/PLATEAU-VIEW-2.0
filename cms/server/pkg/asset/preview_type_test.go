package asset

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestPreviewType_PreviewTypeFrom(t *testing.T) {
	tests := []struct {
		Name     string
		Expected struct {
			TA   PreviewType
			Bool bool
		}
	}{
		{
			Name: "image",
			Expected: struct {
				TA   PreviewType
				Bool bool
			}{
				TA:   PreviewTypeImage,
				Bool: true,
			},
		},
		{
			Name: "IMAGE",
			Expected: struct {
				TA   PreviewType
				Bool bool
			}{
				TA:   PreviewTypeImage,
				Bool: true,
			},
		},
		{
			Name: "image_svg",
			Expected: struct {
				TA   PreviewType
				Bool bool
			}{
				TA:   PreviewTypeImageSvg,
				Bool: true,
			},
		},
		{
			Name: "geo",
			Expected: struct {
				TA   PreviewType
				Bool bool
			}{
				TA:   PreviewTypeGeo,
				Bool: true,
			},
		},
		{
			Name: "geo_3d_tiles",
			Expected: struct {
				TA   PreviewType
				Bool bool
			}{
				TA:   PreviewTypeGeo3dTiles,
				Bool: true,
			},
		},
		{
			Name: "geo_mvt",
			Expected: struct {
				TA   PreviewType
				Bool bool
			}{
				TA:   PreviewTypeGeoMvt,
				Bool: true,
			},
		},
		{
			Name: "model_3d",
			Expected: struct {
				TA   PreviewType
				Bool bool
			}{
				TA:   PreviewTypeModel3d,
				Bool: true,
			},
		},
		{
			Name: "unknown",
			Expected: struct {
				TA   PreviewType
				Bool bool
			}{
				TA:   PreviewTypeUnknown,
				Bool: true,
			},
		},
		{
			Name: "undefined",
			Expected: struct {
				TA   PreviewType
				Bool bool
			}{
				TA:   PreviewType(""),
				Bool: false,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			res, ok := PreviewTypeFrom(tc.Name)
			assert.Equal(t, tc.Expected.TA, res)
			assert.Equal(t, tc.Expected.Bool, ok)
		})
	}
}

func TestPreviewType_PreviewTypeFromRef(t *testing.T) {
	i := PreviewTypeImage
	is := PreviewTypeImageSvg
	g := PreviewTypeGeo
	g3d := PreviewTypeGeo3dTiles
	mvt := PreviewTypeGeoMvt
	m := PreviewTypeModel3d
	u := PreviewTypeUnknown

	tests := []struct {
		Name     string
		Input    *string
		Expected *PreviewType
	}{
		{
			Name:     "image",
			Input:    lo.ToPtr("image"),
			Expected: &i,
		},
		{
			Name:     "upper case image",
			Input:    lo.ToPtr("IMAGE"),
			Expected: &i,
		},
		{
			Name:     "image_svg",
			Input:    lo.ToPtr("image_svg"),
			Expected: &is,
		},
		{
			Name:     "geo",
			Input:    lo.ToPtr("geo"),
			Expected: &g,
		},
		{
			Name:     "geo_3d_tiles",
			Input:    lo.ToPtr("geo_3d_tiles"),
			Expected: &g3d,
		},
		{
			Name:     "geo_mvt",
			Input:    lo.ToPtr("geo_mvt"),
			Expected: &mvt,
		},
		{
			Name:     "model_3d",
			Input:    lo.ToPtr("model_3d"),
			Expected: &m,
		},
		{
			Name:     "unknown",
			Input:    lo.ToPtr("unknown"),
			Expected: &u,
		},
		{
			Name:  "undefined",
			Input: lo.ToPtr("undefined"),
		},
		{
			Name: "nil input",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			res := PreviewTypeFromRef(tc.Input)
			assert.Equal(t, tc.Expected, res)
		})
	}
}

func TestPreviewType_PreviewTypeFromContentType(t *testing.T) {
	c1 := "image/png"
	want1 := lo.ToPtr(PreviewTypeImage)
	got1 := PreviewTypeFromContentType(c1)
	assert.Equal(t, want1, got1)

	c2 := "video/mp4"
	want2 := lo.ToPtr(PreviewTypeUnknown)
	got2 := PreviewTypeFromContentType(c2)
	assert.Equal(t, want2, got2)

	c3 := "image/svg"
	want3 := lo.ToPtr(PreviewTypeImageSvg)
	got3 := PreviewTypeFromContentType(c3)
	assert.Equal(t, want3, got3)
}

func TestPreviewType_String(t *testing.T) {
	s := "image"
	pt := PreviewTypeImage
	assert.Equal(t, s, pt.String())
}

func TestPreviewType_StringRef(t *testing.T) {
	var pt1 *PreviewType
	var pt2 *PreviewType = lo.ToPtr(PreviewTypeImage)
	s := string(*pt2)

	tests := []struct {
		Name     string
		Input    *string
		Expected *string
	}{
		{
			Name:     "nil PreviewType pointer",
			Input:    pt1.StringRef(),
			Expected: nil,
		},
		{
			Name:     "PreviewType pointer",
			Input:    pt2.StringRef(),
			Expected: &s,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, tc.Input)
		})
	}
}
