package fme

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/reearth/reearthx/log"
)

type Interface interface {
	Request(ctx context.Context, r Request) error
}

type Request interface {
	Query() url.Values
	Name() string
}

type FME struct {
	base      *url.URL
	token     string
	resultURL string
	client    *http.Client
}

func New(baseUrl, token, resultURL string) (*FME, error) {
	b, err := url.Parse(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid base url: %w", err)
	}

	return &FME{
		base:      b,
		token:     token,
		resultURL: resultURL,
		client:    http.DefaultClient,
	}, nil
}

func (s *FME) Request(ctx context.Context, r Request) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.url(r), nil)
	if err != nil {
		return fmt.Errorf("failed to init request: %w", err)
	}

	if s.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("fmetoken token=%s", s.token))
	}

	log.Infof("fme: request: %s %s", req.Method, req.URL.String())

	res, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send: %w", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode >= 300 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read body: %w", err)
		}

		return fmt.Errorf("failed to request: code=%d, body=%s", res.StatusCode, body)
	}

	return nil
}

func (s *FME) url(r Request) string {
	u := s.base.JoinPath("fmejobsubmitter", r.Name()+".fmw")
	q := r.Query()
	q.Set("opt_servicemode", "async")
	q.Set("resultUrl", s.resultURL)
	u.RawQuery = q.Encode()
	return u.String()
}
