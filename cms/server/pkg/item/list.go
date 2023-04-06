package item

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type List []*Item

func (l List) SortByID() List {
	m := slices.Clone(l)
	slices.SortFunc(m, func(a, b *Item) bool {
		return a.ID().Compare(b.ID()) < 0
	})
	return m
}

func (l List) ItemsByField(fid id.FieldID, value any) List {
	return lo.Filter(l, func(i *Item, _ int) bool {
		return i.HasField(fid, value)
	})
}

func (l List) SortByTimestamp() List {
	m := slices.Clone(l)
	slices.SortFunc(m, func(a, b *Item) bool {
		return a.timestamp.Before(b.Timestamp())
	})
	return m
}

func (l List) FilterFields(lids id.FieldIDList) List {
	return lo.Map(l, func(i *Item, _ int) *Item {
		return i.FilterFields(lids)
	})
}

type VersionedList []Versioned

func (l VersionedList) SortByCreationDate(dir Direction) VersionedList {
	m := slices.Clone(l)

	if dir == DescDirection {
		slices.SortFunc(m, func(a, b Versioned) bool {
			return a.Value().ID().Timestamp().After(b.Value().ID().Timestamp())
		})
	} else {
		slices.SortFunc(m, func(a, b Versioned) bool {
			return a.Value().ID().Timestamp().Before(b.Value().ID().Timestamp())
		})
	}
	return m
}

func (l VersionedList) SortByTimestamp(dir Direction) VersionedList {
	m := slices.Clone(l)

	if dir == DescDirection {
		slices.SortFunc(m, func(a, b Versioned) bool {
			return a.Value().timestamp.After(b.Value().Timestamp())
		})
	} else {
		slices.SortFunc(m, func(a, b Versioned) bool {
			return a.Value().timestamp.Before(b.Value().Timestamp())
		})
	}
	return m
}

func (l VersionedList) Sort(st *Sort) VersionedList {
	if st == nil {
		return l.Sort(&Sort{
			SortBy: SortTypeCreationDate,
		})
	}

	switch st.SortBy {
	case SortTypeCreationDate:
		return l.SortByCreationDate(st.Direction)
	case SortTypeModificationDate:
		return l.SortByTimestamp(st.Direction)
	default:
		return l.SortByTimestamp(DescDirection)
	}
}

func (l VersionedList) FilterFields(fields id.FieldIDList) VersionedList {
	return lo.Map(l, func(a Versioned, _ int) Versioned {
		return version.ValueFrom(a, a.Value().FilterFields(fields))
	})
}

func (l VersionedList) Unwrap() List {
	if l == nil {
		return nil
	}
	return version.UnwrapValues(l)
}

func (l VersionedList) Item(iid id.ItemID) Versioned {
	if l == nil {
		return nil
	}
	for _, versioned := range l {
		if versioned.Value().ID() == iid {
			return versioned
		}
	}
	return nil
}
