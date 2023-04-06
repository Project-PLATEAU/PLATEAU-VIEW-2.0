package integration

import (
	"net/url"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestWebhook_Active(t *testing.T) {
	tests := []struct {
		name string
		w    *Webhook
		want bool
	}{
		{
			name: "true",
			w:    &Webhook{active: true},
			want: true,
		},
		{
			name: "false",
			w:    &Webhook{active: false},
			want: false,
		},
		{
			name: "not set",
			w:    &Webhook{active: false},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, tt.w.Active(), "Active()")
		})
	}
}

func TestWebhook_Clone(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	tests := []struct {
		name string
		w    *Webhook
		want *Webhook
	}{
		{
			name: "clone",
			w: &Webhook{
				id:     wId,
				name:   "w1",
				url:    lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
				active: true,
				trigger: WebhookTrigger{
					event.ItemCreate:      false,
					event.ItemUpdate:      false,
					event.ItemDelete:      false,
					event.ItemPublish:     false,
					event.ItemUnpublish:   false,
					event.AssetCreate:     false,
					event.AssetDecompress: false,
					event.AssetDelete:     false,
				},
				updatedAt: now,
			},
			want: &Webhook{
				id:     wId,
				name:   "w1",
				url:    lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
				active: true,
				trigger: WebhookTrigger{
					event.ItemCreate:      false,
					event.ItemUpdate:      false,
					event.ItemDelete:      false,
					event.ItemPublish:     false,
					event.ItemUnpublish:   false,
					event.AssetCreate:     false,
					event.AssetDecompress: false,
					event.AssetDelete:     false,
				},
				updatedAt: now,
			},
		},
		{
			name: "nil",
			w:    nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, tt.w.Clone(), "Clone()")
			if tt.want != nil {
				assert.NotSame(t, tt.want, tt.w)
			}
		})
	}
}

func TestWebhook_CreatedAt(t *testing.T) {
	wId := id.NewWebhookID()
	tests := []struct {
		name string
		w    *Webhook
		want time.Time
	}{
		{
			name: "test",
			w:    &Webhook{id: wId},
			want: wId.Timestamp(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, tt.w.CreatedAt(), "CreatedAt()")
		})
	}
}

func TestWebhook_ID(t *testing.T) {
	wId := id.NewWebhookID()
	tests := []struct {
		name string
		w    *Webhook
		want WebhookID
	}{
		{
			name: "set",
			w:    &Webhook{id: wId},
			want: wId,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, tt.w.ID(), "ID()")
		})
	}
}

func TestWebhook_Name(t *testing.T) {
	tests := []struct {
		name string
		w    *Webhook
		want string
	}{
		{
			name: "set",
			w:    &Webhook{name: "test"},
			want: "test",
		},
		{
			name: "not set",
			w:    &Webhook{},
			want: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, tt.w.Name(), "Name()")
		})
	}
}

func TestWebhook_SetActive(t *testing.T) {
	type args struct {
		active bool
	}
	tests := []struct {
		name string
		w    *Webhook
		args args
		want bool
	}{
		{
			name: "set",
			w:    &Webhook{},
			args: args{active: true},
			want: true,
		},
		{
			name: "unset",
			w:    &Webhook{},
			args: args{active: false},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.w.SetActive(tt.args.active)
			assert.Equal(t, tt.want, tt.w.active)
		})
	}
}

func TestWebhook_SetName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		w    *Webhook
		args args
		want string
	}{
		{
			name: "set",
			w:    &Webhook{},
			args: args{name: "xyz"},
			want: "xyz",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.w.SetName(tt.args.name)
			assert.Equal(t, tt.want, tt.w.name)
		})
	}
}

