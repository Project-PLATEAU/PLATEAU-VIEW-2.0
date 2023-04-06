package user

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewMembers(t *testing.T) {
	m := NewMembers()
	assert.NotNil(t, m)
	assert.IsType(t, &Members{}, m)
}

func TestNewMembersWith(t *testing.T) {
	uid := NewID()
	m := NewMembersWith(map[ID]Member{uid: {Role: RoleOwner}})
	assert.NotNil(t, m)
	assert.Equal(t, map[ID]Member{uid: {Role: RoleOwner}}, m.Users())
	assert.Equal(t, false, m.Fixed())
}

func TestNewFixedMembersWith(t *testing.T) {
	uid := NewID()
	m := NewFixedMembersWith(map[ID]Member{uid: {Role: RoleOwner}})
	assert.NotNil(t, m)
	assert.Equal(t, map[ID]Member{uid: {Role: RoleOwner}}, m.Users())
	assert.Equal(t, true, m.Fixed())
}

func TestMembers_HasUser(t *testing.T) {
	uid1 := NewID()
	uid2 := NewID()

	tests := []struct {
		Name     string
		M        *Members
		UID      ID
		Expected bool
	}{
		{
			Name:     "existing user",
			M:        NewMembersWith(map[ID]Member{uid1: {Role: RoleOwner}, uid2: {Role: RoleReader}}),
			UID:      uid1,
			Expected: true,
		},
		{
			Name:     "not existing user",
			M:        NewMembersWith(map[ID]Member{uid2: {Role: RoleReader}}),
			UID:      uid1,
			Expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res := tt.M.HasUser(tt.UID)
			assert.Equal(t, tt.Expected, res)
		})
	}
}

func TestMembers_HasIntegration(t *testing.T) {
	iId1 := id.NewIntegrationID()
	iId2 := id.NewIntegrationID()

	tests := []struct {
		Name     string
		M        *Members
		iId      IntegrationID
		Expected bool
	}{
		{
			Name: "existing integration",
			M: &Members{integrations: map[IntegrationID]Member{iId1: {
				Role:      RoleOwner,
				Disabled:  false,
				InvitedBy: ID{},
			}}},
			iId:      iId1,
			Expected: true,
		},
		{
			Name: "not existing user",
			M: &Members{integrations: map[IntegrationID]Member{iId1: {
				Role:      RoleOwner,
				Disabled:  false,
				InvitedBy: ID{},
			}}},
			iId:      iId2,
			Expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res := tt.M.HasIntegration(tt.iId)
			assert.Equal(t, tt.Expected, res)
		})
	}
}

func TestCopyMembers(t *testing.T) {
	uid := NewID()
	m := NewMembersWith(map[ID]Member{uid: {Role: RoleOwner}})
	m2 := m.Clone()
	assert.Equal(t, m, m2)
}

func TestMembers_Count(t *testing.T) {
	m := NewMembersWith(map[ID]Member{NewID(): {Role: RoleOwner}})
	assert.Equal(t, len(m.Users()), m.Count())
}

func TestMembers_GetUserRole(t *testing.T) {
	uid := NewID()
	m := NewMembersWith(map[ID]Member{uid: {Role: RoleOwner}})
	assert.Equal(t, RoleOwner, m.UserRole(uid))
}

func TestMembers_GetIntegrationRole(t *testing.T) {
	iId := id.NewIntegrationID()
	m := &Members{integrations: map[IntegrationID]Member{iId: {Role: RoleWriter}}}
	assert.Equal(t, RoleWriter, m.IntegrationRole(iId))
}

func TestMembers_IsOnlyOwner(t *testing.T) {
	uid := NewID()
	m := NewMembersWith(map[ID]Member{uid: {Role: RoleOwner}, NewID(): {Role: RoleReader}})
	assert.True(t, m.IsOnlyOwner(uid))
}

