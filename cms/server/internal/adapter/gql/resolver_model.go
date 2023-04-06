package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
)

func (r *Resolver) Model() ModelResolver {
	return &modelResolver{r}
}

type modelResolver struct{ *Resolver }

func (m modelResolver) Project(ctx context.Context, obj *gqlmodel.Model) (*gqlmodel.Project, error) {
	return dataloaders(ctx).Project.Load(obj.ProjectID)
}

func (m modelResolver) Schema(ctx context.Context, obj *gqlmodel.Model) (*gqlmodel.Schema, error) {
	return dataloaders(ctx).Schema.Load(obj.SchemaID)
}
