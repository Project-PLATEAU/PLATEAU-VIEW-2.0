package usecase

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/user"
)

type Operator struct {
	User                   *user.ID
	Integration            *integration.ID
	Machine                bool
	Lang                   string
	ReadableWorkspaces     user.WorkspaceIDList
	WritableWorkspaces     user.WorkspaceIDList
	MaintainableWorkspaces user.WorkspaceIDList
	OwningWorkspaces       user.WorkspaceIDList
	ReadableProjects       project.IDList
	WritableProjects       project.IDList
	OwningProjects         project.IDList
	MaintainableProjects   project.IDList
}

type Ownable interface {
	User() *id.UserID
	Integration() *id.IntegrationID
	Project() id.ProjectID
}

func (o *Operator) Workspaces(r user.Role) []id.WorkspaceID {
	if o == nil {
		return nil
	}
	if r == user.RoleReader {
		return o.ReadableWorkspaces
	}
	if r == user.RoleWriter {
		return o.WritableWorkspaces
	}
	if r == user.RoleMaintainer {
		return o.MaintainableWorkspaces
	}
	if r == user.RoleOwner {
		return o.OwningWorkspaces
	}
	return nil
}

func (o *Operator) AllReadableWorkspaces() user.WorkspaceIDList {
	return append(o.ReadableWorkspaces, o.AllWritableWorkspaces()...)
}

func (o *Operator) AllWritableWorkspaces() user.WorkspaceIDList {
	return append(o.WritableWorkspaces, o.AllMaintainingWorkspaces()...)
}

func (o *Operator) AllMaintainingWorkspaces() user.WorkspaceIDList {
	return append(o.MaintainableWorkspaces, o.AllOwningWorkspaces()...)
}

func (o *Operator) AllOwningWorkspaces() user.WorkspaceIDList {
	return o.OwningWorkspaces
}

func (o *Operator) IsReadableWorkspace(workspace ...id.WorkspaceID) bool {
	return o.AllReadableWorkspaces().Intersect(workspace).Len() > 0
}

func (o *Operator) IsWritableWorkspace(workspace ...id.WorkspaceID) bool {
	return o.AllWritableWorkspaces().Intersect(workspace).Len() > 0
}

func (o *Operator) IsMaintainingWorkspace(workspace ...id.WorkspaceID) bool {
	return o.AllMaintainingWorkspaces().Intersect(workspace).Len() > 0
}

func (o *Operator) IsOwningWorkspace(workspace ...id.WorkspaceID) bool {
	return o.AllOwningWorkspaces().Intersect(workspace).Len() > 0
}

func (o *Operator) AddNewWorkspace(workspace id.WorkspaceID) {
	o.OwningWorkspaces = append(o.OwningWorkspaces, workspace)
}

func (o *Operator) Projects(r user.Role) project.IDList {
	if o == nil {
		return nil
	}
	if r == user.RoleReader {
		return o.ReadableProjects
	}
	if r == user.RoleWriter {
		return o.WritableProjects
	}
	if r == user.RoleMaintainer {
		return o.MaintainableProjects
	}
	if r == user.RoleOwner {
		return o.OwningProjects
	}
	return nil
}

func (o *Operator) AllReadableProjects() project.IDList {
	return append(o.ReadableProjects, o.AllWritableProjects()...)
}

func (o *Operator) AllWritableProjects() project.IDList {
	return append(o.WritableProjects, o.AllMaintainableProjects()...)
}

func (o *Operator) AllMaintainableProjects() project.IDList {
	return append(o.MaintainableProjects, o.AllOwningProjects()...)
}

func (o *Operator) AllOwningProjects() project.IDList {
	return o.OwningProjects
}

func (o *Operator) IsReadableProject(projects ...project.ID) bool {
	return o.AllReadableProjects().Intersect(projects).Len() > 0
}

func (o *Operator) IsWritableProject(projects ...project.ID) bool {
	return o.AllWritableProjects().Intersect(projects).Len() > 0
}

func (o *Operator) IsMaintainingProject(projects ...project.ID) bool {
	return o.AllMaintainableProjects().Intersect(projects).Len() > 0
}

func (o *Operator) IsOwningProject(projects ...project.ID) bool {
	return o.AllOwningProjects().Intersect(projects).Len() > 0
}

func (o *Operator) AddNewProject(p project.ID) {
	o.OwningProjects = append(o.OwningProjects, p)
}

func (o *Operator) Operator() operator.Operator {
	var eOp operator.Operator
	if o.User != nil {
		eOp = operator.OperatorFromUser(*o.User)
	}
	if o.Integration != nil {
		eOp = operator.OperatorFromIntegration(*o.Integration)
	}
	if o.Machine {
		eOp = operator.OperatorFromMachine()
	}
	return eOp
}

func (o *Operator) CanUpdate(obj Ownable) bool {
	isWriter := o.IsWritableProject(obj.Project())
	isMaintainer := o.IsMaintainingProject(obj.Project())
	return isMaintainer || (isWriter && o.Owns(obj)) || o.Machine
}

func (o *Operator) Owns(obj Ownable) bool {
	return (o.User != nil && obj.User() != nil && *o.User == *obj.User()) ||
		(o.Integration != nil && obj.Integration() != nil && *o.Integration == *obj.Integration())
}
