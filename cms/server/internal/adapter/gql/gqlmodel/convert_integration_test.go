package gqlmodel

import (
	"net/url"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestToIntegration(t *testing.T) {
	iId := integration.NewID()
	uId, dId := id.NewUserID(), id.NewUserID()
	now := time.Now()
	tests := []struct {
		name        string
		integration *integration.Integration
		user        id.UserID
		want        *Integration
	}{
		{
			name:        "nil",
			integration: nil,
			user:        id.NewUserID(),
			want:        nil,
		},
		{
			name: "success",
			integration: integration.New().ID(iId).Name("N1").Description("D1").
				LogoUrl(lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))).
				Type(integration.TypePrivate).
				Developer(dId).
				UpdatedAt(now).
				Token("t1").
				Webhook([]*integration.Webhook{}).
				MustBuild(),
			want: &Integration{
				ID:          IDFrom(iId),
				Name:        "N1",
				Description: lo.ToPtr("D1"),
				LogoURL:     *lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
				IType:       "Private",
				DeveloperID: IDFrom(dId),
				Developer:   nil,
				Config:      nil,
				CreatedAt:   iId.Timestamp(),
				UpdatedAt:   now,
			},
			user: uId,
		},
		{
			name: "success developer",
			integration: integration.New().ID(iId).Name("N1").Description("D1").
				LogoUrl(lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))).
				Type(integration.TypePrivate).
				Developer(dId).
				UpdatedAt(now).
				Token("t1").
				Webhook([]*integration.Webhook{}).
				MustBuild(),
			want: &Integration{
				ID:          IDFrom(iId),
				Name:        "N1",
				Description: lo.ToPtr("D1"),
				LogoURL:     *lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
				IType:       "Private",
				DeveloperID: IDFrom(dId),
				Developer:   nil,
				Config: &IntegrationConfig{
					Token:    "t1",
					Webhooks: []*Webhook{},
				},
				CreatedAt: iId.Timestamp(),
				UpdatedAt: now,
			},
			user: dId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ToIntegration(tt.integration, &tt.user)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToIntegrationType(t *testing.T) {
	tests := []struct {
		name string
		args integration.Type
		want IntegrationType
	}{
		{name: "public", args: integration.TypePublic, want: IntegrationTypePublic},
		{name: "private", args: integration.TypePrivate, want: IntegrationTypePrivate},
		{name: "default", args: "", want: ""},
		{name: "default2", args: "some value", want: ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, ToIntegrationType(tt.args))
		})
	}
}

func TestToWebhook(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	tests := []struct {
		name string
		args *integration.Webhook
		want *Webhook
	}{
		{
			name: "nil",
			args: nil,
			want: nil,
		},
		{
			name: "success",
			args: integration.NewWebhookBuilder().ID(wId).Name("WH1").Active(true).UpdatedAt(now).
				Url(lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))).
				MustBuild(),
			want: &Webhook{
				ID:     IDFrom(wId),
				Name:   "WH1",
				URL:    *lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
				Active: true,
				Trigger: &WebhookTrigger{
					OnItemCreate:      lo.ToPtr(false),
					OnItemUpdate:      lo.ToPtr(false),
					OnItemDelete:      lo.ToPtr(false),
					OnItemPublish:     lo.ToPtr(false),
					OnItemUnPublish:   lo.ToPtr(false),
					OnAssetUpload:     lo.ToPtr(false),
					OnAssetDecompress: lo.ToPtr(false),
					OnAssetDelete:     lo.ToPtr(false),
				},
				CreatedAt: wId.Timestamp(),
				UpdatedAt: now,
			},
		},
		{
			name: "success 2",
			args: integration.NewWebhookBuilder().ID(wId).Name("WH1").Active(true).UpdatedAt(now).
				Url(lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))).
				Trigger(integration.WebhookTrigger{
					event.ItemCreate:      true,
					event.ItemUpdate:      true,
					event.ItemDelete:      true,
					event.ItemPublish:     true,
					event.ItemUnpublish:   true,
					event.AssetCreate:     true,
					event.AssetDecompress: true,
					event.AssetDelete:     true,
				}).
				MustBuild(),
			want: &Webhook{
				ID:     IDFrom(wId),
				Name:   "WH1",
				URL:    *lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
				Active: true,
				Trigger: &WebhookTrigger{
					OnItemCreate:      lo.ToPtr(true),
					OnItemUpdate:      lo.ToPtr(true),
					OnItemDelete:      lo.ToPtr(true),
					OnItemPublish:     lo.ToPtr(true),
					OnItemUnPublish:   lo.ToPtr(true),
					OnAssetUpload:     lo.ToPtr(true),
					OnAssetDecompress: lo.ToPtr(true),
					OnAssetDelete:     lo.ToPtr(true),
				},
				CreatedAt: wId.Timestamp(),
				UpdatedAt: now,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, ToWebhook(tt.args))
		})
	}
}

func TestToWebhooks(t *testing.T) {
	wId := id.NewWebhookID()
	now := time.Now()
	tests := []struct {
		name string
		args []*integration.Webhook
		want []*Webhook
	}{
		{
			name: "nil",
			args: nil,
			want: []*Webhook{},
		},
		{
			name: "empty",
			args: []*integration.Webhook{},
			want: []*Webhook{},
		},
		{
			name: "success",
			args: []*integration.Webhook{
				integration.NewWebhookBuilder().ID(wId).Name("WH1").Active(true).UpdatedAt(now).
					Url(lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))).
					MustBuild(),
				integration.NewWebhookBuilder().ID(wId).Name("WH1").Active(true).UpdatedAt(now).
					Url(lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test"))).
					Trigger(integration.WebhookTrigger{
						event.ItemCreate:      true,
						event.ItemUpdate:      true,
						event.ItemDelete:      true,
						event.ItemPublish:     true,
						event.ItemUnpublish:   true,
						event.AssetCreate:     true,
						event.AssetDecompress: true,
						event.AssetDelete:     true,
					}).
					MustBuild(),
			},
			want: []*Webhook{
				{
					ID:     IDFrom(wId),
					Name:   "WH1",
					URL:    *lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
					Active: true,
					Trigger: &WebhookTrigger{
						OnItemCreate:      lo.ToPtr(false),
						OnItemUpdate:      lo.ToPtr(false),
						OnItemDelete:      lo.ToPtr(false),
						OnItemPublish:     lo.ToPtr(false),
						OnItemUnPublish:   lo.ToPtr(false),
						OnAssetUpload:     lo.ToPtr(false),
						OnAssetDecompress: lo.ToPtr(false),
						OnAssetDelete:     lo.ToPtr(false),
					},
					CreatedAt: wId.Timestamp(),
					UpdatedAt: now,
				},
				{
					ID:     IDFrom(wId),
					Name:   "WH1",
					URL:    *lo.Must(url.Parse("https://sub.hugo.com/dir?p=1#test")),
					Active: true,
					Trigger: &WebhookTrigger{
						OnItemCreate:      lo.ToPtr(true),
						OnItemUpdate:      lo.ToPtr(true),
						OnItemDelete:      lo.ToPtr(true),
						OnItemPublish:     lo.ToPtr(true),
						OnItemUnPublish:   lo.ToPtr(true),
						OnAssetUpload:     lo.ToPtr(true),
						OnAssetDecompress: lo.ToPtr(true),
						OnAssetDelete:     lo.ToPtr(true),
					},
					CreatedAt: wId.Timestamp(),
					UpdatedAt: now,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, ToWebhooks(tt.args))
		})
	}
}
