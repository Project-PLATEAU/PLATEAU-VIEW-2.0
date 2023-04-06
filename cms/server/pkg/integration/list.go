package integration

import (
	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type List []*Integration

func (l List) SortByID() List {
	m := slices.Clone(l)
	slices.SortFunc(m, func(a, b *Integration) bool {
		return a.ID().Compare(b.ID()) < 0
	})
	return m
}

func (l List) Clone() List {
	return util.Map(l, func(m *Integration) *Integration { return m.Clone() })
}

func (l List) ActiveWebhooks(ty event.Type) []*Webhook {
	return lo.FlatMap(l, func(i *Integration, _ int) []*Webhook {
		return lo.Filter(i.Webhooks(), func(w *Webhook, _ int) bool {
			return w.Trigger().IsActive(ty) && w.Active()
		})
	})
}
