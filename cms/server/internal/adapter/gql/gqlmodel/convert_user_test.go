package gqlmodel

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestToRole(t *testing.T) {
	tests := []struct {
		name string
		arg  user.Role
		want Role
	}{
		{
			name: "RoleOwner",
			arg:  user.RoleOwner,
			want: RoleOwner,
		},
		{
			name: "RoleMaintainer",
			arg:  user.RoleMaintainer,
			want: RoleMaintainer,
		},
		{
			name: "RoleWriter",
			arg:  user.RoleWriter,
			want: RoleWriter,
		},
		{
			name: "RoleReader",
			arg:  user.RoleReader,
			want: RoleReader,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, ToRole(tt.arg))
		})
	}
}

func TestFromRole(t *testing.T) {
	tests := []struct {
		name string
		arg  Role
		want user.Role
	}{
		{
			name: "RoleOwner",
			arg:  RoleOwner,
			want: user.RoleOwner,
		},
		{
			name: "RoleMaintainer",
			arg:  RoleMaintainer,
			want: user.RoleMaintainer,
		},
		{
			name: "RoleWriter",
			arg:  RoleWriter,
			want: user.RoleWriter,
		},
		{
			name: "RoleReader",
			arg:  RoleReader,
			want: user.RoleReader,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, FromRole(tt.arg))
		})
	}
}
