package repo

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"
)

type Workspace interface {
	FindByID(context.Context, id.WorkspaceID) (*user.Workspace, error)
	FindByIDs(context.Context, id.WorkspaceIDList) (user.WorkspaceList, error)
	FindByUser(context.Context, id.UserID) (user.WorkspaceList, error)
	FindByIntegration(context.Context, id.IntegrationID) (user.WorkspaceList, error)
	Save(context.Context, *user.Workspace) error
	SaveAll(context.Context, []*user.Workspace) error
	Remove(context.Context, id.WorkspaceID) error
	RemoveAll(context.Context, id.WorkspaceIDList) error
}
