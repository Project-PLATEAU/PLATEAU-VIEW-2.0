package gql

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

type ItemLoader struct {
	usecase       interfaces.Item
	schemaUsecase interfaces.Schema
}

func NewItemLoader(usecase interfaces.Item, schemaUsecase interfaces.Schema) *ItemLoader {
	return &ItemLoader{usecase: usecase, schemaUsecase: schemaUsecase}
}
func (c *ItemLoader) Fetch(ctx context.Context, ids []gqlmodel.ID) ([]*gqlmodel.Item, []error) {
	op := getOperator(ctx)
	iIds, err := util.TryMap(ids, gqlmodel.ToID[id.Item])
	if err != nil {
		return nil, []error{err}
	}

	res, err := c.usecase.FindByIDs(ctx, iIds, op)
	if err != nil {
		return nil, []error{err}
	}

	sIds := lo.SliceToMap(res, func(v item.Versioned) (id.ItemID, id.SchemaID) {
		return v.Value().ID(), v.Value().Schema()
	})

	ss, err := c.schemaUsecase.FindByIDs(ctx, lo.Uniq(lo.Values(sIds)), op)
	if err != nil {
		return nil, []error{err}
	}

	return lo.Map(res, func(m item.Versioned, i int) *gqlmodel.Item {
		s, _ := lo.Find(ss, func(s *schema.Schema) bool {
			return s.ID() == sIds[m.Value().ID()]
		})
		return gqlmodel.ToItem(m.Value(), s)
	}), nil
}

func (c *ItemLoader) FindVersionedItem(ctx context.Context, itemID gqlmodel.ID) (*gqlmodel.VersionedItem, error) {
	op := getOperator(ctx)
	iId, err := gqlmodel.ToID[id.Item](itemID)
	if err != nil {
		return nil, err
	}

	itm, err := c.usecase.FindByID(ctx, iId, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	s, err := c.schemaUsecase.FindByID(ctx, itm.Value().Schema(), op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return gqlmodel.ToVersionedItem(itm, s), nil
}

func (c *ItemLoader) FindVersionedItems(ctx context.Context, itemID gqlmodel.ID) ([]*gqlmodel.VersionedItem, error) {
	op := getOperator(ctx)
	iId, err := gqlmodel.ToID[id.Item](itemID)
	if err != nil {
		return nil, err
	}

	res, err := c.usecase.FindAllVersionsByID(ctx, iId, op)
	if err != nil {
		return nil, err
	}

	s, err := c.schemaUsecase.FindByID(ctx, res[0].Value().Schema(), op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	vis := make([]*gqlmodel.VersionedItem, 0, len(res))
	for _, t := range res {
		vis = append(vis, gqlmodel.ToVersionedItem(t, s))
	}
	return vis, nil
}

func (c *ItemLoader) FindBySchema(ctx context.Context, schemaID gqlmodel.ID, sort *gqlmodel.ItemSort, p *gqlmodel.Pagination) (*gqlmodel.ItemConnection, error) {
	op := getOperator(ctx)
	sid, err := gqlmodel.ToID[id.Schema](schemaID)
	if err != nil {
		return nil, err
	}

	s, err := c.schemaUsecase.FindByID(ctx, sid, op)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}

	res, pi, err := c.usecase.FindBySchema(ctx, sid, sort.Into(), p.Into(), op)
	if err != nil {
		return nil, err
	}

	edges := make([]*gqlmodel.ItemEdge, 0, len(res))
	nodes := make([]*gqlmodel.Item, 0, len(res))
	for _, i := range res {
		itm := gqlmodel.ToItem(i.Value(), s)
		edges = append(edges, &gqlmodel.ItemEdge{
			Node:   itm,
			Cursor: usecasex.Cursor(itm.ID),
		})
		nodes = append(nodes, itm)
	}

	return &gqlmodel.ItemConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   gqlmodel.ToPageInfo(pi),
		TotalCount: int(pi.TotalCount),
	}, nil
}

func (c *ItemLoader) FindByProject(ctx context.Context, projectID gqlmodel.ID, p *gqlmodel.Pagination) (*gqlmodel.ItemConnection, error) {
	op := getOperator(ctx)
	pid, err := gqlmodel.ToID[id.Project](projectID)
	if err != nil {
		return nil, err
	}

	res, pi, err := c.usecase.FindByProject(ctx, pid, p.Into(), op)
	if err != nil {
		return nil, err
	}

	sIds := lo.SliceToMap(res, func(v item.Versioned) (id.ItemID, id.SchemaID) {
		return v.Value().ID(), v.Value().Schema()
	})

	ss, err := c.schemaUsecase.FindByIDs(ctx, lo.Uniq(lo.Values(sIds)), op)
	if err != nil {
		return nil, err
	}

	edges := make([]*gqlmodel.ItemEdge, 0, len(res))
	nodes := make([]*gqlmodel.Item, 0, len(res))
	for _, i := range res {
		s, _ := lo.Find(ss, func(s *schema.Schema) bool {
			return s.ID() == sIds[i.Value().ID()]
		})
		itm := gqlmodel.ToItem(i.Value(), s)
		edges = append(edges, &gqlmodel.ItemEdge{
			Node:   itm,
			Cursor: usecasex.Cursor(itm.ID),
		})
		nodes = append(nodes, itm)
	}

	return &gqlmodel.ItemConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   gqlmodel.ToPageInfo(pi),
		TotalCount: int(pi.TotalCount),
	}, nil
}

