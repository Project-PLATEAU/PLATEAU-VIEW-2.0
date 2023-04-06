package gqlmodel

import (
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/util"
)

func ToUser(u *user.User) *User {
	if u == nil {
		return nil
	}

	return &User{
		ID:    IDFrom(u.ID()),
		Name:  u.Name(),
		Email: u.Email(),
	}
}

func ToMe(u *user.User) *Me {
	if u == nil {
		return nil
	}

	return &Me{
		ID:            IDFrom(u.ID()),
		Name:          u.Name(),
		Email:         u.Email(),
		Lang:          u.Lang(),
		Theme:         Theme(u.Theme()),
		MyWorkspaceID: IDFrom(u.Workspace()),
		Auths: util.Map(u.Auths(), func(a user.Auth) string {
			return a.Provider
		}),
	}
}

func ToTheme(t *Theme) *user.Theme {
	if t == nil {
		return nil
	}

	th := user.ThemeDefault
	switch *t {
	case ThemeDark:
		th = user.ThemeDark
	case ThemeLight:
		th = user.ThemeLight
	}
	return &th
}

func ToWorkspace(t *user.Workspace) *Workspace {
	if t == nil {
		return nil
	}

	usersMap := t.Members().Users()
	integrationsMap := t.Members().Integrations()
	members := make([]WorkspaceMember, 0, len(usersMap)+len(integrationsMap))
	for u, m := range usersMap {
		members = append(members, &WorkspaceUserMember{
			UserID: IDFrom(u),
			Role:   ToRole(m.Role),
		})
	}
	for i, m := range integrationsMap {
		members = append(members, &WorkspaceIntegrationMember{
			IntegrationID: IDFrom(i),
			Role:          ToRole(m.Role),
			Active:        !m.Disabled,
			InvitedByID:   IDFrom(m.InvitedBy),
			InvitedBy:     nil,
			Integration:   nil,
		})
	}

	return &Workspace{
		ID:       IDFrom(t.ID()),
		Name:     t.Name(),
		Personal: t.IsPersonal(),
		Members:  members,
	}
}

func FromRole(r Role) user.Role {
	switch r {
	case RoleReader:
		return user.RoleReader
	case RoleWriter:
		return user.RoleWriter
	case RoleMaintainer:
		return user.RoleMaintainer
	case RoleOwner:
		return user.RoleOwner
	}
	return user.Role("")
}

func ToRole(r user.Role) Role {
	switch r {
	case user.RoleReader:
		return RoleReader
	case user.RoleWriter:
		return RoleWriter
	case user.RoleMaintainer:
		return RoleMaintainer
	case user.RoleOwner:
		return RoleOwner
	}
	return Role("")
}
