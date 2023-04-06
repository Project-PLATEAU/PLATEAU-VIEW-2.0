package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
)

type AssetLoader struct {
	usecase interfaces.Asset
}

func NewAssetLoader(usecase interfaces.Asset) *AssetLoader {
	return &AssetLoader{usecase: usecase}
}

func (c *AssetLoader) FindByID(ctx context.Context, assetId gqlmodel.ID) (*gqlmodel.Asset, error) {
	aid, err := gqlmodel.ToID[id.Asset](assetId)
	if err != nil {
		return nil, err
	}

	a, err := c.usecase.FindByID(ctx, aid, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return gqlmodel.ToAsset(a, c.usecase.GetURL), nil
}

func (c *AssetLoader) FindByIDs(ctx context.Context, ids []gqlmodel.ID) ([]*gqlmodel.Asset, []error) {
	ids2, err := util.TryMap(ids, gqlmodel.ToID[id.Asset])
	if err != nil {
		return nil, []error{err}
	}

	res, err := c.usecase.FindByIDs(ctx, ids2, getOperator(ctx))
	if err != nil {
		return nil, []error{err}
	}

	return util.Map(res, func(a *asset.Asset) *gqlmodel.Asset {
		return gqlmodel.ToAsset(a, c.usecase.GetURL)
	}), nil
}

func (c *AssetLoader) FindByProject(ctx context.Context, projectId gqlmodel.ID, keyword *string, sort *gqlmodel.AssetSort, p *gqlmodel.Pagination) (*gqlmodel.AssetConnection, error) {
	pid, err := gqlmodel.ToID[id.Project](projectId)
	if err != nil {
		return nil, err
	}

	f := interfaces.AssetFilter{
		Keyword:    keyword,
		Sort:       sort.Into(),
		Pagination: p.Into(),
	}

	assets, pi, err := c.usecase.FindByProject(ctx, pid, f, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	edges := make([]*gqlmodel.AssetEdge, 0, len(assets))
	nodes := make([]*gqlmodel.Asset, 0, len(assets))
	for _, a := range assets {
		asset := gqlmodel.ToAsset(a, c.usecase.GetURL)
		edges = append(edges, &gqlmodel.AssetEdge{
			Node:   asset,
			Cursor: usecasex.Cursor(asset.ID),
		})
		nodes = append(nodes, asset)
	}

	return &gqlmodel.AssetConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   gqlmodel.ToPageInfo(pi),
		TotalCount: int(pi.TotalCount),
	}, nil
}

type AssetDataLoader interface {
	Load(gqlmodel.ID) (*gqlmodel.Asset, error)
	LoadAll([]gqlmodel.ID) ([]*gqlmodel.Asset, []error)
}

func (c *AssetLoader) DataLoader(ctx context.Context) AssetDataLoader {
	return gqldataloader.NewAssetLoader(gqldataloader.AssetLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Asset, []error) {
			return c.FindByIDs(ctx, keys)
		},
	})
}

func (c *AssetLoader) OrdinaryDataLoader(ctx context.Context) AssetDataLoader {
	return &ordinaryAssetLoader{ctx: ctx, c: c}
}

type ordinaryAssetLoader struct {
	ctx context.Context
	c   *AssetLoader
}

func (l *ordinaryAssetLoader) Load(key gqlmodel.ID) (*gqlmodel.Asset, error) {
	res, errs := l.c.FindByIDs(l.ctx, []gqlmodel.ID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryAssetLoader) LoadAll(keys []gqlmodel.ID) ([]*gqlmodel.Asset, []error) {
	return l.c.FindByIDs(l.ctx, keys)
}
