package gql

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter/gql/gqlmodel"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"
)

func (r *mutationResolver) CreateWorkspace(ctx context.Context, input gqlmodel.CreateWorkspaceInput) (*gqlmodel.CreateWorkspacePayload, error) {
	res, err := usecases(ctx).Workspace.Create(ctx, input.Name, getUser(ctx).ID(), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.CreateWorkspacePayload{Workspace: gqlmodel.ToWorkspace(res)}, nil
}

func (r *mutationResolver) DeleteWorkspace(ctx context.Context, input gqlmodel.DeleteWorkspaceInput) (*gqlmodel.DeleteWorkspacePayload, error) {
	tid, err := gqlmodel.ToID[id.Workspace](input.WorkspaceID)
	if err != nil {
		return nil, err
	}

	if err := usecases(ctx).Workspace.Remove(ctx, tid, getOperator(ctx)); err != nil {
		return nil, err
	}

	return &gqlmodel.DeleteWorkspacePayload{WorkspaceID: input.WorkspaceID}, nil
}

func (r *mutationResolver) UpdateWorkspace(ctx context.Context, input gqlmodel.UpdateWorkspaceInput) (*gqlmodel.UpdateWorkspacePayload, error) {
	tid, err := gqlmodel.ToID[id.Workspace](input.WorkspaceID)
	if err != nil {
		return nil, err
	}

	res, err := usecases(ctx).Workspace.Update(ctx, tid, input.Name, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.UpdateWorkspacePayload{Workspace: gqlmodel.ToWorkspace(res)}, nil
}

func (r *mutationResolver) AddUsersToWorkspace(ctx context.Context, input gqlmodel.AddUsersToWorkspaceInput) (*gqlmodel.AddUsersToWorkspacePayload, error) {
	wid, err := gqlmodel.ToID[id.Workspace](input.WorkspaceID)
	if err != nil {
		return nil, err
	}
	usersMap := make(map[id.UserID]user.Role, len(input.Users))
	for _, u := range input.Users {
		uid, err := gqlmodel.ToID[id.User](u.UserID)
		if err != nil {
			return nil, err
		}
		usersMap[uid] = gqlmodel.FromRole(u.Role)
	}
	res, err := usecases(ctx).Workspace.AddUserMember(ctx, wid, usersMap, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.AddUsersToWorkspacePayload{Workspace: gqlmodel.ToWorkspace(res)}, nil
}

func (r *mutationResolver) AddIntegrationToWorkspace(ctx context.Context, input gqlmodel.AddIntegrationToWorkspaceInput) (*gqlmodel.AddUsersToWorkspacePayload, error) {
	wId, iId, err := gqlmodel.ToID2[id.Workspace, id.Integration](input.WorkspaceID, input.IntegrationID)
	if err != nil {
		return nil, err
	}

	res, err := usecases(ctx).Workspace.AddIntegrationMember(ctx, wId, iId, gqlmodel.FromRole(input.Role), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.AddUsersToWorkspacePayload{Workspace: gqlmodel.ToWorkspace(res)}, nil
}

func (r *mutationResolver) RemoveUserFromWorkspace(ctx context.Context, input gqlmodel.RemoveUserFromWorkspaceInput) (*gqlmodel.RemoveMemberFromWorkspacePayload, error) {
	tid, uid, err := gqlmodel.ToID2[id.Workspace, id.User](input.WorkspaceID, input.UserID)
	if err != nil {
		return nil, err
	}

	res, err := usecases(ctx).Workspace.RemoveUser(ctx, tid, uid, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.RemoveMemberFromWorkspacePayload{Workspace: gqlmodel.ToWorkspace(res)}, nil
}

func (r *mutationResolver) RemoveIntegrationFromWorkspace(ctx context.Context, input gqlmodel.RemoveIntegrationFromWorkspaceInput) (*gqlmodel.RemoveMemberFromWorkspacePayload, error) {
	wId, iId, err := gqlmodel.ToID2[id.Workspace, id.Integration](input.WorkspaceID, input.IntegrationID)
	if err != nil {
		return nil, err
	}

	res, err := usecases(ctx).Workspace.RemoveIntegration(ctx, wId, iId, getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.RemoveMemberFromWorkspacePayload{Workspace: gqlmodel.ToWorkspace(res)}, nil
}

func (r *mutationResolver) UpdateUserOfWorkspace(ctx context.Context, input gqlmodel.UpdateUserOfWorkspaceInput) (*gqlmodel.UpdateMemberOfWorkspacePayload, error) {
	tid, uid, err := gqlmodel.ToID2[id.Workspace, id.User](input.WorkspaceID, input.UserID)
	if err != nil {
		return nil, err
	}

	res, err := usecases(ctx).Workspace.UpdateUser(ctx, tid, uid, gqlmodel.FromRole(input.Role), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.UpdateMemberOfWorkspacePayload{Workspace: gqlmodel.ToWorkspace(res)}, nil
}

func (r *mutationResolver) UpdateIntegrationOfWorkspace(ctx context.Context, input gqlmodel.UpdateIntegrationOfWorkspaceInput) (*gqlmodel.UpdateMemberOfWorkspacePayload, error) {
	wId, iId, err := gqlmodel.ToID2[id.Workspace, id.Integration](input.WorkspaceID, input.IntegrationID)
	if err != nil {
		return nil, err
	}

	res, err := usecases(ctx).Workspace.UpdateIntegration(ctx, wId, iId, gqlmodel.FromRole(input.Role), getOperator(ctx))
	if err != nil {
		return nil, err
	}

	return &gqlmodel.UpdateMemberOfWorkspacePayload{Workspace: gqlmodel.ToWorkspace(res)}, nil
}