func TestMembers_Leave(t *testing.T) {
	uid := NewID()

	tests := []struct {
		Name string
		M    *Members
		UID  ID
		err  error
	}{
		{
			Name: "success user left",
			M:    NewMembersWith(map[ID]Member{uid: {Role: RoleWriter}, NewID(): {Role: RoleOwner}}),
			UID:  uid,
			err:  nil,
		},
		{
			Name: "fail personal workspace",
			M:    NewFixedMembers(uid),
			UID:  uid,
			err:  ErrCannotModifyPersonalWorkspace,
		},
		{
			Name: "fail user not in the workspace",
			M:    NewMembersWith(map[ID]Member{uid: {Role: RoleWriter}, NewID(): {Role: RoleOwner}}),
			UID:  NewID(),
			err:  ErrTargetUserNotInTheWorkspace,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			err := tt.M.Leave(tt.UID)
			if tt.err == nil {
				assert.False(t, tt.M.HasUser(tt.UID))
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func TestMembers_Members(t *testing.T) {
	uid := NewID()
	m := NewMembersWith(map[ID]Member{uid: {Role: RoleOwner}})
	assert.Equal(t, map[ID]Member{uid: {Role: RoleOwner}}, m.Users())
}

func TestMembers_UpdateUserRole(t *testing.T) {
	uid := NewID()

	tests := []struct {
		Name              string
		M                 *Members
		UID               ID
		NewRole, Expected Role
		err               error
	}{
		{
			Name:     "success role updated",
			M:        NewMembersWith(map[ID]Member{uid: {Role: RoleWriter}}),
			UID:      uid,
			NewRole:  RoleOwner,
			Expected: RoleOwner,
			err:      nil,
		},
		{
			Name:     "nil role",
			M:        NewMembersWith(map[ID]Member{uid: {Role: RoleOwner}}),
			UID:      uid,
			NewRole:  "",
			Expected: RoleOwner,
			err:      nil,
		},
		{
			Name:    "fail personal workspace",
			M:       NewFixedMembers(uid),
			UID:     uid,
			NewRole: RoleOwner,
			err:     ErrCannotModifyPersonalWorkspace,
		},
		{
			Name:    "fail user not in the workspace",
			M:       NewMembersWith(map[ID]Member{uid: {Role: RoleOwner}}),
			UID:     NewID(),
			NewRole: RoleOwner,
			err:     ErrTargetUserNotInTheWorkspace,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			err := tt.M.UpdateUserRole(tt.UID, tt.NewRole)
			if tt.err == nil {
				assert.Equal(t, tt.Expected, tt.M.UserRole(tt.UID))
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func TestMembers_UpdateIntegrationRole(t *testing.T) {
	iId := id.NewIntegrationID()

	tests := []struct {
		name          string
		m             *Members
		iId           IntegrationID
		newRole, want Role
		err           error
	}{
		{
			name:    "success role updated",
			m:       &Members{integrations: map[IntegrationID]Member{iId: {Role: RoleWriter}}},
			iId:     iId,
			newRole: RoleOwner,
			want:    RoleOwner,
			err:     nil,
		},
		{
			name:    "nil role",
			m:       &Members{integrations: map[IntegrationID]Member{iId: {Role: RoleWriter}}},
			iId:     iId,
			newRole: "",
			want:    RoleWriter,
			err:     nil,
		},
		{
			name:    "fail user not in the workspace",
			m:       &Members{integrations: map[IntegrationID]Member{iId: {Role: RoleWriter}}},
			iId:     id.NewIntegrationID(),
			newRole: RoleOwner,
			err:     ErrTargetUserNotInTheWorkspace,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.m.UpdateIntegrationRole(tt.iId, tt.newRole)
			if tt.err == nil {
				assert.Equal(t, tt.want, tt.m.IntegrationRole(tt.iId))
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func TestMembers_IntegrationIDs(t *testing.T) {
	i1 := integration.NewID()
	i2 := integration.NewID()
	u1 := NewID()
	m := NewMembersWith(map[ID]Member{u1: {Role: RoleOwner}})
	lo.Must0(m.AddIntegration(i1, RoleWriter, u1))
	lo.Must0(m.AddIntegration(i2, RoleWriter, u1))

	assert.Equal(t, IntegrationIDList{i1, i2}, m.IntegrationIDs())
}

func TestMembers_Join(t *testing.T) {
	uid := NewID()
	uid2 := NewID()

	tests := []struct {
		Name                   string
		M                      *Members
		UID                    ID
		JoinRole, ExpectedRole Role
		err                    error
	}{
		{
			Name:         "success join user",
			M:            NewMembersWith(map[ID]Member{uid: {Role: RoleWriter}}),
			UID:          uid2,
			JoinRole:     "xxx",
			ExpectedRole: "xxx",
			err:          nil,
		},
		{
			Name:         "success join user",
			M:            NewMembersWith(map[ID]Member{uid: {Role: RoleWriter}}),
			UID:          uid2,
			JoinRole:     "",
			ExpectedRole: RoleReader,
			err:          nil,
		},
		{
			Name:     "fail personal workspace",
			M:        NewFixedMembers(uid),
			UID:      uid2,
			JoinRole: "xxx",
			err:      ErrCannotModifyPersonalWorkspace,
		},
		{
			Name:     "fail user already joined",
			M:        NewMembersWith(map[ID]Member{uid: {Role: RoleOwner}}),
			UID:      uid,
			JoinRole: "",
			err:      ErrUserAlreadyJoined,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			err := tt.M.JoinUser(tt.UID, tt.JoinRole, NewID())
			if tt.err == nil {
				assert.True(t, tt.M.HasUser(tt.UID))
				assert.Equal(t, tt.ExpectedRole, tt.M.UserRole(tt.UID))
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func TestMembers_UsersByRole(t *testing.T) {
	uid := NewID()
	uid2 := NewID()

	tests := []struct {
		Name     string
		M        *Members
		Role     Role
		Expected []ID
		err      error
	}{
		{
			Name:     "success join user",
			M:        NewMembersWith(map[ID]Member{uid: {Role: "xxx"}, uid2: {Role: "xxx"}}),
			Role:     "xxx",
			Expected: []ID{uid2, uid},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res := tt.M.UsersByRole(tt.Role)
			assert.Equal(t, tt.Expected, res)
		})
	}
}

func TestMembers_Fixed(t *testing.T) {
	tests := []struct {
		name   string
		target *Members
		want   bool
	}{
		{
			name: "true",
			target: &Members{
				fixed: true,
			},
			want: true,
		},
		{
			name:   "empty",
			target: &Members{},
			want:   false,
		},
		{
			name: "nil",
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.target.Fixed())
		})
	}
}
