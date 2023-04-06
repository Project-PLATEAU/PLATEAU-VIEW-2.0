package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

type RequestLoader struct {
	usecase interfaces.Request
}

func NewRequestLoader(usecase interfaces.Request) *RequestLoader {
	return &RequestLoader{usecase: usecase}
}

func (c *RequestLoader) Fetch(ctx context.Context, ids []gqlmodel.ID) ([]*gqlmodel.Request, []error) {
	ids2, err := util.TryMap(ids, gqlmodel.ToID[id.Request])
	if err != nil {
		return nil, []error{err}
	}

	res, err := c.usecase.FindByIDs(ctx, ids2, getOperator(ctx))
	if err != nil {
		return nil, []error{err}
	}

	return util.Map(res, func(req *request.Request) *gqlmodel.Request {
		return gqlmodel.ToRequest(req)
	}), nil
}

func (c *RequestLoader) FindByProject(ctx context.Context, projectId gqlmodel.ID, keyword *string, state []gqlmodel.RequestState, createdBy, reviewer *gqlmodel.ID, p *gqlmodel.Pagination, sort *gqlmodel.Sort) (*gqlmodel.RequestConnection, error) {
	pid, err := gqlmodel.ToID[id.Project](projectId)
	if err != nil {
		return nil, err
	}

	f := interfaces.RequestFilter{
		Keyword: keyword,
	}
	if state != nil {
		f.State = lo.Map(state, func(s gqlmodel.RequestState, _ int) request.State {
			return request.StateFrom(s.String())
		})
	}
	f.Reviewer = gqlmodel.ToIDRef[id.User](reviewer)
	f.CreatedBy = gqlmodel.ToIDRef[id.User](createdBy)

	requests, pi, err := c.usecase.FindByProject(ctx, pid, f, sort.Into(), p.Into(), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	edges := make([]*gqlmodel.RequestEdge, 0, len(requests))
	nodes := make([]*gqlmodel.Request, 0, len(requests))
	for _, req := range requests {
		request := gqlmodel.ToRequest(req)
		edges = append(edges, &gqlmodel.RequestEdge{
			Node:   request,
			Cursor: usecasex.Cursor(request.ID),
		})
		nodes = append(nodes, request)
	}

	return &gqlmodel.RequestConnection{
		Edges:      edges,
		Nodes:      nodes,
		PageInfo:   gqlmodel.ToPageInfo(pi),
		TotalCount: int(pi.TotalCount),
	}, nil
}

type RequestDataLoader interface {
	Load(gqlmodel.ID) (*gqlmodel.Request, error)
	LoadAll([]gqlmodel.ID) ([]*gqlmodel.Request, []error)
}

func (c *RequestLoader) DataLoader(ctx context.Context) RequestDataLoader {
	return gqldataloader.NewRequestLoader(gqldataloader.RequestLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Request, []error) {
			return c.Fetch(ctx, keys)
		},
	})
}

func (c *RequestLoader) OrdinaryDataLoader(ctx context.Context) RequestDataLoader {
	return &ordinaryRequestLoader{ctx: ctx, c: c}
}

type ordinaryRequestLoader struct {
	ctx context.Context
	c   *RequestLoader
}

func (l *ordinaryRequestLoader) Load(key gqlmodel.ID) (*gqlmodel.Request, error) {
	res, errs := l.c.Fetch(l.ctx, []gqlmodel.ID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryRequestLoader) LoadAll(keys []gqlmodel.ID) ([]*gqlmodel.Request, []error) {
	return l.c.Fetch(l.ctx, keys)
}