func TestWebhook_SetTrigger(t *testing.T) {
	type args struct {
		trigger WebhookTrigger
	}
	tests := []struct {
		name string
		w    *Webhook
		args args
		want WebhookTrigger
	}{
		{
			name: "set",
			w:    &Webhook{},
			args: args{trigger: WebhookTrigger{
				event.ItemCreate:      false,
				event.ItemUpdate:      false,
				event.ItemDelete:      false,
				event.ItemPublish:     false,
				event.ItemUnpublish:   false,
				event.AssetCreate:     false,
				event.AssetDecompress: false,
				event.AssetDelete:     false,
			}},
			want: WebhookTrigger{
				event.ItemCreate:      false,
				event.ItemUpdate:      false,
				event.ItemDelete:      false,
				event.ItemPublish:     false,
				event.ItemUnpublish:   false,
				event.AssetCreate:     false,
				event.AssetDecompress: false,
				event.AssetDelete:     false,
			},
		},
		{
			name: "set true",
			w:    &Webhook{},
			args: args{trigger: WebhookTrigger{
				event.ItemCreate:      true,
				event.ItemUpdate:      true,
				event.ItemDelete:      true,
				event.ItemPublish:     true,
				event.ItemUnpublish:   true,
				event.AssetCreate:     true,
				event.AssetDecompress: true,
				event.AssetDelete:     true,
			}},
			want: WebhookTrigger{
				event.ItemCreate:      true,
				event.ItemUpdate:      true,
				event.ItemDelete:      true,
				event.ItemPublish:     true,
				event.ItemUnpublish:   true,
				event.AssetCreate:     true,
				event.AssetDecompress: true,
				event.AssetDelete:     true,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.w.SetTrigger(tt.args.trigger)
			assert.Equal(t, tt.want, tt.w.trigger)
		})
	}
}

func TestWebhook_SetUpdatedAt(t *testing.T) {
	now := time.Now()
	type args struct {
		updatedAt time.Time
	}
	tests := []struct {
		name string
		w    *Webhook
		args args
		want time.Time
	}{
		{
			name: "set",
			w:    &Webhook{},
			args: args{updatedAt: now},
			want: now,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.w.SetUpdatedAt(tt.args.updatedAt)
			assert.Equal(t, tt.want, tt.w.updatedAt)
		})
	}
}

func TestWebhook_SetUrl(t *testing.T) {
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name string
		w    *Webhook
		args args
		want string
	}{
		{
			name: "set",
			w:    &Webhook{},
			args: args{lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))},
			want: "https://sub.hugo.com/dir?p=1#test",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.w.SetURL(tt.args.url)
			assert.Equal(t, tt.want, tt.w.url.String())
		})
	}
}

func TestWebhook_Trigger(t *testing.T) {
	tests := []struct {
		name string
		w    *Webhook
		want WebhookTrigger
	}{
		{
			name: "get falsy",
			w: &Webhook{trigger: WebhookTrigger{
				event.ItemCreate:      false,
				event.ItemUpdate:      false,
				event.ItemDelete:      false,
				event.ItemPublish:     false,
				event.ItemUnpublish:   false,
				event.AssetCreate:     false,
				event.AssetDecompress: false,
				event.AssetDelete:     false,
			}},
			want: WebhookTrigger{
				event.ItemCreate:      false,
				event.ItemUpdate:      false,
				event.ItemDelete:      false,
				event.ItemPublish:     false,
				event.ItemUnpublish:   false,
				event.AssetCreate:     false,
				event.AssetDecompress: false,
				event.AssetDelete:     false,
			},
		},
		{
			name: "get true",
			w: &Webhook{trigger: WebhookTrigger{
				event.ItemCreate:      true,
				event.ItemUpdate:      true,
				event.ItemDelete:      true,
				event.ItemPublish:     true,
				event.ItemUnpublish:   true,
				event.AssetCreate:     true,
				event.AssetDecompress: true,
				event.AssetDelete:     true,
			}},
			want: WebhookTrigger{
				event.ItemCreate:      true,
				event.ItemUpdate:      true,
				event.ItemDelete:      true,
				event.ItemPublish:     true,
				event.ItemUnpublish:   true,
				event.AssetCreate:     true,
				event.AssetDecompress: true,
				event.AssetDelete:     true,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, tt.w.Trigger(), "Trigger()")
		})
	}
}

func TestWebhook_UpdatedAt(t *testing.T) {
	now := time.Now()
	wId := id.NewWebhookID()
	tests := []struct {
		name string
		w    *Webhook
		want time.Time
	}{
		{
			name: "set",
			w:    &Webhook{id: wId, updatedAt: now},
			want: now,
		},
		{
			name: "not set",
			w:    &Webhook{id: wId},
			want: wId.Timestamp(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, tt.w.UpdatedAt(), "UpdatedAt()")
		})
	}
}

func TestWebhook_Url(t *testing.T) {
	tests := []struct {
		name string
		w    *Webhook
		want string
	}{
		{
			name: "set",
			w:    &Webhook{url: lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))},
			want: "https://sub.hugo.com/dir?p=1#test",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, tt.w.URL().String(), "Url()")
		})
	}
}
