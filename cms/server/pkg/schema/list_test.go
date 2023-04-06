package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/stretchr/testify/assert"
)

func TestFieldList_Find(t *testing.T) {
	f := NewField(NewText(nil).TypeProperty()).NewID().Key(key.Random()).MustBuild()
	f2 := NewField(NewText(nil).TypeProperty()).NewID().Key(key.Random()).MustBuild()
	assert.Equal(t, f, FieldList{f, f2}.Find(f.ID()))
	assert.Equal(t, f2, FieldList{f, f2}.Find(f2.ID()))
	assert.Nil(t, FieldList{f, f2}.Find(NewFieldID()))
}

func TestList_SortByID(t *testing.T) {
	id1 := NewID()
	id2 := NewID()

	list := List{
		&Schema{id: id2},
		&Schema{id: id1},
	}
	res := list.SortByID()
	assert.Equal(t, List{
		&Schema{id: id1},
		&Schema{id: id2},
	}, res)
	// test whether original list is not modified
	assert.Equal(t, List{
		&Schema{id: id2},
		&Schema{id: id1},
	}, list)
}

func TestList_Clone(t *testing.T) {
	id := NewID()
	list := List{&Schema{id: id}}
	got := list.Clone()
	assert.Equal(t, list, got)
	assert.NotSame(t, list[0], got[0])

	got[0].id = NewID()
	// test whether original list is not modified
	assert.Equal(t, list, List{&Schema{id: id}})
}

func TestFieldList_Clone(t *testing.T) {
	id := NewFieldID()
	list := FieldList{&Field{id: id}}
	got := list.Clone()
	assert.Equal(t, list, got)
	assert.NotSame(t, list[0], got[0])

	got[0].id = NewFieldID()
	// test whether original list is not modified
	assert.Equal(t, list, FieldList{&Field{id: id}})
}

func TestFieldList_SortByID(t *testing.T) {
	id1 := NewFieldID()
	id2 := NewFieldID()

	list := FieldList{
		&Field{id: id2},
		&Field{id: id1},
	}
	res := list.SortByID()
	assert.Equal(t, FieldList{
		&Field{id: id1},
		&Field{id: id2},
	}, res)
	// test whether original list is not modified
	assert.Equal(t, FieldList{
		&Field{id: id2},
		&Field{id: id1},
	}, list)
}

func TestFieldList_IDs(t *testing.T) {
	id1 := NewFieldID()
	id2 := NewFieldID()

	list := FieldList{
		&Field{id: id1},
		&Field{id: id2},
	}
	assert.Equal(t, id.FieldIDList{id1, id2}, list.IDs())
}

func TestFieldList_Ordered(t *testing.T) {
	id1 := NewFieldID()
	id2 := NewFieldID()

	list := FieldList{
		&Field{id: id2, order: 2},
		&Field{id: id1, order: 1},
	}
	res := list.Ordered()
	assert.Equal(t, FieldList{
		&Field{id: id1, order: 1},
		&Field{id: id2, order: 2},
	}, res)
}
