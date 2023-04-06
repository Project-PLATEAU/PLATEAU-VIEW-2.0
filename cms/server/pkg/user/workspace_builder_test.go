package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkspaceBuilder_ID(t *testing.T) {
	tid := NewWorkspaceID()
	tm := NewWorkspace().ID(tid).MustBuild()
	assert.Equal(t, tid, tm.ID())
}

func TestWorkspaceBuilder_Members(t *testing.T) {
	m := map[ID]Member{NewID(): {Role: RoleOwner}}
	tm := NewWorkspace().NewID().Members(m).MustBuild()
	assert.Equal(t, m, tm.Members().Users())
}

func TestWorkspaceBuilder_Personal(t *testing.T) {
	tm := NewWorkspace().NewID().Personal(true).MustBuild()
	assert.True(t, tm.IsPersonal())
}

func TestWorkspaceBuilder_Name(t *testing.T) {
	tm := NewWorkspace().NewID().Name("xxx").MustBuild()
	assert.Equal(t, "xxx", tm.Name())
}

func TestWorkspaceBuilder_NewID(t *testing.T) {
	tm := NewWorkspace().NewID().MustBuild()
	assert.NotNil(t, tm.ID())
}

func TestWorkspaceBuilder_Build(t *testing.T) {
	tid := NewWorkspaceID()
	uid := NewID()

	type args struct {
		ID       WorkspaceID
		Name     string
		Personal bool
		Members  map[ID]Member
	}

	tests := []struct {
		Name     string
		Args     args
		Expected *Workspace
		Err      error
	}{
		{
			Name: "success create workspace",
			Args: args{
				ID:       tid,
				Name:     "xxx",
				Personal: true,
				Members:  map[ID]Member{uid: {Role: RoleOwner}},
			},
			Expected: &Workspace{
				id:   tid,
				name: "xxx",
				members: &Members{
					users:        map[ID]Member{uid: {Role: RoleOwner}},
					integrations: map[IntegrationID]Member{},
					fixed:        true,
				},
			},
		}, {
			Name: "success create workspace with nil users",
			Args: args{
				ID:   tid,
				Name: "xxx",
			},
			Expected: &Workspace{
				id:   tid,
				name: "xxx",
				members: &Members{
					users:        map[ID]Member{},
					integrations: map[IntegrationID]Member{},
					fixed:        false,
				},
			},
		},
		{
			Name: "fail invalid id",
			Err:  ErrInvalidID,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			res, err := NewWorkspace().
				ID(tt.Args.ID).
				Members(tt.Args.Members).
				Personal(tt.Args.Personal).
				Name(tt.Args.Name).
				Build()
			if tt.Err == nil {
				assert.Equal(t, tt.Expected, res)
			} else {
				assert.Equal(t, tt.Err, err)
			}
		})
	}
}

func TestWorkspaceBuilder_MustBuild(t *testing.T) {
	tid := NewWorkspaceID()
	uid := NewID()

	type args struct {
		ID       WorkspaceID
		Name     string
		Personal bool
		Members  map[ID]Member
	}

	tests := []struct {
		Name     string
		Args     args
		Expected *Workspace
		Err      error
	}{
		{
			Name: "success create workspace",
			Args: args{
				ID:       tid,
				Name:     "xxx",
				Personal: true,
				Members:  map[ID]Member{uid: {Role: RoleOwner}},
			},
			Expected: &Workspace{
				id:   tid,
				name: "xxx",
				members: &Members{
					users:        map[ID]Member{uid: {Role: RoleOwner}},
					integrations: map[IntegrationID]Member{},
					fixed:        true,
				},
			},
		}, {
			Name: "success create workspace with nil users",
			Args: args{
				ID:   tid,
				Name: "xxx",
			},
			Expected: &Workspace{
				id:   tid,
				name: "xxx",
				members: &Members{
					users:        map[ID]Member{},
					integrations: map[IntegrationID]Member{},
					fixed:        false,
				},
			},
		},
		{
			Name: "fail invalid id",
			Err:  ErrInvalidID,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			build := func() *Workspace {
				t.Helper()
				return NewWorkspace().ID(tt.Args.ID).Members(tt.Args.Members).Personal(tt.Args.Personal).Name(tt.Args.Name).MustBuild()
			}

			if tt.Err != nil {
				assert.PanicsWithValue(t, tt.Err, func() { _ = build() })
			} else {
				assert.Equal(t, tt.Expected, build())
			}
		})
	}
}
