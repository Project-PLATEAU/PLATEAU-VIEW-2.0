package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
)

func (r *Resolver) Integration() IntegrationResolver {
	return &integrationResolver{r}
}

type integrationResolver struct{ *Resolver }

func (i integrationResolver) Developer(ctx context.Context, obj *gqlmodel.Integration) (*gqlmodel.User, error) {
	return dataloaders(ctx).User.Load(obj.DeveloperID)
}
