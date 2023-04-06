package interfaces

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
)

type CreateProjectParam struct {
	WorkspaceID id.WorkspaceID
	Name        *string
	Description *string
	Alias       *string
}

type UpdateProjectParam struct {
	ID          id.ProjectID
	Name        *string
	Description *string
	Alias       *string
	Publication *UpdateProjectPublicationParam
}

type UpdateProjectPublicationParam struct {
	Scope       *project.PublicationScope
	AssetPublic *bool
}

var (
	ErrProjectAliasIsNotSet    error = rerror.NewE(i18n.T("project alias is not set"))
	ErrProjectAliasAlreadyUsed error = rerror.NewE(i18n.T("project alias is already used by another project"))
)

type Project interface {
	Fetch(context.Context, []id.ProjectID, *usecase.Operator) (project.List, error)
	FindByIDOrAlias(context.Context, project.IDOrAlias, *usecase.Operator) (*project.Project, error)
	FindByWorkspace(context.Context, id.WorkspaceID, *usecasex.Pagination, *usecase.Operator) (project.List, *usecasex.PageInfo, error)
	Create(context.Context, CreateProjectParam, *usecase.Operator) (*project.Project, error)
	Update(context.Context, UpdateProjectParam, *usecase.Operator) (*project.Project, error)
	CheckAlias(context.Context, string) (bool, error)
	Delete(context.Context, id.ProjectID, *usecase.Operator) error
}
