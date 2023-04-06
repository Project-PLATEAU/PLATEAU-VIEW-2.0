package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_SortByID(t *testing.T) {
	id1 := NewID()
	id2 := NewID()

	list := List{
		&Model{id: id2},
		&Model{id: id1},
	}
	res := list.SortByID()
	assert.Equal(t, List{
		&Model{id: id1},
		&Model{id: id2},
	}, res)
	// test whether original list is not modified
	assert.Equal(t, List{
		&Model{id: id2},
		&Model{id: id1},
	}, list)
}

func TestList_Clone(t *testing.T) {
	id := NewID()
	list := List{&Model{id: id}}
	got := list.Clone()
	assert.Equal(t, list, got)
	assert.NotSame(t, list[0], got[0])

	got[0].id = NewID()
	// test whether original list is not modified
	assert.Equal(t, list, List{&Model{id: id}})
}
