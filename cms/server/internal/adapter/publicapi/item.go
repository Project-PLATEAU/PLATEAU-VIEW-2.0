package publicapi

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

func (c *Controller) GetItem(ctx context.Context, prj, mkey, i string) (Item, error) {
	pr, err := c.checkProject(ctx, prj)
	if err != nil {
		return Item{}, err
	}

	if mkey == "" {
		return Item{}, rerror.ErrNotFound
	}

	iid, err := id.ItemIDFrom(i)
	if err != nil {
		return Item{}, rerror.ErrNotFound
	}

	it, err := c.usecases.Item.FindPublicByID(ctx, iid, nil)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return Item{}, rerror.ErrNotFound
		}
		return Item{}, err
	}

	itv := it.Value()
	m, err := c.usecases.Model.FindByID(ctx, itv.Model(), nil)
	if err != nil {
		return Item{}, err
	}

	if m.Key().String() != mkey || !m.Public() {
		return Item{}, rerror.ErrNotFound
	}

	s, err := c.usecases.Schema.FindByID(ctx, m.Schema(), nil)
	if err != nil {
		return Item{}, err
	}

	var assets asset.List
	if pr.Publication().AssetPublic() {
		assets, err = c.usecases.Asset.FindByIDs(ctx, itv.AssetIDs(), nil)
		if err != nil {
			return Item{}, err
		}
	}

	return NewItem(itv, s, assets, c.assetUrlResolver), nil
}

func (c *Controller) GetItems(ctx context.Context, prj, model string, p ListParam) (ListResult[Item], error) {
	pr, err := c.checkProject(ctx, prj)
	if err != nil {
		return ListResult[Item]{}, err
	}

	m, err := c.usecases.Model.FindByKey(ctx, pr.ID(), model, nil)
	if err != nil {
		return ListResult[Item]{}, err
	}
	if !m.Public() {
		return ListResult[Item]{}, rerror.ErrNotFound
	}

	s, err := c.usecases.Schema.FindByID(ctx, m.Schema(), nil)
	if err != nil {
		return ListResult[Item]{}, err
	}

	items, pi, err := c.usecases.Item.FindPublicByModel(ctx, m.ID(), p.Pagination, nil)
	if err != nil {
		return ListResult[Item]{}, err
	}

	var assets asset.List
	if pr.Publication().AssetPublic() {
		assetIDs := lo.FlatMap(items.Unwrap(), func(i *item.Item, _ int) []id.AssetID {
			return i.AssetIDs()
		})
		assets, err = c.usecases.Asset.FindByIDs(ctx, assetIDs, nil)
		if err != nil {
			return ListResult[Item]{}, err
		}
	}

	res := NewListResult(util.Map(items.Unwrap(), func(i *item.Item) Item {
		return NewItem(i, s, assets, c.assetUrlResolver)
	}), pi, p.Pagination)
	return res, nil
}
