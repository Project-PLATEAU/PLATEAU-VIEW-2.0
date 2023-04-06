package searchindex

import (
	"context"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

const storageModel = "itemasset"

type StorageItems []StorageItem

func (s StorageItems) FindByAsset(aid string) *StorageItem {
	for _, i := range s {
		if slices.Contains(i.Asset, aid) {
			return &i
		}
	}
	return nil
}

func (s StorageItems) FindByItem(iid string) *StorageItem {
	for _, i := range s {
		if i.Item == iid {
			return &i
		}
	}
	return nil
}

type StorageItem struct {
	ID    string   `json:"id,omitempty" cms:"id"`
	Item  string   `json:"item,omitempty" cms:"item,text"`
	Asset []string `json:"asset,omitempty" cms:"asset,asset"`
}

func (s StorageItem) RemoveAsset(aid string) StorageItem {
	s.Asset = lo.Filter(s.Asset, func(a string, _ int) bool { return a != aid })
	return s
}

type Storage struct {
	c   cms.Interface
	pid string
	mid string
}

func NewStorage(c cms.Interface, pid, mid string) *Storage {
	if mid == "" {
		mid = storageModel
	}
	return &Storage{
		c:   c,
		pid: pid,
		mid: mid,
	}
}

func (s *Storage) All(ctx context.Context) (r StorageItems, err error) {
	items, err := s.c.GetItemsByKey(ctx, s.pid, s.mid, false)
	if err != nil {
		return nil, err
	}

	for _, i := range items.Items {
		s := StorageItem{}
		i.Unmarshal(&s)
		if s.Item != "" {
			r = append(r, s)
		}
	}
	return
}

func (s *Storage) FindByAsset(ctx context.Context, assetid string) (r StorageItem, err error) {
	all, err := s.All(ctx)
	if err != nil {
		return
	}

	r2 := all.FindByAsset(assetid)
	if r2 == nil {
		return r, rerror.ErrNotFound
	}
	return *r2, nil
}

func (s *Storage) FindByItem(ctx context.Context, itemid string) (r StorageItem, err error) {
	all, err := s.All(ctx)
	if err != nil {
		return
	}

	r2 := all.FindByItem(itemid)
	if r2 == nil {
		return r, rerror.ErrNotFound
	}
	return *r2, nil
}

func (s *Storage) Set(ctx context.Context, item StorageItem) (err error) {
	citem := cms.Item{}
	cms.Marshal(item, &citem)
	if citem.ID == "" {
		_, err = s.c.CreateItemByKey(ctx, s.pid, storageModel, citem.Fields)
	} else if len(item.Asset) > 0 {
		_, err = s.c.UpdateItem(ctx, citem.ID, citem.Fields)
	} else {
		err = s.Delete(ctx, item.ID)
	}
	return
}

func (s *Storage) Delete(ctx context.Context, id string) error {
	return s.c.DeleteItem(ctx, id)
}
