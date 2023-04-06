package interactor

import (
	"context"
	"strings"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/usecasex"
	"golang.org/x/exp/maps"
)

type Workspace struct {
	repos       *repo.Container
	gateways    *gateway.Container
	transaction usecasex.Transaction
}

func NewWorkspace(r *repo.Container, g *gateway.Container) interfaces.Workspace {
	return &Workspace{
		repos:       r,
		gateways:    g,
		transaction: r.Transaction,
	}
}

func (i *Workspace) Fetch(ctx context.Context, ids []id.WorkspaceID, operator *usecase.Operator) ([]*user.Workspace, error) {
	return Run1(ctx, operator, i.repos, Usecase().Transaction(), func(ctx context.Context) ([]*user.Workspace, error) {
		res, err := i.repos.Workspace.FindByIDs(ctx, ids)
		res2, err := i.filterWorkspaces(res, operator, err)
		return res2, err
	})
}

func (i *Workspace) FindByUser(ctx context.Context, id id.UserID, operator *usecase.Operator) ([]*user.Workspace, error) {
	return Run1(ctx, operator, i.repos, Usecase().Transaction(), func(ctx context.Context) ([]*user.Workspace, error) {
		res, err := i.repos.Workspace.FindByUser(ctx, id)
		res2, err := i.filterWorkspaces(res, operator, err)
		return res2, err
	})
}

func (i *Workspace) Create(ctx context.Context, name string, firstUser id.UserID, operator *usecase.Operator) (_ *user.Workspace, err error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(), func(ctx context.Context) (*user.Workspace, error) {
		if len(strings.TrimSpace(name)) == 0 {
			return nil, user.ErrInvalidName
		}

		workspace, err := user.NewWorkspace().
			NewID().
			Name(name).
			Build()
		if err != nil {
			return nil, err
		}

		if err := workspace.Members().JoinUser(firstUser, user.RoleOwner, *operator.User); err != nil {
			return nil, err
		}

		if err := i.repos.Workspace.Save(ctx, workspace); err != nil {
			return nil, err
		}

		operator.AddNewWorkspace(workspace.ID())
		return workspace, nil
	})
}

func (i *Workspace) Update(ctx context.Context, id id.WorkspaceID, name string, operator *usecase.Operator) (_ *user.Workspace, err error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(), func(ctx context.Context) (*user.Workspace, error) {
		workspace, err := i.repos.Workspace.FindByID(ctx, id)
		if err != nil {
			return nil, err
		}
		if workspace.IsPersonal() {
			return nil, user.ErrCannotModifyPersonalWorkspace
		}
		if workspace.Members().UserRole(*operator.User) != user.RoleOwner {
			return nil, interfaces.ErrOperationDenied
		}

		if len(strings.TrimSpace(name)) == 0 {
			return nil, user.ErrInvalidName
		}

		workspace.Rename(name)

		err = i.repos.Workspace.Save(ctx, workspace)
		if err != nil {
			return nil, err
		}

		return workspace, nil
	})
}

func (i *Workspace) AddUserMember(ctx context.Context, workspaceID id.WorkspaceID, users map[id.UserID]user.Role, operator *usecase.Operator) (_ *user.Workspace, err error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction().WithOwnableWorkspaces(workspaceID), func(ctx context.Context) (*user.Workspace, error) {
		workspace, err := i.repos.Workspace.FindByID(ctx, workspaceID)
		if err != nil {
			return nil, err
		}
		if workspace.IsPersonal() {
			return nil, user.ErrCannotModifyPersonalWorkspace
		}

		ul, err := i.repos.User.FindByIDs(ctx, maps.Keys(users))
		if err != nil {
			return nil, err
		}

		for _, m := range ul {
			err = workspace.Members().JoinUser(m.ID(), users[m.ID()], *operator.User)
			if err != nil {
				return nil, err
			}
		}

		err = i.repos.Workspace.Save(ctx, workspace)
		if err != nil {
			return nil, err
		}

		return workspace, nil
	})
}

func (i *Workspace) AddIntegrationMember(ctx context.Context, wId id.WorkspaceID, iId id.IntegrationID, role user.Role, operator *usecase.Operator) (_ *user.Workspace, err error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction().WithOwnableWorkspaces(wId), func(ctx context.Context) (*user.Workspace, error) {
		workspace, err := i.repos.Workspace.FindByID(ctx, wId)
		if err != nil {
			return nil, err
		}

		_, err = i.repos.Integration.FindByID(ctx, iId)
		if err != nil {
			return nil, err
		}

		err = workspace.Members().AddIntegration(iId, role, *operator.User)
		if err != nil {
			return nil, err
		}

		err = i.repos.Workspace.Save(ctx, workspace)
		if err != nil {
			return nil, err
		}

		return workspace, nil
	})
}

