package request

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/version"
)

type ItemList []*Item

type Item struct {
	item    ItemID
	pointer version.VersionOrRef
}

func (i *Item) Item() ItemID {
	return i.item
}

func (i *Item) Pointer() version.VersionOrRef {
	return i.pointer
}

func NewItemWithVersion(i ItemID, v version.VersionOrRef) (*Item, error) {
	if i.IsNil() {
		return nil, ErrInvalidID
	}
	return &Item{
		item:    i,
		pointer: v,
	}, nil
}

func NewItem(i ItemID) (*Item, error) {
	ptr := version.Public.OrVersion()
	return NewItemWithVersion(i, ptr)
}

func (l ItemList) IDs() id.ItemIDList {
	ids := id.ItemIDList{}
	for _, item := range l {
		ids = ids.Add(item.Item())
	}
	return ids
}

func (l ItemList) HasDuplication() bool {
	lmap := make(map[id.ItemID]struct{})
	for _, i := range l {
		if _, ok := lmap[i.Item()]; ok {
			return true
		}
		lmap[i.Item()] = struct{}{}
	}

	return false
}
