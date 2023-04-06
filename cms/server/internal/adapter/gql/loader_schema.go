package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqldataloader"
	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

type SchemaLoader struct {
	usecase interfaces.Schema
}

func NewSchemaLoader(usecase interfaces.Schema) *SchemaLoader {
	return &SchemaLoader{usecase: usecase}
}

func (c *SchemaLoader) Fetch(ctx context.Context, ids []gqlmodel.ID) ([]*gqlmodel.Schema, []error) {
	sIds, err := util.TryMap(ids, gqlmodel.ToID[id.Schema])
	if err != nil {
		return nil, []error{err}
	}

	res, err := c.usecase.FindByIDs(ctx, sIds, getOperator(ctx))
	if err != nil {
		return nil, []error{err}
	}

	return lo.Map(res, func(m *schema.Schema, _ int) *gqlmodel.Schema {
		return gqlmodel.ToSchema(m)
	}), nil
}

// data loaders

type SchemaDataLoader interface {
	Load(gqlmodel.ID) (*gqlmodel.Schema, error)
	LoadAll([]gqlmodel.ID) ([]*gqlmodel.Schema, []error)
}

func (c *SchemaLoader) DataLoader(ctx context.Context) SchemaDataLoader {
	return gqldataloader.NewSchemaLoader(gqldataloader.SchemaLoaderConfig{
		Wait:     dataLoaderWait,
		MaxBatch: dataLoaderMaxBatch,
		Fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Schema, []error) {
			return c.Fetch(ctx, keys)
		},
	})
}

func (c *SchemaLoader) OrdinaryDataLoader(ctx context.Context) SchemaDataLoader {
	return &ordinarySchemaLoader{
		fetch: func(keys []gqlmodel.ID) ([]*gqlmodel.Schema, []error) {
			return c.Fetch(ctx, keys)
		},
	}
}

type ordinarySchemaLoader struct {
	fetch func(keys []gqlmodel.ID) ([]*gqlmodel.Schema, []error)
}

func (l *ordinarySchemaLoader) Load(key gqlmodel.ID) (*gqlmodel.Schema, error) {
	res, errs := l.fetch([]gqlmodel.ID{key})
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if len(res) > 0 {
		return res[0], nil
	}
	return nil, nil
}

func (l *ordinarySchemaLoader) LoadAll(keys []gqlmodel.ID) ([]*gqlmodel.Schema, []error) {
	return l.fetch(keys)
}
