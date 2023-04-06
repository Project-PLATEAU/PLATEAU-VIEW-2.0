package publicapi

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/rerror"
)

func (c *Controller) GetAsset(ctx context.Context, prj, i string) (Asset, error) {
	_, err := c.checkProject(ctx, prj)
	if err != nil {
		return Asset{}, err
	}

	iid, err := id.AssetIDFrom(i)
	if err != nil {
		return Asset{}, rerror.ErrNotFound
	}

	a, err := c.usecases.Asset.FindByID(ctx, iid, nil)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return Asset{}, rerror.ErrNotFound
		}
		return Asset{}, err
	}

	f, err := c.usecases.Asset.FindFileByID(ctx, iid, nil)
	if err != nil {
		return Asset{}, err
	}

	return NewAsset(a, f, c.assetUrlResolver), nil
}
