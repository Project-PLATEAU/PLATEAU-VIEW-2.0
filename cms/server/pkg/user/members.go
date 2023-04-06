package user

import (
	"sort"

	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

var (
	ErrUserAlreadyJoined             = rerror.NewE(i18n.T("user already joined"))
	ErrCannotModifyPersonalWorkspace = rerror.NewE(i18n.T("personal workspace cannot be modified"))
	ErrTargetUserNotInTheWorkspace   = rerror.NewE(i18n.T("target user does not exist in the workspace"))
	ErrInvalidWorkspaceName          = rerror.NewE(i18n.T("invalid workspace name"))
)

type Member struct {
	Role      Role
	Disabled  bool
	InvitedBy ID
}

type Members struct {
	users        map[ID]Member
	integrations map[IntegrationID]Member
	fixed        bool
}

func NewMembers() *Members {
	m := &Members{
		users:        map[ID]Member{},
		integrations: map[IntegrationID]Member{},
	}
	return m
}

func NewFixedMembers(u ID) *Members {
	m := &Members{
		users: map[ID]Member{
			u: {
				Role:      RoleOwner,
				Disabled:  false,
				InvitedBy: u,
			},
		},
		integrations: map[IntegrationID]Member{},
		fixed:        true,
	}
	return m
}

func NewMembersWith(users map[ID]Member) *Members {
	m := &Members{
		users:        maps.Clone(users),
		integrations: map[IntegrationID]Member{},
	}
	return m
}

func NewFixedMembersWith(users map[ID]Member) *Members {
	m := &Members{
		users:        maps.Clone(users),
		integrations: map[IntegrationID]Member{},
		fixed:        true,
	}
	return m
}

func (m *Members) Clone() *Members {
	c := &Members{
		users:        maps.Clone(m.users),
		integrations: maps.Clone(m.integrations),
		fixed:        m.fixed,
	}
	return c
}

func (m *Members) Users() map[ID]Member {
	return maps.Clone(m.users)
}

func (m *Members) Integrations() map[IntegrationID]Member {
	return maps.Clone(m.integrations)
}

func (m *Members) IntegrationIDs() IntegrationIDList {
	return IntegrationIDList(lo.Keys(m.integrations)).Sort()
}

func (m *Members) HasUser(u ID) bool {
	_, ok := m.users[u]
	return ok
}

func (m *Members) HasIntegration(i IntegrationID) bool {
	_, ok := m.integrations[i]
	return ok
}

func (m *Members) Count() int {
	return len(m.users)
}

func (m *Members) UserRole(u ID) Role {
	return m.users[u].Role
}

func (m *Members) IntegrationRole(iId IntegrationID) Role {
	return m.integrations[iId].Role
}

func (m *Members) UpdateUserRole(u ID, role Role) error {
	if m.fixed {
		return ErrCannotModifyPersonalWorkspace
	}
	if role == Role("") {
		return nil
	}
	if _, ok := m.users[u]; !ok {
		return ErrTargetUserNotInTheWorkspace
	}
	mm := m.users[u]
	mm.Role = role
	m.users[u] = mm
	return nil
}

func (m *Members) UpdateIntegrationRole(iId IntegrationID, role Role) error {
	if !role.Valid() {
		return nil
	}
	if _, ok := m.integrations[iId]; !ok {
		return ErrTargetUserNotInTheWorkspace
	}
	mm := m.integrations[iId]
	mm.Role = role
	m.integrations[iId] = mm
	return nil
}

func (m *Members) JoinUser(u ID, role Role, i ID) error {
	if m.fixed {
		return ErrCannotModifyPersonalWorkspace
	}
	if _, ok := m.users[u]; ok {
		return ErrUserAlreadyJoined
	}
	if role == Role("") {
		role = RoleReader
	}
	m.users[u] = Member{
		Role:      role,
		Disabled:  false,
		InvitedBy: i,
	}
	return nil
}

func (m *Members) AddIntegration(iId IntegrationID, role Role, i ID) error {
	if _, ok := m.integrations[iId]; ok {
		return ErrUserAlreadyJoined
	}
	if role == Role("") {
		role = RoleReader
	}
	m.integrations[iId] = Member{
		Role:      role,
		Disabled:  false,
		InvitedBy: i,
	}
	return nil
}

func (m *Members) Leave(u ID) error {
	if m.fixed {
		return ErrCannotModifyPersonalWorkspace
	}
	if _, ok := m.users[u]; ok {
		delete(m.users, u)
	} else {
		return ErrTargetUserNotInTheWorkspace
	}
	return nil
}

func (m *Members) DeleteIntegration(iId IntegrationID) error {
	if _, ok := m.integrations[iId]; ok {
		delete(m.integrations, iId)
	} else {
		return ErrTargetUserNotInTheWorkspace
	}
	return nil
}

func (m *Members) UsersByRole(role Role) []ID {
	users := make([]ID, 0, len(m.users))
	for u, m := range m.users {
		if m.Role == role {
			users = append(users, u)
		}
	}

	sort.SliceStable(users, func(a, b int) bool {
		return users[a].Compare(users[b]) > 0
	})

	return users
}

func (m *Members) IsOnlyOwner(u ID) bool {
	return len(m.UsersByRole(RoleOwner)) == 1 && m.users[u].Role == RoleOwner
}

func (m *Members) IsOwnerOrMaintainer(u ID) bool {
	return m.users[u].Role == RoleOwner || m.users[u].Role == RoleMaintainer
}

func (m *Members) Fixed() bool {
	if m == nil {
		return false
	}
	return m.fixed
}