func (c *ItemLoader) Search(ctx context.Context, query gqlmodel.ItemQuery, sort *gqlmodel.ItemSort, p *gqlmodel.Pagination) (*gqlmodel.ItemConnection, error) {
	op := getOperator(ctx)
	q := gqlmodel.ToItemQuery(query)
	res, pi, err := c.usecase.Search(ctx, q, sort.Into(), p.Into(), op)
	if err != nil {
		return nil, err
	}

	sIds := lo.SliceToMap(res, func(v item.Versioned) (id.ItemID, id.SchemaID) {
		return v.Value().ID(), v.Value().Schema()
	})

	ss, err := c.schemaUsecase.FindByIDs(ctx, lo.Uniq(lo.Values(sIds)), op)
	if err != nil {
		return nil, err
	}

	edges := make([]*gqlmodel.ItemEdge, 0, len(res))
	nodes := make([]*gqlmodel.Item, 0, len(res))
	for _, i := range res {
		s, _ := lo.Find(ss, func(s *schema.Schema) bool {
			return s.ID() == sIds[i.Value().ID()]
		})
		itm := gqlmodel.ToItem(i.Value(), s)
		edges = append(edges, &gqlmodel.ItemEdge{
			Node:   itm,
			Cursor: usecasex.Cursor(itm.ID),
		})
		nodes = append(nodes, itm)
	}

	return &gqlmodel.ItemConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   gqlmodel.ToPageInfo(pi),
		TotalCount: int(pi.TotalCount),
	}, nil
}

// data loader

type ItemDataLoader interface {
	Load(gqlmodel.ID) (*gqlmodel.Item, error)
	LoadAll([]gqlmodel.ID) ([]*gqlmodel.Item, []error)
}

func (c *ItemLoader) DataLoader(ctx context.Context) ItemDataLoader {
	return gqldataloader.NewItemLoader(gqldataloader.ItemLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Item, []error) {
			return c.Fetch(ctx, keys)
		},
	})
}

func (c *ItemLoader) OrdinaryDataLoader(ctx context.Context) ItemDataLoader {
	return &ordinaryItemLoader{
		fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Item, []error) {
			return c.Fetch(ctx, keys)
		},
	}
}

type ordinaryItemLoader struct {
	fetch func(keys []gqlmodel.ID) ([]*gqlmodel.Item, []error)
}

func (l *ordinaryItemLoader) Load(key gqlmodel.ID) (*gqlmodel.Item, error) {
	res, errs := l.fetch([]gqlmodel.ID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryItemLoader) LoadAll(keys []gqlmodel.ID) ([]*gqlmodel.Item, []error) {
	return l.fetch(keys)
}
