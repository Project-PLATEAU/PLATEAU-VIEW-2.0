package memory

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
)

type AssetFile struct {
	data *util.SyncMap[asset.ID, *asset.File]
	err  error
}

func NewAssetFile() *AssetFile {
	return &AssetFile{
		data: &util.SyncMap[id.AssetID, *asset.File]{},
	}
}

func (r *AssetFile) FindByID(ctx context.Context, id id.AssetID) (*asset.File, error) {
	if r.err != nil {
		return nil, r.err
	}

	f := r.data.Find(func(key asset.ID, value *asset.File) bool {
		return key == id
	}).Clone()
	return rerror.ErrIfNil(f, rerror.ErrNotFound)
}

func (r *AssetFile) Save(ctx context.Context, id id.AssetID, file *asset.File) error {
	if r.err != nil {
		return r.err
	}

	r.data.Store(id, file.Clone())
	return nil
}
