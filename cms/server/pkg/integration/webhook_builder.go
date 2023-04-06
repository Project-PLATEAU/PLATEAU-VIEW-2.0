package integration

import (
	"net/url"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
)

type WebhookBuilder struct {
	w *Webhook
}

func NewWebhookBuilder() *WebhookBuilder {
	return &WebhookBuilder{w: &Webhook{}}
}

func (b *WebhookBuilder) Build() (*Webhook, error) {
	if b.w.id.IsNil() {
		return nil, ErrInvalidID
	}
	if b.w.updatedAt.IsZero() {
		b.w.updatedAt = b.w.CreatedAt()
	}
	return b.w, nil
}

func (b *WebhookBuilder) MustBuild() *Webhook {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *WebhookBuilder) NewID() *WebhookBuilder {
	b.w.id = id.NewWebhookID()
	return b
}

func (b *WebhookBuilder) ID(wId WebhookID) *WebhookBuilder {
	b.w.id = wId
	return b
}

func (b *WebhookBuilder) Name(name string) *WebhookBuilder {
	b.w.name = name
	return b
}

func (b *WebhookBuilder) Url(url *url.URL) *WebhookBuilder {
	b.w.url = url
	return b
}

func (b *WebhookBuilder) Active(active bool) *WebhookBuilder {
	b.w.active = active
	return b
}

func (b *WebhookBuilder) Trigger(trigger WebhookTrigger) *WebhookBuilder {
	b.w.trigger = trigger
	return b
}

func (b *WebhookBuilder) UpdatedAt(updatedAt time.Time) *WebhookBuilder {
	b.w.updatedAt = updatedAt
	return b
}

func (b *WebhookBuilder) Secret(secret string) *WebhookBuilder {
	b.w.secret = secret
	return b
}
