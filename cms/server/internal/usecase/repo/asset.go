package repo

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/usecasex"
)

type AssetFilter struct {
	Sort       *usecasex.Sort
	Keyword    *string
	Pagination *usecasex.Pagination
}

type Asset interface {
	Filtered(ProjectFilter) Asset
	FindByProject(context.Context, id.ProjectID, AssetFilter) ([]*asset.Asset, *usecasex.PageInfo, error)
	FindByID(context.Context, id.AssetID) (*asset.Asset, error)
	FindByIDs(context.Context, id.AssetIDList) ([]*asset.Asset, error)
	Save(context.Context, *asset.Asset) error
	Delete(context.Context, id.AssetID) error
}

type AssetFile interface {
	FindByID(context.Context, id.AssetID) (*asset.File, error)
	Save(context.Context, id.AssetID, *asset.File) error
}
