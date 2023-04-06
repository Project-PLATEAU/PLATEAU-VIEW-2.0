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

func TestNewWebhookBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *WebhookBuilder
	}{
		{
			name: "name",
			want: &WebhookBuilder{
				w: &Webhook{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, NewWebhookBuilder(), "NewWebhookBuilder()")
		})
	}
}

func TestWebhookBuilder_Active(t *testing.T) {
	type fields struct {
		w *Webhook
	}
	type args struct {
		active bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *WebhookBuilder
	}{
		{
			name:   "true",
			fields: fields{w: &Webhook{}},
			args:   args{active: true},
			want:   &WebhookBuilder{w: &Webhook{active: true}},
		},
		{
			name:   "false",
			fields: fields{w: &Webhook{}},
			args:   args{active: false},
			want:   &WebhookBuilder{w: &Webhook{active: false}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &WebhookBuilder{
				w: tt.fields.w,
			}
			assert.Equalf(t, tt.want, b.Active(tt.args.active), "Active(%v)", tt.args.active)
		})
	}
}

func TestWebhookBuilder_Build(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	type fields struct {
		w *Webhook
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Webhook
		wantErr error
	}{
		{
			name:    "no id",
			fields:  fields{w: &Webhook{}},
			want:    nil,
			wantErr: ErrInvalidID,
		},
		{
			name:    "no update at",
			fields:  fields{w: &Webhook{id: wId}},
			want:    &Webhook{id: wId, updatedAt: wId.Timestamp()},
			wantErr: nil,
		},
		{
			name:    "full",
			fields:  fields{w: &Webhook{id: wId, updatedAt: now, active: true, name: "xyz", trigger: WebhookTrigger{}}},
			want:    &Webhook{id: wId, updatedAt: now, active: true, name: "xyz", trigger: WebhookTrigger{}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &WebhookBuilder{
				w: tt.fields.w,
			}
			got, err := b.Build()
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				assert.Nil(t, got)
				return
			}
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "Build()")
		})
	}
}

func TestWebhookBuilder_ID(t *testing.T) {
	wId := id.NewWebhookID()
	type fields struct {
		w *Webhook
	}
	type args struct {
		wId WebhookID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *WebhookBuilder
	}{
		{
			name:   "set",
			fields: fields{w: &Webhook{}},
			args:   args{wId: wId},
			want:   &WebhookBuilder{w: &Webhook{id: wId}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &WebhookBuilder{
				w: tt.fields.w,
			}
			assert.Equalf(t, tt.want, b.ID(tt.args.wId), "ID(%v)", tt.args.wId)
		})
	}
}

func TestWebhookBuilder_MustBuild(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	type fields struct {
		w *Webhook
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Webhook
		wantErr error
	}{
		{
			name:    "no id",
			fields:  fields{w: &Webhook{}},
			want:    nil,
			wantErr: ErrInvalidID,
		},
		{
			name:    "no update at",
			fields:  fields{w: &Webhook{id: wId}},
			want:    &Webhook{id: wId, updatedAt: wId.Timestamp()},
			wantErr: nil,
		},
		{
			name:    "full",
			fields:  fields{w: &Webhook{id: wId, updatedAt: now, active: true, name: "xyz", trigger: WebhookTrigger{}}},
			want:    &Webhook{id: wId, updatedAt: now, active: true, name: "xyz", trigger: WebhookTrigger{}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &WebhookBuilder{
				w: tt.fields.w,
			}

			if tt.wantErr != nil {
				assert.PanicsWithValue(t, tt.wantErr, func() { _ = b.MustBuild() })
			} else {
				assert.Equal(t, tt.want, b.MustBuild())
			}
		})
	}
}

func TestWebhookBuilder_Name(t *testing.T) {
	type fields struct {
		w *Webhook
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *WebhookBuilder
	}{
		{
			name:   "set",
			fields: fields{w: &Webhook{}},
			args:   args{name: "test"},
			want:   &WebhookBuilder{w: &Webhook{name: "test"}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &WebhookBuilder{
				w: tt.fields.w,
			}
			assert.Equalf(t, tt.want, b.Name(tt.args.name), "Name(%v)", tt.args.name)
		})
	}
}

func TestWebhookBuilder_Trigger(t *testing.T) {
	type fields struct {
		w *Webhook
	}
	type args struct {
		trigger WebhookTrigger
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *WebhookBuilder
	}{
		{
			name:   "set",
			fields: fields{w: &Webhook{}},
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
			want: &WebhookBuilder{w: &Webhook{trigger: WebhookTrigger{
				event.ItemCreate:      true,
				event.ItemUpdate:      true,
				event.ItemDelete:      true,
				event.ItemPublish:     true,
				event.ItemUnpublish:   true,
				event.AssetCreate:     true,
				event.AssetDecompress: true,
				event.AssetDelete:     true,
			}}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &WebhookBuilder{
				w: tt.fields.w,
			}
			assert.Equalf(t, tt.want, b.Trigger(tt.args.trigger), "Trigger(%v)", tt.args.trigger)
		})
	}
}

func TestWebhookBuilder_UpdatedAt(t *testing.T) {
	now := time.Now()
	type fields struct {
		w *Webhook
	}
	type args struct {
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *WebhookBuilder
	}{
		{
			name:   "set",
			fields: fields{w: &Webhook{}},
			args:   args{updatedAt: now},
			want:   &WebhookBuilder{w: &Webhook{updatedAt: now}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &WebhookBuilder{
				w: tt.fields.w,
			}
			assert.Equalf(t, tt.want, b.UpdatedAt(tt.args.updatedAt), "UpdatedAt(%v)", tt.args.updatedAt)
		})
	}
}

func TestWebhookBuilder_Url(t *testing.T) {
	type fields struct {
		w *Webhook
	}
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *WebhookBuilder
	}{
		{
			name:   "set",
			fields: fields{w: &Webhook{}},
			args:   args{url: lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))},
			want:   &WebhookBuilder{w: &Webhook{url: lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &WebhookBuilder{
				w: tt.fields.w,
			}
			assert.Equalf(t, tt.want, b.Url(tt.args.url), "Url(%v)", tt.args.url)
		})
	}
}

func TestWebhookBuilder_NewID(t *testing.T) {
	type fields struct {
		w *Webhook
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "test",
			fields: fields{w: &Webhook{}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &WebhookBuilder{
				w: tt.fields.w,
			}
			b.NewID()
			assert.False(t, b.w.id.IsNil())
			assert.False(t, b.w.id.IsEmpty())
		})
	}
}
