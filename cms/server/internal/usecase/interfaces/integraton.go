package interfaces

import (
	"context"
	"net/url"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
)

type CreateIntegrationParam struct {
	Name        string
	Description *string
	Type        integration.Type
	Logo        url.URL
}

type UpdateIntegrationParam struct {
	Name        *string
	Description *string
	Logo        *url.URL
}

type CreateWebhookParam struct {
	Name    string
	URL     url.URL
	Secret  string
	Active  bool
	Trigger *WebhookTriggerParam
}

type UpdateWebhookParam struct {
	Name    *string
	URL     *url.URL
	Active  *bool
	Trigger *WebhookTriggerParam
	Secret  *string
}

type WebhookTriggerParam map[event.Type]bool

type Integration interface {
	FindByIDs(context.Context, id.IntegrationIDList, *usecase.Operator) (integration.List, error)
	FindByMe(context.Context, *usecase.Operator) (integration.List, error)
	Create(context.Context, CreateIntegrationParam, *usecase.Operator) (*integration.Integration, error)
	Update(context.Context, id.IntegrationID, UpdateIntegrationParam, *usecase.Operator) (*integration.Integration, error)
	Delete(context.Context, id.IntegrationID, *usecase.Operator) error

	CreateWebhook(context.Context, id.IntegrationID, CreateWebhookParam, *usecase.Operator) (*integration.Webhook, error)
	UpdateWebhook(context.Context, id.IntegrationID, id.WebhookID, UpdateWebhookParam, *usecase.Operator) (*integration.Webhook, error)
	DeleteWebhook(context.Context, id.IntegrationID, id.WebhookID, *usecase.Operator) error
}
