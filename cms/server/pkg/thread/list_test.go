package thread

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestList_SortByID(t *testing.T) {
	id1 := NewID()
	id2 := NewID()

	list := List{
		&Thread{id: id2},
		&Thread{id: id1},
	}
	res := list.SortByID()
	assert.Equal(t, List{
		&Thread{id: id1},
		&Thread{id: id2},
	}, res)
	// test whether original list is not modified
	assert.Equal(t, List{
		&Thread{id: id2},
		&Thread{id: id1},
	}, list)
}

func TestList_Clone(t *testing.T) {
	th := New().NewID().Workspace(id.NewWorkspaceID()).MustBuild()

	list := List{th}
	got := list.Clone()
	assert.Equal(t, list, got)
	assert.NotSame(t, list[0], got[0])
}
