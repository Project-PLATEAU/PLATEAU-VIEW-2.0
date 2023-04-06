package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkspace_ID(t *testing.T) {
	tid := NewWorkspaceID()
	tm := NewWorkspace().ID(tid).MustBuild()
	assert.Equal(t, tid, tm.ID())
}

func TestWorkspace_Name(t *testing.T) {
	tm := NewWorkspace().NewID().Name("ttt").MustBuild()
	assert.Equal(t, "ttt", tm.Name())
}

func TestWorkspace_Members(t *testing.T) {
	m := map[ID]Member{
		NewID(): {Role: RoleOwner},
	}
	tm := NewWorkspace().NewID().Members(m).MustBuild()
	assert.Equal(t, m, tm.Members().Users())
}

func TestWorkspace_IsPersonal(t *testing.T) {
	tm := NewWorkspace().NewID().Personal(true).MustBuild()
	assert.Equal(t, true, tm.IsPersonal())
}

func TestWorkspace_Rename(t *testing.T) {
	tm := NewWorkspace().NewID().Name("ttt").MustBuild()
	tm.Rename("ccc")
	assert.Equal(t, "ccc", tm.Name())
}
