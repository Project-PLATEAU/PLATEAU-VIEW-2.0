package integration

import (
	"net/url"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestIntegration_Clone(t *testing.T) {
	iId := id.NewIntegrationID()
	uId := id.NewUserID()
	wId := id.NewWebhookID()
	now := time.Now()
	tests := []struct {
		name string
		i    *Integration
		want *Integration
	}{
		{
			name: "test",
			i: &Integration{
				id:          iId,
				name:        "xyz",
				description: "xyz d",
				logoUrl:     lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
				iType:       TypePublic,
				token:       "token",
				developer:   uId,
				webhooks: []*Webhook{
					{
						id:        wId,
						name:      "w xyz",
						url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
						active:    true,
						trigger:   WebhookTrigger{},
						updatedAt: now,
					},
				},
				updatedAt: now,
			},
			want: &Integration{
				id:          iId,
				name:        "xyz",
				description: "xyz d",
				logoUrl:     lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
				iType:       TypePublic,
				token:       "token",
				developer:   uId,
				webhooks: []*Webhook{
					{
						id:        wId,
						name:      "w xyz",
						url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
						active:    true,
						trigger:   WebhookTrigger{},
						updatedAt: now,
					},
				},
				updatedAt: now,
			},
		},
		{
			name: "nil",
			i:    nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, tt.i.Clone(), "Clone()")
		})
	}
}

func TestIntegration_CreatedAt(t *testing.T) {
	iId := id.NewIntegrationID()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name:   "test",
			fields: fields{id: iId},
			want:   iId.Timestamp(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.CreatedAt(), "CreatedAt()")
		})
	}
}

func TestIntegration_Description(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "test",
			fields: fields{description: "xyz"},
			want:   "xyz",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.Description(), "Description()")
		})
	}
}

func TestIntegration_Developer(t *testing.T) {
	uId := id.NewUserID()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   UserID
	}{
		{
			name:   "test",
			fields: fields{developer: uId},
			want:   uId,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.Developer(), "Developer()")
		})
	}
}

func TestIntegration_ID(t *testing.T) {
	iId := id.NewIntegrationID()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   ID
	}{
		{
			name:   "test",
			fields: fields{id: iId},
			want:   iId,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.ID(), "ID()")
		})
	}
}

func TestIntegration_LogoUrl(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *url.URL
	}{
		{
			name:   "test",
			fields: fields{logoUrl: lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))},
			want:   lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.LogoUrl(), "LogoUrl()")
		})
	}
}

func TestIntegration_Name(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "test",
			fields: fields{name: "xyz"},
			want:   "xyz",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.Name(), "Name()")
		})
	}
}

func TestIntegration_SetDescription(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{description: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			i.SetDescription(tt.args.description)
			assert.Equal(t, tt.want, i.description)
		})
	}
}

func TestIntegration_SetDeveloper(t *testing.T) {
	uId := id.NewUserID()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		developer UserID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   UserID
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{developer: uId},
			want:   uId,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			i.SetDeveloper(tt.args.developer)
			assert.Equal(t, tt.want, i.developer)
		})
	}
}

func TestIntegration_SetLogoUrl(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		logoUrl *url.URL
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{logoUrl: lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))},
			want:   "https://sub.hugo.com/dir?p=1#test",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			i.SetLogoUrl(tt.args.logoUrl)
			assert.Equal(t, tt.want, i.logoUrl.String())
		})
	}
}

func TestIntegration_SetName(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{name: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			i.SetName(tt.args.name)
			assert.Equal(t, tt.want, i.name)
		})
	}
}

func TestIntegration_SetToken(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "test",
			fields: fields{},
			args:   args{token: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			i.SetToken(tt.args.token)
			assert.Equal(t, tt.want, i.token)
		})
	}
}

func TestIntegration_SetType(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		t Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Type
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{t: TypePublic},
			want:   TypePublic,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			i.SetType(tt.args.t)
			assert.Equal(t, tt.want, i.iType)
		})
	}
}

func TestIntegration_SetUpdatedAt(t *testing.T) {
	now := time.Now()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{updatedAt: now},
			want:   now,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			i.SetUpdatedAt(tt.args.updatedAt)
			assert.Equal(t, tt.want, i.updatedAt)
		})
	}
}

func TestIntegration_SetWebhook(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		webhook []*Webhook
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Webhook
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{webhook: []*Webhook{}},
			want:   []*Webhook{},
		},
		{
			name:   "set",
			fields: fields{},
			args: args{webhook: []*Webhook{{
				id:        wId,
				name:      "xyz",
				url:       nil,
				active:    false,
				trigger:   WebhookTrigger{},
				updatedAt: now,
			}}},
			want: []*Webhook{{
				id:        wId,
				name:      "xyz",
				url:       nil,
				active:    false,
				trigger:   WebhookTrigger{},
				updatedAt: now,
			}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			i.SetWebhook(tt.args.webhook)
			assert.Equal(t, tt.want, i.webhooks)
		})
	}
}

func TestIntegration_Token(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "test",
			fields: fields{token: "xyz"},
			want:   "xyz",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.Token(), "Token()")
		})
	}
}

func TestIntegration_Type(t *testing.T) {
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "test",
			fields: fields{iType: TypePublic},
			want:   TypePublic,
		},
		{
			name:   "test",
			fields: fields{iType: TypePrivate},
			want:   TypePrivate,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.Type(), "Type()")
		})
	}
}

func TestIntegration_UpdatedAt(t *testing.T) {
	iId := id.NewIntegrationID()
	now := time.Now()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name:   "test",
			fields: fields{updatedAt: now},
			want:   now,
		},
		{
			name:   "test",
			fields: fields{id: iId},
			want:   iId.Timestamp(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.UpdatedAt(), "UpdatedAt()")
		})
	}
}

