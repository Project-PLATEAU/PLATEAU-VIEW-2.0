package webhook

import (
	"context"
	"crypto/hmac"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	now := time.Date(2022, 10, 10, 1, 1, 1, 1, time.UTC)
	defer util.MockNow(now)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint1 := "https://example.com"
	endpoint2 := "https://example.com/fail"
	w := &Webhook{
		URL:       endpoint1,
		Secret:    "secret",
		Timestamp: now,
		EventID:   "event",
		EventType: "asset.create",
		EventData: `{
			"id": "aaa",
			"name": "name"
		}`,
	}

	httpmock.RegisterResponder("POST", endpoint1, func(req *http.Request) (*http.Response, error) {

		//check signature
		sign := req.Header.Get("Reearth-Signature")
		rawSign := lo.Must(hex.DecodeString(strings.Split(sign, ",")[2]))

		expectedSign := Sign(lo.Must(io.ReadAll(req.Body)), []byte(w.Secret), util.Now(), "v1")
		rawExpectedSign := lo.Must(hex.DecodeString(strings.Split(expectedSign, ",")[2]))

		if !hmac.Equal(rawSign, rawExpectedSign) {
			return nil, errors.New("invalid signature")
		}

		return httpmock.NewStringResponse(200, "hello"), nil
	})

	httpmock.RegisterResponder("POST", endpoint2, httpmock.NewStringResponder(500, `error`))

	// success
	err := Send(context.Background(), w)
	info := httpmock.GetCallCountInfo()
	assert.NoError(t, err)
	assert.Equal(t, 1, info["POST "+endpoint1])

	// should fail
	w.URL = endpoint2
	err = Send(context.Background(), w)
	info = httpmock.GetCallCountInfo()
	assert.Error(t, err)
	assert.Equal(t, 1, info["POST "+endpoint2])
}

func TestWebhook_requestBody(t *testing.T) {
	time := time.Date(2022, 10, 10, 1, 1, 1, 1, time.UTC)

	rawExpected := requestBody{
		ID:        "event",
		Timestamp: time,
		Type:      "asset.create",
		Data: `{
			"id": "aaa",
			"name": "name"
		}`,
	}
	expected, err := json.Marshal(rawExpected)
	assert.NoError(t, err)

	w := &Webhook{
		URL:       "https://example.com",
		Secret:    "secret",
		Timestamp: time,
		EventID:   "event",
		EventType: "asset.create",
		EventData: `{
			"id": "aaa",
			"name": "name"
		}`,
	}

	res, err := w.requestBody()
	assert.NoError(t, err)
	assert.Equal(t, expected, res)

}

func TestSign(t *testing.T) {
	time := time.Date(2022, 10, 10, 1, 1, 1, 1, time.UTC)
	assert.Equal(t, "v1,t=1665363661,df7340aca1823ae160b7f72e3851b4d6da7c026588b467a6c59fd0f7350c1e7f", Sign([]byte("a"), []byte("b"), time, "v1"))
}
