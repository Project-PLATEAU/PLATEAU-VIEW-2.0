package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

type IntegrationLoader struct {
	usecase interfaces.Integration
}

func NewIntegrationLoader(usecase interfaces.Integration) *IntegrationLoader {
	return &IntegrationLoader{usecase: usecase}
}

func (c *IntegrationLoader) Fetch(ctx context.Context, ids []gqlmodel.ID) ([]*gqlmodel.Integration, []error) {
	sIds, err := util.TryMap(ids, gqlmodel.ToID[id.Integration])
	if err != nil {
		return nil, []error{err}
	}

	op := getOperator(ctx)
	if err != nil {
		return nil, []error{err}
	}

	res, err := c.usecase.FindByIDs(ctx, sIds, op)
	if err != nil {
		return nil, []error{err}
	}

	return lo.Map(res, func(m *integration.Integration, _ int) *gqlmodel.Integration {
		return gqlmodel.ToIntegration(m, op.User)
	}), nil
}

func (c *IntegrationLoader) FindByMe(ctx context.Context) ([]*gqlmodel.Integration, error) {
	op := getOperator(ctx)

	res, err := c.usecase.FindByMe(ctx, op)
	if err != nil {
		return nil, err
	}
	integrations := make([]*gqlmodel.Integration, 0, len(res))
	for _, i := range res {
		integrations = append(integrations, gqlmodel.ToIntegration(i, op.User))
	}
	return integrations, nil
}

// data loaders

type IntegrationDataLoader interface {
	Load(gqlmodel.ID) (*gqlmodel.Integration, error)
	LoadAll([]gqlmodel.ID) ([]*gqlmodel.Integration, []error)
}

func (c *IntegrationLoader) DataLoader(ctx context.Context) IntegrationDataLoader {
	return gqldataloader.NewIntegrationLoader(gqldataloader.IntegrationLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Integration, []error) {
			return c.Fetch(ctx, keys)
		},
	})
}

func (c *IntegrationLoader) OrdinaryDataLoader(ctx context.Context) IntegrationDataLoader {
	return &ordinaryIntegrationLoader{
		fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Integration, []error) {
			return c.Fetch(ctx, keys)
		},
	}
}

type ordinaryIntegrationLoader struct {
	fetch func(keys []gqlmodel.ID) ([]*gqlmodel.Integration, []error)
}

func (l *ordinaryIntegrationLoader) Load(key gqlmodel.ID) (*gqlmodel.Integration, error) {
	res, errs := l.fetch([]gqlmodel.ID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinaryIntegrationLoader) LoadAll(keys []gqlmodel.ID) ([]*gqlmodel.Integration, []error) {
	return l.fetch(keys)
}
