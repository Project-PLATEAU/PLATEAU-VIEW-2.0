package gcp

import (
	"encoding/json"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/integrationapi"
	"github.com/reearth/reearth-cms/server/pkg/task"
)

type webhookData struct {
	URL       string                  `json:"url"`
	Secret    string                  `json:"secret"`
	Timestamp time.Time               `json:"timestamp"`
	WebhookID string                  `json:"webhookId"`
	EventID   string                  `json:"eventId"`
	EventType string                  `json:"type"`
	EventData any                     `json:"data"`
	Operator  integrationapi.Operator `json:"operator"`
}

func marshalWebhookData(w *task.WebhookPayload, urlResolver asset.URLResolver) ([]byte, error) {
	ed, err := integrationapi.NewEventWith(w.Event, w.Override, "", urlResolver)
	if err != nil {
		return nil, err
	}

	d := webhookData{
		URL:       w.Webhook.URL().String(),
		Secret:    w.Webhook.Secret(),
		Timestamp: ed.Timestamp,
		WebhookID: w.Webhook.ID().String(),
		EventID:   ed.ID,
		EventType: ed.Type,
		EventData: ed.Data,
		Operator:  ed.Operator,
	}

	return json.Marshal(d)
}
