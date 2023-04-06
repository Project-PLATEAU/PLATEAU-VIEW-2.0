package ckan

import (
	"context"
	"errors"
	"sort"
	"strings"

	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/thanhpk/randstr"
)

type Mock struct {
	org       string
	packages  *util.SyncMap[string, Package]
	resources *util.SyncMap[string, Resource]
}

func NewMock(org string, packages []Package, resources []Resource) *Mock {
	return &Mock{
		org: org,
		packages: util.SyncMapFrom(lo.SliceToMap(packages, func(p Package) (string, Package) {
			return p.ID, p
		})),
		resources: util.SyncMapFrom(lo.SliceToMap(resources, func(r Resource) (string, Resource) {
			return r.ID, r
		})),
	}
}

func (c *Mock) ShowPackage(ctx context.Context, id string) (Package, error) {
	p, ok := c.packages.Load(id)
	if !ok || p.OwnerOrg != c.org {
		p2 := c.packages.Find(func(_ string, p Package) bool {
			return p.Name == id
		})
		if p2.ID != "" {
			p = p2
		} else {
			return Package{}, rerror.ErrNotFound
		}
	}
	p.Resources = c.resources.FindAll(func(_ string, r Resource) bool {
		return r.PackageID == p.ID
	})
	sort.Slice(p.Resources, func(i, j int) bool {
		return p.Resources[i].Name < p.Resources[j].Name
	})
	return p, nil
}

func (c *Mock) SearchPackageByName(ctx context.Context, name string) (List[Package], error) {
	results := lo.Map(c.packages.FindAll(func(_ string, p Package) bool {
		return strings.Contains(p.Name, name) && p.OwnerOrg == c.org
	}), func(p Package, _ int) Package {
		p.Resources = c.resources.FindAll(func(_ string, r Resource) bool {
			return r.PackageID == p.ID
		})
		sort.Slice(p.Resources, func(i, j int) bool {
			return p.Resources[i].Name < p.Resources[j].Name
		})
		return p
	})
	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})
	return List[Package]{
		Count:   len(results),
		Results: results,
	}, nil
}

func (c *Mock) CreatePackage(ctx context.Context, pkg Package) (Package, error) {
	if pkg.OwnerOrg != c.org {
		return Package{}, errors.New("invalid org")
	}
	if pkg.ID == "" {
		pkg.ID = randstr.Hex(16)
	}
	c.packages.Store(pkg.ID, pkg)
	return pkg, nil
}

func (c *Mock) PatchPackage(ctx context.Context, pkg Package) (Package, error) {
	if pkg.OwnerOrg != c.org {
		return Package{}, errors.New("invalid org")
	}
	// TODO: patch
	pkg.Name = pkg.Name + " PATCHED"
	c.packages.Store(pkg.ID, pkg)
	return pkg, nil
}

func (c *Mock) SavePackage(ctx context.Context, pkg Package) (Package, error) {
	if pkg.ID == "" {
		return c.CreatePackage(ctx, pkg)
	}
	return c.PatchPackage(ctx, pkg)
}

func (c *Mock) CreateResource(ctx context.Context, resource Resource) (Resource, error) {
	p, ok := c.packages.Load(resource.PackageID)
	if !ok {
		return Resource{}, errors.New("invalid package")
	}
	if p.OwnerOrg != c.org {
		return Resource{}, errors.New("invalid org")
	}

	if resource.ID == "" {
		resource.ID = randstr.Hex(16)
	}
	c.resources.Store(resource.ID, resource)
	return resource, nil
}

func (c *Mock) PatchResource(ctx context.Context, resource Resource) (Resource, error) {
	p, ok := c.packages.Load(resource.PackageID)
	if !ok {
		return Resource{}, errors.New("invalid package")
	}
	if p.OwnerOrg != c.org {
		return Resource{}, errors.New("invalid org")
	}

	// TODO: patch
	resource.Name = resource.Name + " PATCHED"
	c.resources.Store(resource.ID, resource)
	return resource, nil
}

func (c *Mock) UploadResource(ctx context.Context, resource Resource, filename string, data []byte) (Resource, error) {
	return c.SaveResource(ctx, resource)
}

func (c *Mock) SaveResource(ctx context.Context, resource Resource) (Resource, error) {
	if resource.ID == "" {
		return c.CreateResource(ctx, resource)
	}
	return c.PatchResource(ctx, resource)
}
