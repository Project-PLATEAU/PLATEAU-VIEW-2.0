package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/samber/lo"
)

func (r *Resolver) Comment() CommentResolver {
	return &commentResolver{r}
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) Author(ctx context.Context, obj *gqlmodel.Comment) (gqlmodel.Operator, error) {
	switch obj.AuthorType {
	case gqlmodel.OperatorTypeUser:
		ws, err := loaders(ctx).Workspace.FindByUser(ctx, obj.AuthorID)
		if err != nil {
			return nil, err
		}
		ok := lo.ContainsBy(ws, func(w *gqlmodel.Workspace) bool {
			return w != nil && (w.ID == obj.WorkspaceID)
		})
		if !ok {
			return nil, nil
		}
		return dataloaders(ctx).User.Load(obj.AuthorID)
	case gqlmodel.OperatorTypeIntegration:
		return dataloaders(ctx).Integration.Load(obj.AuthorID)
	default:
		return nil, nil
	}
}
