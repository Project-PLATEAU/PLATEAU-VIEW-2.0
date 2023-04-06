package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
)

func (r *Resolver) WorkspaceUserMember() WorkspaceUserMemberResolver {
	return &workspaceUserMemberResolver{r}
}

type workspaceUserMemberResolver struct{ *Resolver }

func (w workspaceUserMemberResolver) User(ctx context.Context, obj *gqlmodel.WorkspaceUserMember) (*gqlmodel.User, error) {
	return dataloaders(ctx).User.Load(obj.UserID)
}

func (r *Resolver) WorkspaceIntegrationMember() WorkspaceIntegrationMemberResolver {
	return &workspaceIntegrationMemberResolver{r}
}

type workspaceIntegrationMemberResolver struct{ *Resolver }

func (w workspaceIntegrationMemberResolver) InvitedBy(ctx context.Context, obj *gqlmodel.WorkspaceIntegrationMember) (*gqlmodel.User, error) {
	return dataloaders(ctx).User.Load(obj.InvitedByID)
}

func (w workspaceIntegrationMemberResolver) Integration(ctx context.Context, obj *gqlmodel.WorkspaceIntegrationMember) (*gqlmodel.Integration, error) {
	return dataloaders(ctx).Integration.Load(obj.IntegrationID)
}
