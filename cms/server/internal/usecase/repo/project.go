package repo

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearthx/usecasex"
)

type Project interface {
	Filtered(filter WorkspaceFilter) Project
	FindByIDs(context.Context, id.ProjectIDList) (project.List, error)
	FindByID(context.Context, id.ProjectID) (*project.Project, error)
	FindByIDOrAlias(context.Context, project.IDOrAlias) (*project.Project, error)
	FindByWorkspaces(context.Context, id.WorkspaceIDList, *usecasex.Pagination) (project.List, *usecasex.PageInfo, error)
	FindByPublicName(context.Context, string) (*project.Project, error)
	CountByWorkspace(context.Context, id.WorkspaceID) (int, error)
	Save(context.Context, *project.Project) error
	Remove(context.Context, id.ProjectID) error
}
