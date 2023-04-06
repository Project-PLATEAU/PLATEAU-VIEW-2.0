package http

import (
	"context"

	"github.com/reearth/reearth-cms/worker/internal/usecase/interactor"
	"github.com/reearth/reearth-cms/worker/pkg/webhook"
)

type WebhookController struct {
	usecase *interactor.Usecase
}

func NewWebhookController(u *interactor.Usecase) *WebhookController {
	return &WebhookController{
		usecase: u,
	}
}

func (c *WebhookController) Webhook(ctx context.Context, w *webhook.Webhook) error {
	return c.usecase.SendWebhook(ctx, w)
}
