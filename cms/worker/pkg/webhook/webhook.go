package webhook

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/reearth/reearthx/util"
)

type Webhook struct {
	URL       string    `json:"url"`
	Secret    string    `json:"secret"`
	Timestamp time.Time `json:"timestamp"`
	WebhookID string    `json:"webhookId"`
	EventID   string    `json:"eventId"`
	EventType string    `json:"type"`
	EventData any       `json:"data"`
	Operator  any       `json:"operator"`
}

type requestBody struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
	Data      any       `json:"data"`
	Operator  any       `json:"operator"`
}

func Send(ctx context.Context, w *Webhook) error {
	b, err := w.requestBody()
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", w.URL, bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("failed to create a request: %w", err)
	}

	now := util.Now()
	signature := Sign(b, []byte(w.Secret), now, "v1")

	req.Header.Set("Reearth-Signature", signature)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send a request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 300 {
		return fmt.Errorf("ERROR: id=%s, url=%s, status=%d", w.EventID, w.URL, res.StatusCode)
	}

	return nil
}

func (w Webhook) requestBody() ([]byte, error) {
	b := requestBody{
		ID:        w.EventID,
		Timestamp: w.Timestamp,
		Type:      w.EventType,
		Data:      w.EventData,
		Operator:  w.Operator,
	}
	return json.Marshal(b)
}

func Sign(payload, secret []byte, t time.Time, v string) string {
	mac := hmac.New(sha256.New, secret)
	_, _ = mac.Write([]byte(fmt.Sprintf("%s:%d:", v, t.Unix())))
	_, _ = mac.Write(payload)
	s := hex.EncodeToString(mac.Sum(nil))
	return fmt.Sprintf("%s,t=%d,%s", v, t.Unix(), s)
}
