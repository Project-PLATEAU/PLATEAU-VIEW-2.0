package ckan

import (
	"context"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestMock(t *testing.T) {
	ctx := context.Background()
	m := NewMock("org", nil, nil)

	pkg, err := m.SavePackage(ctx, Package{
		ID:       "",
		Name:     "name",
		OwnerOrg: "org",
		Title:    "title",
	})
	assert.NoError(t, err)
	assert.Equal(t, Package{
		ID:       pkg.ID,
		Name:     "name",
		OwnerOrg: "org",
		Title:    "title",
	}, pkg)

	r, err := m.SaveResource(ctx, Resource{
		ID:        "",
		PackageID: pkg.ID,
		URL:       "https://example.com",
	})
	assert.NoError(t, err)
	assert.Equal(t, Resource{
		ID:        r.ID,
		PackageID: pkg.ID,
		URL:       "https://example.com",
	}, r)

	pkg2, err := m.ShowPackage(ctx, "name")
	assert.NoError(t, err)
	assert.Equal(t, Package{
		ID:       pkg.ID,
		Name:     "name",
		OwnerOrg: "org",
		Title:    "title",
		Resources: []Resource{
			{
				ID:        r.ID,
				PackageID: pkg.ID,
				URL:       "https://example.com",
			},
		},
	}, pkg2)

	pkgs, err := m.SearchPackageByName(ctx, "name")
	assert.NoError(t, err)
	assert.Equal(t, List[Package]{
		Count: 1,
		Results: []Package{{
			ID:       pkg.ID,
			Name:     "name",
			OwnerOrg: "org",
			Title:    "title",
			Resources: []Resource{
				{
					ID:        r.ID,
					PackageID: pkg.ID,
					URL:       "https://example.com",
				},
			},
		}},
	}, pkgs)

	assert.Equal(t, Resource{
		ID:        r.ID,
		PackageID: r.PackageID,
		Name:      r.Name + " PATCHED",
		URL:       r.URL,
	}, lo.Must(m.UploadResource(ctx, r, "", nil)))
}
