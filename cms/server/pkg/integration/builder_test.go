package integration

import (
	"net/url"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Builder
	}{
		{
			name: "new",
			want: &Builder{i: &Integration{}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, New(), "New()")
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	iId := id.NewIntegrationID()
	now := time.Now()
	type fields struct {
		i *Integration
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Integration
		wantErr error
	}{
		{
			name:    "no id",
			fields:  fields{i: &Integration{}},
			want:    nil,
			wantErr: ErrInvalidID,
		},
		{
			name: "no updated at",
			fields: fields{i: &Integration{
				id: iId,
			}},
			want: &Integration{
				id:        iId,
				updatedAt: iId.Timestamp(),
			},
			wantErr: nil,
		},
		{
			name: "full",
			fields: fields{i: &Integration{
				id:        iId,
				updatedAt: now,
			}},
			want: &Integration{
				id:        iId,
				updatedAt: now,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
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

func TestBuilder_MustBuild(t *testing.T) {
	iId := id.NewIntegrationID()
	now := time.Now()
	type fields struct {
		i *Integration
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Integration
		wantErr error
	}{
		{
			name:    "no id",
			fields:  fields{i: &Integration{}},
			want:    nil,
			wantErr: ErrInvalidID,
		},
		{
			name: "no updated at",
			fields: fields{i: &Integration{
				id: iId,
			}},
			want: &Integration{
				id:        iId,
				updatedAt: iId.Timestamp(),
			},
			wantErr: nil,
		},
		{
			name: "full",
			fields: fields{i: &Integration{
				id:        iId,
				updatedAt: now,
			}},
			want: &Integration{
				id:        iId,
				updatedAt: now,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			if tt.wantErr != nil {
				assert.PanicsWithValue(t, tt.wantErr, func() {
					b.MustBuild()
				})
			} else {
				assert.Equal(t, tt.want, b.MustBuild())
			}
		})
	}
}

func TestBuilder_NewID(t *testing.T) {
	type fields struct {
		i *Integration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "new",
			fields: fields{i: &Integration{}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			b.NewID()
			assert.False(t, b.i.id.IsEmpty())
		})
	}
}

func TestBuilder_ID(t *testing.T) {
	iId := id.NewIntegrationID()
	type fields struct {
		i *Integration
	}
	type args struct {
		id ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args:   args{id: iId},
			want:   &Builder{i: &Integration{id: iId}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			assert.Equalf(t, tt.want, b.ID(tt.args.id), "ID(%v)", tt.args.id)
		})
	}
}

func TestBuilder_Name(t *testing.T) {
	type fields struct {
		i *Integration
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args:   args{name: "test"},
			want:   &Builder{i: &Integration{name: "test"}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			assert.Equalf(t, tt.want, b.Name(tt.args.name), "Name(%v)", tt.args.name)
		})
	}
}

func TestBuilder_Description(t *testing.T) {
	type fields struct {
		i *Integration
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name:   "set",
			fields: fields{&Integration{}},
			args:   args{description: "test"},
			want:   &Builder{i: &Integration{description: "test"}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			assert.Equalf(t, tt.want, b.Description(tt.args.description), "Description(%v)", tt.args.description)
		})
	}
}

func TestBuilder_Type(t *testing.T) {
	type fields struct {
		i *Integration
	}
	type args struct {
		t Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args:   args{t: TypePublic},
			want:   &Builder{i: &Integration{iType: TypePublic}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			assert.Equalf(t, tt.want, b.Type(tt.args.t), "Type(%v)", tt.args.t)
		})
	}
}

func TestBuilder_LogoUrl(t *testing.T) {
	type fields struct {
		i *Integration
	}
	type args struct {
		logoURL *url.URL
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args:   args{logoURL: lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))},
			want:   &Builder{i: &Integration{logoUrl: lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))}},
		},
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args:   args{logoURL: nil},
			want:   &Builder{i: &Integration{logoUrl: nil}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			assert.Equalf(t, tt.want, b.LogoUrl(tt.args.logoURL), "LogoUrl(%v)", tt.args.logoURL)
		})
	}
}

func TestBuilder_Developer(t *testing.T) {
	uId := id.NewUserID()
	type fields struct {
		i *Integration
	}
	type args struct {
		developer UserID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args:   args{developer: uId},
			want:   &Builder{i: &Integration{developer: uId}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			assert.Equalf(t, tt.want, b.Developer(tt.args.developer), "Developer(%v)", tt.args.developer)
		})
	}
}

func TestBuilder_Webhook(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	type fields struct {
		i *Integration
	}
	type args struct {
		webhook []*Webhook
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args:   args{webhook: []*Webhook{}},
			want:   &Builder{i: &Integration{webhooks: []*Webhook{}}},
		},
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args: args{webhook: []*Webhook{{
				id:        wId,
				name:      "xyz",
				url:       nil,
				active:    true,
				trigger:   WebhookTrigger{},
				updatedAt: now,
			}}},
			want: &Builder{i: &Integration{webhooks: []*Webhook{{
				id:        wId,
				name:      "xyz",
				url:       nil,
				active:    true,
				trigger:   WebhookTrigger{},
				updatedAt: now,
			}}}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			assert.Equalf(t, tt.want, b.Webhook(tt.args.webhook), "Webhook(%v)", tt.args.webhook)
		})
	}
}

func TestBuilder_Token(t *testing.T) {
	type fields struct {
		i *Integration
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args:   args{token: "xyz"},
			want:   &Builder{i: &Integration{token: "xyz"}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			assert.Equalf(t, tt.want, b.Token(tt.args.token), "Token(%v)", tt.args.token)
		})
	}
}

func TestBuilder_UpdatedAt(t *testing.T) {
	now := time.Now()
	type fields struct {
		i *Integration
	}
	type args struct {
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name:   "set",
			fields: fields{i: &Integration{}},
			args:   args{updatedAt: now},
			want:   &Builder{i: &Integration{updatedAt: now}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				i: tt.fields.i,
			}
			assert.Equalf(t, tt.want, b.UpdatedAt(tt.args.updatedAt), "UpdatedAt(%v)", tt.args.updatedAt)
		})
	}
}
