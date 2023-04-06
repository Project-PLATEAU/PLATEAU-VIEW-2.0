package schema

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type List []*Schema

func (l List) SortByID() List {
	m := slices.Clone(l)
	slices.SortFunc(m, func(a, b *Schema) bool {
		return a.ID().Compare(b.ID()) < 0
	})
	return m
}

func (l List) Clone() List {
	return util.Map(l, func(s *Schema) *Schema { return s.Clone() })
}

type FieldList []*Field

func (l FieldList) Find(fid FieldID) *Field {
	f, _ := lo.Find(l, func(f *Field) bool {
		return f.ID() == fid
	})
	return f
}

func (l FieldList) SortByID() FieldList {
	m := slices.Clone(l)
	slices.SortFunc(m, func(a, b *Field) bool {
		return a.ID().Compare(b.ID()) < 0
	})
	return m
}

func (l FieldList) Clone() FieldList {
	return util.Map(l, func(f *Field) *Field { return f.Clone() })
}

func (l FieldList) IDs() (ids id.FieldIDList) {
	for _, sf := range l {
		ids = ids.Add(sf.ID())
	}
	return
}

func (l FieldList) Ordered() FieldList {
	o := slices.Clone(l)
	slices.SortFunc(o, func(a, b *Field) bool {
		return a.Order() < b.Order()
	})
	return o
}

func (l FieldList) Count() int {
	return len(l)
}
