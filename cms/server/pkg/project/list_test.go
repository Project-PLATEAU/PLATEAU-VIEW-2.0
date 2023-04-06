package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_SortByID(t *testing.T) {
	id1 := NewID()
	id2 := NewID()

	list := List{
		&Project{id: id2},
		&Project{id: id1},
	}
	res := list.SortByID()
	assert.Equal(t, List{
		&Project{id: id1},
		&Project{id: id2},
	}, res)
	// test whether original list is not modified
	assert.Equal(t, List{
		&Project{id: id2},
		&Project{id: id1},
	}, list)
}

func TestList_Clone(t *testing.T) {
	p := New().NewID().Name("a").MustBuild()

	list := List{p}
	got := list.Clone()
	assert.Equal(t, list, got)
	assert.NotSame(t, list[0], got[0])
}