func TestIntegration_Webhooks(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhook     []*Webhook
		updatedAt   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Webhook
	}{
		{
			name:   "test",
			fields: fields{webhook: []*Webhook{}},
			want:   []*Webhook{},
		},
		{
			name: "test",
			fields: fields{webhook: []*Webhook{
				{
					id:        wId,
					name:      "w xyz",
					url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
					active:    true,
					trigger:   WebhookTrigger{},
					updatedAt: now,
				},
			}},
			want: []*Webhook{
				{
					id:        wId,
					name:      "w xyz",
					url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
					active:    true,
					trigger:   WebhookTrigger{},
					updatedAt: now,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhook,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equalf(t, tt.want, i.Webhooks(), "Webhook()")
		})
	}
}

func TestIntegration_Webhook(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhooks    []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		wId WebhookID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Webhook
		want1  bool
	}{
		{
			name:   "test",
			fields: fields{webhooks: []*Webhook{}},
			args:   args{wId: id.NewWebhookID()},
			want:   nil,
			want1:  false,
		},
		{
			name: "test",
			fields: fields{webhooks: []*Webhook{
				{
					id:        wId,
					name:      "w xyz",
					url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
					active:    true,
					trigger:   WebhookTrigger{},
					updatedAt: now,
				},
			}},
			args: args{wId: wId},
			want: &Webhook{
				id:        wId,
				name:      "w xyz",
				url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
				active:    true,
				trigger:   WebhookTrigger{},
				updatedAt: now,
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhooks,
				updatedAt:   tt.fields.updatedAt,
			}
			got, got1 := i.Webhook(tt.args.wId)
			assert.Equalf(t, tt.want, got, "Webhook(%v)", tt.args.wId)
			assert.Equalf(t, tt.want1, got1, "Webhook(%v)", tt.args.wId)
		})
	}
}

func TestIntegration_AddWebhook(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhooks    []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		w *Webhook
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Webhook
	}{
		{
			name:   "test",
			fields: fields{webhooks: []*Webhook{}},
			args:   args{w: nil},
			want:   []*Webhook{},
		},
		{
			name:   "test",
			fields: fields{webhooks: []*Webhook{}},
			args: args{w: &Webhook{
				id:        wId,
				name:      "w xyz",
				url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
				active:    true,
				trigger:   WebhookTrigger{},
				updatedAt: now,
			}},
			want: []*Webhook{
				{
					id:        wId,
					name:      "w xyz",
					url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
					active:    true,
					trigger:   WebhookTrigger{},
					updatedAt: now,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhooks,
				updatedAt:   tt.fields.updatedAt,
			}
			i.AddWebhook(tt.args.w)
			assert.Equal(t, tt.want, i.webhooks)
		})
	}
}

func TestIntegration_UpdateWebhook(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhooks    []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		wId WebhookID
		w   *Webhook
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Webhook
		want1  bool
	}{
		{
			name:   "test",
			fields: fields{webhooks: []*Webhook{}},
			args:   args{wId: id.NewWebhookID(), w: nil},
			want:   []*Webhook{},
			want1:  false,
		},
		{
			name:   "test",
			fields: fields{webhooks: []*Webhook{}},
			args: args{
				wId: wId,
				w: &Webhook{
					id:        wId,
					name:      "w xyz",
					url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
					active:    true,
					trigger:   WebhookTrigger{},
					updatedAt: now,
				}},
			want:  []*Webhook{},
			want1: false,
		},
		{
			name: "test",
			fields: fields{webhooks: []*Webhook{
				{
					id:        wId,
					name:      "w xyz",
					url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
					active:    true,
					trigger:   WebhookTrigger{},
					updatedAt: now,
				},
			}},
			args: args{
				wId: wId,
				w: &Webhook{
					id:        wId,
					name:      "w xyz updated",
					url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
					active:    true,
					trigger:   WebhookTrigger{},
					updatedAt: now,
				}},
			want: []*Webhook{
				{
					id:        wId,
					name:      "w xyz updated",
					url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
					active:    true,
					trigger:   WebhookTrigger{},
					updatedAt: now,
				},
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhooks,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equal(t, tt.want1, i.UpdateWebhook(tt.args.wId, tt.args.w))
			assert.Equal(t, tt.want, i.webhooks)
		})
	}
}

func TestIntegration_DeleteWebhook(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	type fields struct {
		id          ID
		name        string
		description string
		logoUrl     *url.URL
		iType       Type
		token       string
		developer   UserID
		webhooks    []*Webhook
		updatedAt   time.Time
	}
	type args struct {
		wId WebhookID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Webhook
		want1  bool
	}{
		{
			name:   "test",
			fields: fields{webhooks: []*Webhook{}},
			args:   args{wId: id.NewWebhookID()},
			want:   []*Webhook{},
			want1:  false,
		},
		{
			name:   "test",
			fields: fields{webhooks: []*Webhook{}},
			args:   args{wId: wId},
			want:   []*Webhook{},
			want1:  false,
		},
		{
			name: "test",
			fields: fields{webhooks: []*Webhook{
				{
					id:        wId,
					name:      "w xyz",
					url:       lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test")),
					active:    true,
					trigger:   WebhookTrigger{},
					updatedAt: now,
				},
			}},
			args:  args{wId: wId},
			want:  []*Webhook{},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Integration{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				logoUrl:     tt.fields.logoUrl,
				iType:       tt.fields.iType,
				token:       tt.fields.token,
				developer:   tt.fields.developer,
				webhooks:    tt.fields.webhooks,
				updatedAt:   tt.fields.updatedAt,
			}
			assert.Equal(t, tt.want1, i.DeleteWebhook(tt.args.wId))
			assert.Equal(t, tt.want, i.webhooks)
		})
	}
}
