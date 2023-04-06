package fme

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var _ Interface = (*FME)(nil)

func TestFME(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	ctx := context.Background()
	wantReq := ConversionRequest{
		ID:                 "xxx",
		Target:             "target",
		PRCS:               "6669",
		DevideODC:          true,
		QualityCheckParams: "params",
		QualityCheck:       true,
	}

	// valid
	calls := mockFMEServer(t, "http://fme.example.com", "TOKEN", wantReq, "https://example.com")
	f := lo.Must(New("http://fme.example.com", "TOKEN", "https://example.com"))
	req := wantReq
	req.QualityCheck = false
	assert.NoError(t, f.Request(ctx, req))
	assert.Equal(t, 1, calls("convert-all"))
	req.QualityCheck = true
	assert.NoError(t, f.Request(ctx, req))
	assert.Equal(t, 1, calls("quality-check-and-convert-all"))

	// invalid token
	httpmock.Reset()
	calls = mockFMEServer(t, "http://fme.example.com", "TOKEN", wantReq, "https://example.com")
	f = lo.Must(New("http://fme.example.com", "TOKEN2", "https://example.com"))
	req = wantReq
	req.QualityCheck = false
	assert.ErrorContains(t, f.Request(ctx, req), "failed to request: code=401")
	assert.Equal(t, 1, calls("convert-all"))
	req.QualityCheck = true
	assert.ErrorContains(t, f.Request(ctx, req), "failed to request: code=401")
	assert.Equal(t, 1, calls("quality-check-and-convert-all"))

	// invalid queries
	httpmock.Reset()
	calls = mockFMEServer(t, "http://fme.example.com", "TOKEN", wantReq, "https://example.com")
	f = lo.Must(New("http://fme.example.com", "TOKEN", "https://example.com"))
	req = ConversionRequest{
		ID:     wantReq.ID,
		Target: "target!",
		PRCS:   wantReq.PRCS,
	}
	req.QualityCheck = false
	assert.ErrorContains(t, f.Request(ctx, req), "failed to request: code=400")
	assert.Equal(t, 1, calls("convert-all"))
	req.QualityCheck = true
	assert.ErrorContains(t, f.Request(ctx, req), "failed to request: code=400")
	assert.Equal(t, 1, calls("quality-check-and-convert-all"))
}

func mockFMEServer(t *testing.T, host, token string, r ConversionRequest, resultURL string) func(string) int {
	t.Helper()
	u := host + "/fmejobsubmitter/plateau2022-cms/"

	responder := func(req *http.Request) (*http.Response, error) {
		if t := parseFMEToken(req); t != token {
			return httpmock.NewJsonResponse(http.StatusUnauthorized, map[string]any{
				"statusInfo": map[string]any{
					"message": "failure",
					"status":  "failure",
				},
			})
		}

		q := req.URL.Query()
		invalid := false

		if q.Get("opt_servicemode") != "async" ||
			resultURL != q.Get("resultUrl") {
			invalid = true
		}

		q.Del("opt_servicemode")
		q.Del("resultUrl")
		if !reflect.DeepEqual(r.Query(), q) {
			invalid = true
		}

		if invalid {
			return httpmock.NewJsonResponse(http.StatusBadRequest, map[string]any{
				"statusInfo": map[string]any{
					"message": "failure",
					"status":  "failure",
				},
			})
		}

		return httpmock.NewJsonResponse(200, map[string]any{
			"statusInfo": map[string]any{
				"message": "success",
				"status":  "success",
			},
		})
	}

	httpmock.RegisterResponder("POST", u+"convert-all.fmw", responder)
	httpmock.RegisterResponder("POST", u+"quality-check-and-convert-all.fmw", responder)

	return func(ws string) int {
		return httpmock.GetCallCountInfo()[fmt.Sprintf("POST %s%s.fmw", u, ws)]
	}
}

func parseFMEToken(r *http.Request) string {
	aut := r.Header.Get("Authorization")
	_, token, found := strings.Cut(aut, "fmetoken token=")
	if !found {
		return ""
	}
	return token
}
