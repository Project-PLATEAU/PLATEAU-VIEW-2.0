package gqlmodel

import (
	"testing"

	"github.com/google/uuid"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestToAsset(t *testing.T) {
	pid1 := id.NewProjectID()
	uid1 := id.NewUserID()
	id1 := id.NewAssetID()
	var pti asset.PreviewType = asset.PreviewTypeImage
	uuid := uuid.New().String()
	thid := id.NewThreadID()
	a1 := asset.New().ID(id1).Project(pid1).CreatedByUser(uid1).FileName("aaa.jpg").Size(1000).Type(&pti).UUID(uuid).Thread(thid).MustBuild()

	want1 := Asset{
		ID:            ID(id1.String()),
		ProjectID:     ID(pid1.String()),
		CreatedAt:     id1.Timestamp(),
		CreatedByID:   ID(uid1.String()),
		CreatedByType: OperatorTypeUser,
		PreviewType:   ToPreviewType(&pti),
		UUID:          uuid,
		URL:           "xxx",
		ThreadID:      ID(thid.String()),
		Size:          1000,
	}

	var a2 *asset.Asset = nil
	want2 := (*Asset)(nil)

	tests := []struct {
		name string
		arg  *asset.Asset
		want *Asset
	}{
		{
			name: "to asset valid",
			arg:  a1,
			want: &want1,
		},
		{
			name: "to asset nil",
			arg:  a2,
			want: want2,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			resolver := func(_ *asset.Asset) string {
				return "xxx"
			}
			got := ToAsset(tc.arg, resolver)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestConvertAsset_FromPreviewType(t *testing.T) {
	var pt1 PreviewType = PreviewTypeImage
	want1 := asset.PreviewTypeImage
	got1 := FromPreviewType(&pt1)
	assert.Equal(t, &want1, got1)

	var pt2 PreviewType = PreviewTypeGeo
	want2 := asset.PreviewTypeGeo
	got2 := FromPreviewType(&pt2)
	assert.Equal(t, &want2, got2)

	var pt3 PreviewType = PreviewTypeGeo3dTiles
	want3 := asset.PreviewTypeGeo3dTiles
	got3 := FromPreviewType(&pt3)
	assert.Equal(t, &want3, got3)

	var pt4 PreviewType = PreviewTypeGeoMvt
	want4 := asset.PreviewTypeGeoMvt
	got4 := FromPreviewType(&pt4)
	assert.Equal(t, &want4, got4)

	var pt5 PreviewType = PreviewTypeModel3d
	want5 := asset.PreviewTypeModel3d
	got5 := FromPreviewType(&pt5)
	assert.Equal(t, &want5, got5)

	var pt6 *PreviewType = nil
	want6 := (*asset.PreviewType)(nil)
	got6 := FromPreviewType(pt6)
	assert.Equal(t, want6, got6)

	var pt7 PreviewType = "test"
	want7 := (*asset.PreviewType)(nil)
	got7 := FromPreviewType(&pt7)
	assert.Equal(t, want7, got7)

	var pt8 PreviewType = PreviewTypeUnknown
	want8 := asset.PreviewTypeUnknown
	got8 := FromPreviewType(&pt8)
	assert.Equal(t, &want8, got8)

	var pt9 PreviewType = PreviewTypeImageSVG
	want9 := asset.PreviewTypeImageSvg
	got9 := FromPreviewType(&pt9)
	assert.Equal(t, &want9, got9)
}

func TestConvertAsset_ToPreviewType(t *testing.T) {
	var pt1 asset.PreviewType = asset.PreviewTypeImage
	want1 := PreviewTypeImage
	got1 := ToPreviewType(&pt1)
	assert.Equal(t, &want1, got1)

	var pt2 asset.PreviewType = asset.PreviewTypeGeo
	want2 := PreviewTypeGeo
	got2 := ToPreviewType(&pt2)
	assert.Equal(t, &want2, got2)

	var pt3 asset.PreviewType = asset.PreviewTypeGeo3dTiles
	want3 := PreviewTypeGeo3dTiles
	got3 := ToPreviewType(&pt3)
	assert.Equal(t, &want3, got3)

	var pt4 asset.PreviewType = asset.PreviewTypeModel3d
	want4 := PreviewTypeModel3d
	got4 := ToPreviewType(&pt4)
	assert.Equal(t, &want4, got4)

	var pt5 asset.PreviewType = asset.PreviewTypeModel3d
	want5 := PreviewTypeModel3d
	got5 := ToPreviewType(&pt5)
	assert.Equal(t, &want5, got5)

	var pt6 *asset.PreviewType = nil
	want6 := (*PreviewType)(nil)
	got6 := ToPreviewType(pt6)
	assert.Equal(t, want6, got6)

	var pt7 asset.PreviewType = "test"
	want7 := (*PreviewType)(nil)
	got7 := ToPreviewType(&pt7)
	assert.Equal(t, want7, got7)

	var pt8 asset.PreviewType = asset.PreviewTypeUnknown
	want8 := PreviewTypeUnknown
	got8 := ToPreviewType(&pt8)
	assert.Equal(t, &want8, got8)

	var pt9 asset.PreviewType = asset.PreviewTypeImageSvg
	want9 := PreviewTypeImageSVG
	got9 := ToPreviewType(&pt9)
	assert.Equal(t, &want9, got9)
}

func TestConvertAsset_ToStatus(t *testing.T) {
	s0 := asset.ArchiveExtractionStatusSkipped
	want0 := ArchiveExtractionStatusSkipped
	got0 := ToArchiveExtractionStatus(&s0)
	assert.Equal(t, &want0, got0)

	s1 := asset.ArchiveExtractionStatusPending
	want1 := ArchiveExtractionStatusPending
	got1 := ToArchiveExtractionStatus(&s1)
	assert.Equal(t, &want1, got1)

	s2 := asset.ArchiveExtractionStatusInProgress
	want2 := ArchiveExtractionStatusInProgress
	got2 := ToArchiveExtractionStatus(&s2)
	assert.Equal(t, &want2, got2)

	s3 := asset.ArchiveExtractionStatusDone
	want3 := ArchiveExtractionStatusDone
	got3 := ToArchiveExtractionStatus(&s3)
	assert.Equal(t, &want3, got3)

	s4 := asset.ArchiveExtractionStatusFailed
	want4 := ArchiveExtractionStatusFailed
	got4 := ToArchiveExtractionStatus(&s4)
	assert.Equal(t, &want4, got4)

	var s5 *asset.ArchiveExtractionStatus = nil
	want5 := (*ArchiveExtractionStatus)(nil)
	got5 := ToArchiveExtractionStatus(s5)
	assert.Equal(t, want5, got5)

	var s6 asset.ArchiveExtractionStatus = "test"
	want6 := (*ArchiveExtractionStatus)(nil)
	got6 := ToArchiveExtractionStatus(&s6)
	assert.Equal(t, want6, got6)
}

func TestConvertAsset_ToAssetFile(t *testing.T) {
	c := []*asset.File{}
	f1 := asset.NewFile().Name("aaa.jpg").Size(1000).ContentType("image/jpg").Path("/").Children(c).Build()

	want1 := AssetFile{
		Name:        "aaa.jpg",
		Size:        int64(1000),
		ContentType: lo.ToPtr("image/jpg"),
		Path:        "/",
		Children:    lo.Map(c, func(a *asset.File, _ int) *AssetFile { return ToAssetFile(a) }),
	}

	var f2 *asset.File = nil
	want2 := (*AssetFile)(nil)

	tests := []struct {
		name string
		arg  *asset.File
		want *AssetFile
	}{
		{
			name: "to asset file valid",
			arg:  f1,
			want: &want1,
		},
		{
			name: "to asset file nil",
			arg:  f2,
			want: want2,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := ToAssetFile(tc.arg)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAssetSort_Into(t *testing.T) {
	tests := []struct {
		name string
		sort *AssetSort
		want *usecasex.Sort
	}{
		{
			name: "success",
			sort: &AssetSort{
				SortBy:    "NAME",
				Direction: lo.ToPtr(SortDirectionAsc),
			},
			want: &usecasex.Sort{
				Key:      "filename",
				Reverted: false,
			},
		},
		{
			name: "success",
			sort: &AssetSort{
				SortBy:    "NAME",
				Direction: nil,
			},
			want: &usecasex.Sort{
				Key:      "filename",
				Reverted: false,
			},
		},
		{
			name: "success",
			sort: &AssetSort{
				SortBy:    "NAME",
				Direction: lo.ToPtr(SortDirectionDesc),
			},
			want: &usecasex.Sort{
				Key:      "filename",
				Reverted: true,
			},
		},
		{
			name: "success",
			sort: &AssetSort{
				SortBy:    "NAME",
				Direction: nil,
			},
			want: &usecasex.Sort{
				Key:      "filename",
				Reverted: false,
			},
		},
		{
			name: "success",
			sort: &AssetSort{
				SortBy:    "SIZE",
				Direction: nil,
			},
			want: &usecasex.Sort{
				Key:      "size",
				Reverted: false,
			},
		},
		{
			name: "success",
			sort: &AssetSort{
				SortBy:    "DATE",
				Direction: nil,
			},
			want: &usecasex.Sort{
				Key:      "createdat",
				Reverted: false,
			},
		},
		{
			name: "success",
			sort: &AssetSort{
				SortBy:    "xxx",
				Direction: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sort.Into())
		})
	}
}