func (i *Workspace) RemoveUser(ctx context.Context, id id.WorkspaceID, u id.UserID, operator *usecase.Operator) (_ *user.Workspace, err error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction(), func(ctx context.Context) (*user.Workspace, error) {
		workspace, err := i.repos.Workspace.FindByID(ctx, id)
		if err != nil {
			return nil, err
		}
		if workspace.IsPersonal() {
			return nil, user.ErrCannotModifyPersonalWorkspace
		}

		isOwner := workspace.Members().UserRole(*operator.User) == user.RoleOwner
		isSelfLeave := *operator.User == u
		if !isOwner && !isSelfLeave {
			return nil, interfaces.ErrOperationDenied
		}

		if isSelfLeave && workspace.Members().IsOnlyOwner(u) {
			return nil, interfaces.ErrOwnerCannotLeaveTheWorkspace
		}

		err = workspace.Members().Leave(u)
		if err != nil {
			return nil, err
		}

		err = i.repos.Workspace.Save(ctx, workspace)
		if err != nil {
			return nil, err
		}

		return workspace, nil
	})
}

func (i *Workspace) RemoveIntegration(ctx context.Context, wId id.WorkspaceID, iId id.IntegrationID, operator *usecase.Operator) (_ *user.Workspace, err error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().WithOwnableWorkspaces(wId).Transaction(), func(ctx context.Context) (*user.Workspace, error) {
		workspace, err := i.repos.Workspace.FindByID(ctx, wId)
		if err != nil {
			return nil, err
		}

		err = workspace.Members().DeleteIntegration(iId)
		if err != nil {
			return nil, err
		}

		err = i.repos.Workspace.Save(ctx, workspace)
		if err != nil {
			return nil, err
		}

		return workspace, nil
	})
}

func (i *Workspace) UpdateUser(ctx context.Context, id id.WorkspaceID, u id.UserID, role user.Role, operator *usecase.Operator) (_ *user.Workspace, err error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().Transaction().WithOwnableWorkspaces(id), func(ctx context.Context) (*user.Workspace, error) {
		workspace, err := i.repos.Workspace.FindByID(ctx, id)
		if err != nil {
			return nil, err
		}
		if workspace.IsPersonal() {
			return nil, user.ErrCannotModifyPersonalWorkspace
		}

		if u == *operator.User {
			return nil, interfaces.ErrCannotChangeOwnerRole
		}

		err = workspace.Members().UpdateUserRole(u, role)
		if err != nil {
			return nil, err
		}

		err = i.repos.Workspace.Save(ctx, workspace)
		if err != nil {
			return nil, err
		}

		return workspace, nil
	})
}

func (i *Workspace) UpdateIntegration(ctx context.Context, wId id.WorkspaceID, iId id.IntegrationID, role user.Role, operator *usecase.Operator) (_ *user.Workspace, err error) {
	if operator.User == nil {
		return nil, interfaces.ErrInvalidOperator
	}
	return Run1(ctx, operator, i.repos, Usecase().WithOwnableWorkspaces(wId).Transaction(), func(ctx context.Context) (*user.Workspace, error) {
		workspace, err := i.repos.Workspace.FindByID(ctx, wId)
		if err != nil {
			return nil, err
		}

		err = workspace.Members().UpdateIntegrationRole(iId, role)
		if err != nil {
			return nil, err
		}

		err = i.repos.Workspace.Save(ctx, workspace)
		if err != nil {
			return nil, err
		}

		return workspace, nil
	})
}

func (i *Workspace) Remove(ctx context.Context, id id.WorkspaceID, operator *usecase.Operator) error {
	if operator.User == nil {
		return interfaces.ErrInvalidOperator
	}
	return Run0(ctx, operator, i.repos, Usecase().Transaction().WithOwnableWorkspaces(id), func(ctx context.Context) error {
		workspace, err := i.repos.Workspace.FindByID(ctx, id)
		if err != nil {
			return err
		}
		if workspace.IsPersonal() {
			return user.ErrCannotModifyPersonalWorkspace
		}

		projectCount, err := i.repos.Project.CountByWorkspace(ctx, id)
		if err != nil {
			return err
		}
		if projectCount > 0 {
			return interfaces.ErrWorkspaceWithProjects
		}

		err = i.repos.Workspace.Remove(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (i *Workspace) filterWorkspaces(workspaces []*user.Workspace, operator *usecase.Operator, err error) ([]*user.Workspace, error) {
	if err != nil {
		return nil, err
	}
	if operator == nil {
		return make([]*user.Workspace, len(workspaces)), nil
	}
	for i, t := range workspaces {
		if t == nil || !operator.IsReadableWorkspace(t.ID()) {
			workspaces[i] = nil
		}
	}
	return workspaces, nil
}
