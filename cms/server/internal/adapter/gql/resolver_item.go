package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/samber/lo"
)

func (r *Resolver) Item() ItemResolver {
	return &itemResolver{r}
}

type itemResolver struct{ *Resolver }

func (i itemResolver) Project(ctx context.Context, obj *gqlmodel.Item) (*gqlmodel.Project, error) {
	return dataloaders(ctx).Project.Load(obj.ProjectID)
}

func (i itemResolver) Schema(ctx context.Context, obj *gqlmodel.Item) (*gqlmodel.Schema, error) {
	return dataloaders(ctx).Schema.Load(obj.SchemaID)
}

func (i itemResolver) Thread(ctx context.Context, obj *gqlmodel.Item) (*gqlmodel.Thread, error) {
	return dataloaders(ctx).Thread.Load(obj.ThreadID)
}

func (i itemResolver) Model(ctx context.Context, obj *gqlmodel.Item) (*gqlmodel.Model, error) {
	return dataloaders(ctx).Model.Load(obj.ModelID)
}

func (i itemResolver) User(ctx context.Context, obj *gqlmodel.Item) (*gqlmodel.User, error) {
	if obj.UserID != nil {
		return dataloaders(ctx).User.Load(*obj.UserID)
	}
	return nil, nil
}

func (i itemResolver) Integration(ctx context.Context, obj *gqlmodel.Item) (*gqlmodel.Integration, error) {
	if obj.IntegrationID != nil {
		return dataloaders(ctx).Integration.Load(*obj.IntegrationID)
	}
	return nil, nil
}

func (i itemResolver) Status(ctx context.Context, obj *gqlmodel.Item) (gqlmodel.ItemStatus, error) {
	return dataloaders(ctx).ItemStatus.Load(obj.ID)
}

func (i itemResolver) Assets(ctx context.Context, obj *gqlmodel.Item) ([]*gqlmodel.Asset, error) {

	aIds := lo.FlatMap(obj.Fields, func(f *gqlmodel.ItemField, _ int) []gqlmodel.ID {
		if f.Type != gqlmodel.SchemaFieldTypeAsset || f.Value == nil {
			return nil
		}
		if s, ok := f.Value.(string); ok {
			return []gqlmodel.ID{gqlmodel.ID(s)}
		}
		if ss, ok := f.Value.([]any); ok {
			return lo.FilterMap(ss, func(i any, _ int) (gqlmodel.ID, bool) {
				str, ok := i.(string)
				return gqlmodel.ID(str), ok
			})
		}
		return nil
	})

	assets, err := dataloaders(ctx).Asset.LoadAll(aIds)
	if len(err) > 0 && err[0] != nil {
		return nil, err[0]
	}
	return assets, nil
}
