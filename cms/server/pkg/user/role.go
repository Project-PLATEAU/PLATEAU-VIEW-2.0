package user

import (
	"strings"

	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var (
	// RoleOwner is a role who can have full controll of project
	RoleOwner = Role("owner")
	// RoleMaintainer is a role who can maintain a project
	RoleMaintainer = Role("maintainer")
	// RoleWriter is a role who can read and write project
	RoleWriter = Role("writer")
	// RoleReader is a role who can read project
	RoleReader = Role("reader")

	roles = []Role{
		RoleOwner,
		RoleMaintainer,
		RoleWriter,
		RoleReader,
	}

	ErrInvalidRole = rerror.NewE(i18n.T("invalid role"))
)

type Role string

func (r Role) Valid() bool {
	switch r {
	case RoleOwner:
		return true
	case RoleMaintainer:
		return true
	case RoleWriter:
		return true
	case RoleReader:
		return true
	}
	return false
}

func RoleFromString(r string) (Role, error) {
	role := Role(strings.ToLower(r))

	if role.Valid() {
		return role, nil
	}
	return role, ErrInvalidRole
}

func (r Role) Includes(role Role) bool {
	for i, r2 := range roles {
		if r == r2 {
			for _, r3 := range roles[i:] {
				if role == r3 {
					return true
				}
			}
		}
	}
	return false
}
