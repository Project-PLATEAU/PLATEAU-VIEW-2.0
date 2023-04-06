package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearthx/util"
)

type ThreadLoader struct {
	usecase interfaces.Thread
}

func NewThreadLoader(usecase interfaces.Thread) *ThreadLoader {
	return &ThreadLoader{usecase: usecase}
}

type ThreadDataLoader interface {
	Load(gqlmodel.ID) (*gqlmodel.Thread, error)
	LoadAll([]gqlmodel.ID) ([]*gqlmodel.Thread, []error)
}

func (c *ThreadLoader) DataLoader(ctx context.Context) ThreadDataLoader {
	return gqldataloader.NewThreadLoader(gqldataloader.ThreadLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Thread, []error) {
			return c.FindByIDs(ctx, keys)
		},
	})
}

func (c *ThreadLoader) OrdinaryDataLoader(ctx context.Context) ThreadDataLoader {
	return &ordinaryThreadLoader{ctx: ctx, c: c}
}

type ordinaryThreadLoader struct {
	ctx context.Context
	c   *ThreadLoader
}

func (l *ordinaryThreadLoader) Load(key gqlmodel.ID) (*gqlmodel.Thread, error) {
	res, errs := l.c.FindByIDs(l.ctx, []gqlmodel.ID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryThreadLoader) LoadAll(keys []gqlmodel.ID) ([]*gqlmodel.Thread, []error) {
	return l.c.FindByIDs(l.ctx, keys)
}

func (c *ThreadLoader) FindByIDs(ctx context.Context, ids []gqlmodel.ID) ([]*gqlmodel.Thread, []error) {
	ids2, err := util.TryMap(ids, gqlmodel.ToID[id.Thread])
	if err != nil {
		return nil, []error{err}
	}

	res, err := c.usecase.FindByIDs(ctx, ids2, getOperator(ctx))
	if err != nil {
		return nil, []error{err}
	}

	return util.Map(res, func(a *thread.Thread) *gqlmodel.Thread {
		return gqlmodel.ToThread(a)
	}), nil
}
