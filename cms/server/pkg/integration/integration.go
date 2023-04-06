package integration

import (
	"crypto/rand"
	"math/big"
	"net/url"
	"time"

	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Integration struct {
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

func (i *Integration) ID() ID {
	return i.id
}

func (i *Integration) Name() string {
	return i.name
}

func (i *Integration) SetName(name string) {
	i.name = name
}

func (i *Integration) Description() string {
	return i.description
}

func (i *Integration) SetDescription(description string) {
	i.description = description
}

func (i *Integration) Type() Type {
	return i.iType
}

func (i *Integration) SetType(t Type) {
	i.iType = t
}

func (i *Integration) LogoUrl() *url.URL {
	return i.logoUrl
}

func (i *Integration) SetLogoUrl(logoUrl *url.URL) {
	i.logoUrl = logoUrl
}

func (i *Integration) Token() string {
	return i.token
}

func (i *Integration) SetToken(token string) {
	i.token = token
}

func (i *Integration) RandomToken() {
	t, err := randomString(43)
	if err != nil {
		return
	}
	i.token = "secret_" + t
}

func (i *Integration) Developer() UserID {
	return i.developer
}

func (i *Integration) SetDeveloper(developer UserID) {
	i.developer = developer
}

func (i *Integration) Webhooks() []*Webhook {
	return i.webhooks
}

func (i *Integration) Webhook(wId WebhookID) (*Webhook, bool) {
	return lo.Find(i.webhooks, func(w *Webhook) bool { return w.id == wId })
}

func (i *Integration) AddWebhook(w *Webhook) {
	if w == nil {
		return
	}
	i.webhooks = append(i.webhooks, w)
}

func (i *Integration) UpdateWebhook(wId WebhookID, w *Webhook) bool {
	if w == nil {
		return false
	}
	_, idx, ok := lo.FindIndexOf(i.webhooks, func(w *Webhook) bool { return w.id == wId })
	if !ok || idx >= len(i.webhooks) {
		return false
	}
	i.webhooks[idx] = w
	return true
}

func (i *Integration) DeleteWebhook(wId WebhookID) bool {
	_, idx, ok := lo.FindIndexOf(i.webhooks, func(w *Webhook) bool { return w.id == wId })
	if !ok || idx >= len(i.webhooks) {
		return false
	}
	i.webhooks = slices.Delete(i.webhooks, idx, idx+1)
	return true
}

func (i *Integration) SetWebhook(webhook []*Webhook) {
	i.webhooks = webhook
}

func (i *Integration) UpdatedAt() time.Time {
	if i.updatedAt.IsZero() {
		return i.id.Timestamp()
	}
	return i.updatedAt
}

func (i *Integration) SetUpdatedAt(updatedAt time.Time) {
	i.updatedAt = updatedAt
}

func (i *Integration) CreatedAt() time.Time {
	return i.id.Timestamp()
}

func (i *Integration) Clone() *Integration {
	if i == nil {
		return nil
	}

	var u *url.URL = nil
	if i.logoUrl != nil {
		u, _ = url.Parse(i.logoUrl.String())
	}
	return &Integration{
		id:          i.id.Clone(),
		name:        i.name,
		description: i.description,
		logoUrl:     u,
		iType:       i.iType,
		token:       i.token,
		developer:   i.developer,
		webhooks:    util.Map(i.webhooks, func(w *Webhook) *Webhook { return w.Clone() }),
		updatedAt:   i.updatedAt,
	}
}

func randomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[randIndex.Int64()]
	}

	return string(result), nil
}
