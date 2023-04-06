package request

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	vor := version.New().OrRef()
	iid := id.NewItemID()
	itm, err := NewItemWithVersion(iid, vor)
	assert.NoError(t, err)
	assert.Equal(t, vor, itm.Pointer())
	assert.Equal(t, iid, itm.Item())
}

func TestItemList_IDs(t *testing.T) {
	i1, _ := NewItem(id.NewItemID())
	i2, _ := NewItem(id.NewItemID())
	i3, _ := NewItem(id.NewItemID())
	list := ItemList{i1, i2, i3}
	ids := id.ItemIDList{i1.Item(), i2.Item(), i3.Item()}
	assert.Equal(t, ids, list.IDs())
}

func TestItemList_HasDuplication(t *testing.T) {
	i1, _ := NewItem(id.NewItemID())
	i2, _ := NewItem(id.NewItemID())
	i3, _ := NewItem(id.NewItemID())
	input1 := ItemList{i1, i2, i3, i2}
	input2 := ItemList{i1, i2, i3}
	assert.Equal(t, true, input1.HasDuplication())
	assert.Equal(t, false, input2.HasDuplication())
}
