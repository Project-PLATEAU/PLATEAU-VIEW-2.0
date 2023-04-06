package integration

import (
	"net/url"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestList_SortByID(t *testing.T) {
	id1 := NewID()
	id2 := NewID()

	list := List{
		&Integration{id: id2},
		&Integration{id: id1},
	}
	res := list.SortByID()
	assert.Equal(t, List{
		&Integration{id: id1},
		&Integration{id: id2},
	}, res)
	// test whether original list is not modified
	assert.Equal(t, List{
		&Integration{id: id2},
		&Integration{id: id1},
	}, list)
}

func TestList_Clone(t *testing.T) {
	id := NewID()
	list := List{&Integration{id: id}}
	got := list.Clone()
	assert.Equal(t, list, got)
	assert.NotSame(t, list[0], got[0])

	got[0].id = NewID()
	// test whether original list is not modified
	assert.Equal(t, list, List{&Integration{id: id}})
}

func TestList_ActiveWebhooks(t *testing.T) {
	now := time.Now()
	iID1 := NewID()
	iID2 := NewID()
	iID3 := NewID()
	uID := user.NewID()
	wID1 := NewWebhookID()
	wID2 := NewWebhookID()
	wID3 := NewWebhookID()
	wID4 := NewWebhookID()

	w1 := &Webhook{
		id:     wID1,
		name:   "w xyz",
		url:    lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
		active: true,
		trigger: WebhookTrigger{
			event.ItemCreate:      true,
			event.ItemUpdate:      true,
			event.ItemDelete:      true,
			event.ItemPublish:     false,
			event.ItemUnpublish:   false,
			event.AssetCreate:     false,
			event.AssetDecompress: false,
			event.AssetDelete:     false,
		},
		updatedAt: now,
	}
	w2 := &Webhook{
		id:     wID2,
		name:   "w abc",
		url:    lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
		active: true,
		trigger: WebhookTrigger{
			event.ItemCreate:      true,
			event.ItemUpdate:      false,
			event.ItemDelete:      false,
			event.ItemPublish:     false,
			event.ItemUnpublish:   false,
			event.AssetCreate:     false,
			event.AssetDecompress: false,
			event.AssetDelete:     false,
		},
		updatedAt: now,
	}
	w3 := &Webhook{
		id:     wID3,
		name:   "xxx",
		url:    lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
		active: true,
		trigger: WebhookTrigger{
			event.ItemCreate:      true,
			event.ItemUpdate:      true,
			event.ItemDelete:      false,
			event.ItemPublish:     false,
			event.ItemUnpublish:   false,
			event.AssetCreate:     false,
			event.AssetDecompress: false,
			event.AssetDelete:     false},
		updatedAt: now,
	}
	w4 := &Webhook{
		id:        wID4,
		name:      "xxx",
		url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
		active:    false,
		updatedAt: now,
	}

	i1 := &Integration{
		id:          iID1,
		name:        "xyz",
		description: "xyz d",
		logoUrl:     lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
		iType:       TypePublic,
		token:       "token",
		developer:   uID,
		webhooks: []*Webhook{
			w1, w2,
		},
	}
	i2 := &Integration{
		id:          iID2,
		name:        "xxx",
		description: "xyz d",
		logoUrl:     lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
		iType:       TypePublic,
		token:       "token",
		developer:   uID,
		webhooks: []*Webhook{
			w3,
		},
	}
	i3 := &Integration{
		id:          iID3,
		name:        "xxx",
		description: "xyz d",
		logoUrl:     lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
		iType:       TypePublic,
		token:       "token",
		developer:   uID,
		webhooks: []*Webhook{
			w4,
		},
	}
	iList := List{i1, i2, i3}

	// type test struct {
	// 	eType event.Type
	// 	expected []*Webhook
	// }

	tests := []struct {
		name     string
		eType    event.Type
		expected []*Webhook
	}{
		{
			name:  "integrations have multiple active webhooks",
			eType: event.Type("item.create"),
			expected: []*Webhook{
				w1, w2, w3,
			},
		},
		{
			name:  "integrations have one active webhook each",
			eType: event.Type("item.update"),
			expected: []*Webhook{
				w1, w3,
			},
		},
		{
			name:  "one integration have one active webhook",
			eType: event.Type("item.delete"),
			expected: []*Webhook{
				w1,
			},
		},
		{
			name:     "no integration have active webhooks",
			eType:    event.Type("item.publish"),
			expected: []*Webhook{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, iList.ActiveWebhooks(tc.eType))
		})
	}
}
