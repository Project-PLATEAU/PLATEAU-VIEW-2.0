package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

type ModelLoader struct {
	usecase interfaces.Model
}

func NewModelLoader(usecase interfaces.Model) *ModelLoader {
	return &ModelLoader{usecase: usecase}
}

func (c *ModelLoader) Fetch(ctx context.Context, ids []gqlmodel.ID) ([]*gqlmodel.Model, []error) {
	mIds, err := util.TryMap(ids, gqlmodel.ToID[id.Model])
	if err != nil {
		return nil, []error{err}
	}

	res, err := c.usecase.FindByIDs(ctx, mIds, getOperator(ctx))
	if err != nil {
		return nil, []error{err}
	}

	return lo.Map(res, func(m *model.Model, i int) *gqlmodel.Model {
		return gqlmodel.ToModel(m)
	}), nil
}

func (c *ModelLoader) FindByProject(ctx context.Context, projectId gqlmodel.ID, p *gqlmodel.Pagination) (*gqlmodel.ModelConnection, error) {
	pId, err := gqlmodel.ToID[id.Project](projectId)
	if err != nil {
		return nil, err
	}

	res, pi, err := c.usecase.FindByProject(ctx, pId, p.Into(), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	edges := make([]*gqlmodel.ModelEdge, 0, len(res))
	nodes := make([]*gqlmodel.Model, 0, len(res))
	for _, r := range res {
		m := gqlmodel.ToModel(r)
		edges = append(edges, &gqlmodel.ModelEdge{
			Node:   m,
			Cursor: usecasex.Cursor(m.ID),
		})
		nodes = append(nodes, m)
	}

	return &gqlmodel.ModelConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   gqlmodel.ToPageInfo(pi),
		TotalCount: int(pi.TotalCount),
	}, nil
}

func (c *ModelLoader) CheckKey(ctx context.Context, projectID gqlmodel.ID, key string) (*gqlmodel.KeyAvailability, error) {
	pId, err := gqlmodel.ToID[id.Project](projectID)
	if err != nil {
		return nil, err
	}

	ok, err := c.usecase.CheckKey(ctx, pId, key)
	if err != nil {
		return nil, err
	}

	return &gqlmodel.KeyAvailability{Key: key, Available: ok}, nil
}

// data loaders

type ModelDataLoader interface {
	Load(gqlmodel.ID) (*gqlmodel.Model, error)
	LoadAll([]gqlmodel.ID) ([]*gqlmodel.Model, []error)
}

func (c *ModelLoader) DataLoader(ctx context.Context) ModelDataLoader {
	return gqldataloader.NewModelLoader(gqldataloader.ModelLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Model, []error) {
			return c.Fetch(ctx, keys)
		},
	})
}

func (c *ModelLoader) OrdinaryDataLoader(ctx context.Context) ModelDataLoader {
	return &ordinaryModelLoader{
		fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Model, []error) {
			return c.Fetch(ctx, keys)
		},
	}
}

type ordinaryModelLoader struct {
	fetch func(keys []gqlmodel.ID) ([]*gqlmodel.Model, []error)
}

func (l *ordinaryModelLoader) Load(key gqlmodel.ID) (*gqlmodel.Model, error) {
	res, errs := l.fetch([]gqlmodel.ID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryModelLoader) LoadAll(keys []gqlmodel.ID) ([]*gqlmodel.Model, []error) {
	return l.fetch(keys)
}
