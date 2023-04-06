package request

import "golang.org/x/exp/slices"

type List []*Request

func (l List) UpdateStatus(state State) {
	for _, request := range l {
		request.SetState(state)
	}
}

func (l List) Ordered() List {
	o := slices.Clone(l)
	slices.SortFunc(o, func(a, b *Request) bool {
		return a.UpdatedAt().After(b.UpdatedAt())
	})
	return o
}
