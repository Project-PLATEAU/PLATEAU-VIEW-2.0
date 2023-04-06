package interfaces

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var (
	ErrOwnerCannotLeaveTheWorkspace = rerror.NewE(i18n.T("owner user cannot leave from the workspace"))
	ErrCannotChangeOwnerRole        = rerror.NewE(i18n.T("cannot change the role of the workspace owner"))
	ErrCannotDeleteWorkspace        = rerror.NewE(i18n.T("cannot delete workspace because at least one project is left"))
	ErrWorkspaceWithProjects        = rerror.NewE(i18n.T("target workspace still has some project"))
)

type Workspace interface {
	Fetch(context.Context, []id.WorkspaceID, *usecase.Operator) ([]*user.Workspace, error)
	FindByUser(context.Context, id.UserID, *usecase.Operator) ([]*user.Workspace, error)
	Create(context.Context, string, id.UserID, *usecase.Operator) (*user.Workspace, error)
	Update(context.Context, id.WorkspaceID, string, *usecase.Operator) (*user.Workspace, error)
	AddUserMember(context.Context, id.WorkspaceID, map[id.UserID]user.Role, *usecase.Operator) (*user.Workspace, error)
	AddIntegrationMember(context.Context, id.WorkspaceID, id.IntegrationID, user.Role, *usecase.Operator) (*user.Workspace, error)
	UpdateUser(context.Context, id.WorkspaceID, id.UserID, user.Role, *usecase.Operator) (*user.Workspace, error)
	UpdateIntegration(context.Context, id.WorkspaceID, id.IntegrationID, user.Role, *usecase.Operator) (*user.Workspace, error)
	RemoveUser(context.Context, id.WorkspaceID, id.UserID, *usecase.Operator) (*user.Workspace, error)
	RemoveIntegration(context.Context, id.WorkspaceID, id.IntegrationID, *usecase.Operator) (*user.Workspace, error)
	Remove(context.Context, id.WorkspaceID, *usecase.Operator) error
}
