package asset

import "github.com/samber/lo"

type Map map[ID]*Asset

func (m Map) List() List {
	return lo.MapToSlice(m, func(_ ID, v *Asset) *Asset {
		return v
	})
}

func (m Map) ListFrom(ids IDList) (res List) {
	for _, id := range ids {
		if a, ok := m[id]; ok {
			res = append(res, a)
		}
	}
	return
}
