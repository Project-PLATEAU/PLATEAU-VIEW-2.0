package integrationapi

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestToAssetFile(t *testing.T) {
	f3 := asset.NewFile().Name("aaa").Path("/a/aaa").Build()
	f2 := asset.NewFile().Name("a").Path("/a").Size(10).Children([]*asset.File{f3}).Build()
	f1 := asset.NewFile().Name("").Path("/").Size(11).Children([]*asset.File{f2}).Build()

	a := ToAssetFile(f1, true)
	e := &File{
		Name:        lo.ToPtr(""),
		Path:        lo.ToPtr("/"),
		Size:        lo.ToPtr(float32(11)),
		ContentType: lo.ToPtr(""),
		Children: lo.ToPtr([]File{
			{
				Name:        lo.ToPtr("a"),
				Path:        lo.ToPtr("/a"),
				Size:        lo.ToPtr(float32(10)),
				ContentType: lo.ToPtr(""),
				Children: lo.ToPtr([]File{
					{
						Name:        lo.ToPtr("aaa"),
						Path:        lo.ToPtr("/a/aaa"),
						Size:        lo.ToPtr(float32(0)),
						ContentType: lo.ToPtr(""),
					},
				}),
			},
		}),
	}
	assert.Equal(t, e, a)

	a = ToAssetFile(f1, false)
	e = &File{
		Name:        lo.ToPtr(""),
		Path:        lo.ToPtr("/"),
		Size:        lo.ToPtr(float32(11)),
		ContentType: lo.ToPtr(""),
	}
	assert.Equal(t, e, a)
}
