package cms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
)

type Interface interface {
	GetModel(ctx context.Context, modelID string) (*Model, error)
	GetModelByKey(ctx context.Context, proejctID, modelID string) (*Model, error)
	GetItem(ctx context.Context, itemID string, asset bool) (*Item, error)
	GetItemsPartially(ctx context.Context, modelID string, page, perPage int, asset bool) (*Items, error)
	GetItems(ctx context.Context, modelID string, asset bool) (*Items, error)
	GetItemsPartiallyByKey(ctx context.Context, projectIDOrAlias, modelIDOrKey string, page, perPage int, asset bool) (*Items, error)
	GetItemsByKey(ctx context.Context, projectIDOrAlias, modelIDOrKey string, asset bool) (*Items, error)
	CreateItem(ctx context.Context, modelID string, fields []Field) (*Item, error)
	CreateItemByKey(ctx context.Context, projectID, modelID string, fields []Field) (*Item, error)
	UpdateItem(ctx context.Context, itemID string, fields []Field) (*Item, error)
	DeleteItem(ctx context.Context, itemID string) error
	Asset(ctx context.Context, id string) (*Asset, error)
	UploadAsset(ctx context.Context, projectID, url string) (string, error)
	UploadAssetDirectly(ctx context.Context, projectID, name string, data io.Reader) (string, error)
	CommentToItem(ctx context.Context, assetID, content string) error
	CommentToAsset(ctx context.Context, assetID, content string) error
}

type CMS struct {
	base   *url.URL
	token  string
	client *http.Client
}

func New(base, token string) (*CMS, error) {
	b, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %w", err)
	}

	return &CMS{
		base:   b,
		token:  token,
		client: http.DefaultClient,
	}, nil
}

func (c *CMS) assetParam(asset bool) map[string][]string {
	if !asset {
		return make(map[string][]string)
	}
	return map[string][]string{
		"asset": {"true"},
	}
}

func (c *CMS) GetModel(ctx context.Context, modelID string) (*Model, error) {
	b, err := c.send(ctx, http.MethodGet, []string{"api", "models", modelID}, "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get an model: %w", err)
	}
	defer func() { _ = b.Close() }()

	model := &Model{}
	if err := json.NewDecoder(b).Decode(model); err != nil {
		return nil, fmt.Errorf("failed to parse an model: %w", err)
	}

	return model, nil
}

func (c *CMS) GetModelByKey(ctx context.Context, projectKey, modelKey string) (*Model, error) {
	b, err := c.send(ctx, http.MethodGet, []string{"api", "projects", projectKey, "models", modelKey}, "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get an model: %w", err)
	}
	defer func() { _ = b.Close() }()

	model := &Model{}
	if err := json.NewDecoder(b).Decode(model); err != nil {
		return nil, fmt.Errorf("failed to parse an model: %w", err)
	}

	return model, nil
}

func (c *CMS) GetItem(ctx context.Context, itemID string, asset bool) (*Item, error) {
	b, err := c.send(ctx, http.MethodGet, []string{"api", "items", itemID}, "", c.assetParam(asset))
	if err != nil {
		return nil, fmt.Errorf("failed to get an item: %w", err)
	}
	defer func() { _ = b.Close() }()

	item := &Item{}
	if err := json.NewDecoder(b).Decode(item); err != nil {
		return nil, fmt.Errorf("failed to parse an item: %w", err)
	}

	return item, nil
}

func (c *CMS) GetItemsPartially(ctx context.Context, modelID string, page, perPage int, asset bool) (*Items, error) {
	q := c.assetParam(asset)
	if page >= 1 {
		q["page"] = []string{strconv.Itoa(page)}
	}
	if perPage >= 1 {
		q["perPage"] = []string{strconv.Itoa(perPage)}
	}

	b, err := c.send(ctx, http.MethodGet, []string{"api", "models", modelID, "items"}, "", q)
	if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	}
	defer func() { _ = b.Close() }()

	items := &Items{}
	if err := json.NewDecoder(b).Decode(items); err != nil {
		return nil, fmt.Errorf("failed to parse items: %w", err)
	}

	return items, nil
}

func (c *CMS) GetItems(ctx context.Context, modelID string, asset bool) (*Items, error) {
	var items *Items
	const perPage = 100
	for p := 1; ; p++ {
		i, err := c.GetItemsPartially(ctx, modelID, p, perPage, asset)
		if err != nil {
			return nil, err
		}

		if i == nil || i.PerPage <= 0 {
			return nil, fmt.Errorf("invalid response: %#v", i)
		}

		if items == nil {
			items = i
		} else {
			items.Items = append(items.Items, i.Items...)
		}

		allPageCount := i.TotalCount / i.PerPage
		if i.Page >= allPageCount {
			break
		}
	}

	return items, nil
}

func (c *CMS) GetItemsPartiallyByKey(ctx context.Context, projectIDOrAlias, modelIDOrAlias string, page, perPage int, asset bool) (*Items, error) {
	q := c.assetParam(asset)
	if page >= 1 {
		q["page"] = []string{strconv.Itoa(page)}
	}
	if perPage >= 1 {
		q["perPage"] = []string{strconv.Itoa(perPage)}
	}

	b, err := c.send(ctx, http.MethodGet, []string{"api", "projects", projectIDOrAlias, "models", modelIDOrAlias, "items"}, "", q)
	if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	}
	defer func() { _ = b.Close() }()

	items := &Items{}
	if err := json.NewDecoder(b).Decode(items); err != nil {
		return nil, fmt.Errorf("failed to parse items: %w", err)
	}

	return items, nil
}

func (c *CMS) GetItemsByKey(ctx context.Context, projectIDOrAlias, modelIDOrAlias string, asset bool) (*Items, error) {
	var items *Items
	const perPage = 100
	for p := 1; ; p++ {
		i, err := c.GetItemsPartiallyByKey(ctx, projectIDOrAlias, modelIDOrAlias, p, perPage, asset)
		if err != nil {
			return nil, err
		}

		if i == nil || i.PerPage <= 0 {
			return nil, fmt.Errorf("invalid response: %#v", i)
		}

		if items == nil {
			items = i
		} else {
			items.Items = append(items.Items, i.Items...)
		}

		if !i.HasNext() {
			break
		}
	}

	return items, nil
}

func (c *CMS) CreateItem(ctx context.Context, modelID string, fields []Field) (*Item, error) {
	rb := map[string]any{
		"fields": fields,
	}

	b, err := c.send(ctx, http.MethodPost, []string{"api", "models", modelID, "items"}, "", rb)
	if err != nil {
		return nil, fmt.Errorf("failed to create an item: %w", err)
	}
	defer func() { _ = b.Close() }()

	item := &Item{}
	if err := json.NewDecoder(b).Decode(&item); err != nil {
		return nil, fmt.Errorf("failed to parse an item: %w", err)
	}

	return item, nil
}

func (c *CMS) CreateItemByKey(ctx context.Context, projectID, modelID string, fields []Field) (*Item, error) {
	rb := map[string]any{
		"fields": fields,
	}

	b, err := c.send(ctx, http.MethodPost, []string{"api", "projects", projectID, "models", modelID, "items"}, "", rb)
	if err != nil {
		return nil, fmt.Errorf("failed to create an item: %w", err)
	}
	defer func() { _ = b.Close() }()

	item := &Item{}
	if err := json.NewDecoder(b).Decode(&item); err != nil {
		return nil, fmt.Errorf("failed to parse an item: %w", err)
	}

	return item, nil
}

func (c *CMS) UpdateItem(ctx context.Context, itemID string, fields []Field) (*Item, error) {
	rb := map[string]any{
		"fields": fields,
	}

	b, err := c.send(ctx, http.MethodPatch, []string{"api", "items", itemID}, "", rb)
	if err != nil {
		return nil, fmt.Errorf("failed to update an item: %w", err)
	}
	defer func() { _ = b.Close() }()

	item := &Item{}
	if err := json.NewDecoder(b).Decode(&item); err != nil {
		return nil, fmt.Errorf("failed to parse an item: %w", err)
	}

	return item, nil
}

func (c *CMS) DeleteItem(ctx context.Context, itemID string) error {
	b, err := c.send(ctx, http.MethodDelete, []string{"api", "items", itemID}, "", nil)
	if err != nil {
		return fmt.Errorf("failed to delete an item: %w", err)
	}
	defer func() { _ = b.Close() }()
	return nil
}

func (c *CMS) UploadAsset(ctx context.Context, projectID, url string) (string, error) {
	rb := map[string]string{
		"url": url,
	}

	b, err2 := c.send(ctx, http.MethodPost, []string{"api", "projects", projectID, "assets"}, "", rb)
	if err2 != nil {
		log.Errorf("cms: upload asset: failed to upload an asset: %s", err2)
		return "", fmt.Errorf("failed to upload an asset: %w", err2)
	}

	defer func() { _ = b.Close() }()

	body, err2 := io.ReadAll(b)
	if err2 != nil {
		return "", fmt.Errorf("failed to read body: %w", err2)
	}

	type res struct {
		ID string `json:"id"`
	}

	r := &res{}
	if err2 := json.Unmarshal(body, &r); err2 != nil {
		return "", fmt.Errorf("failed to parse body: %w", err2)
	}

	return r.ID, nil
}

func (c *CMS) UploadAssetDirectly(ctx context.Context, projectID, name string, data io.Reader) (string, error) {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	go func() {
		var err error
		defer func() {
			_ = mw.Close()
			_ = pw.CloseWithError(err)
		}()

		fw, err2 := mw.CreateFormFile("file", name)
		if err2 != nil {
			err = err2
			return
		}
		_, err = io.Copy(fw, data)
	}()

	b, err2 := c.send(ctx, http.MethodPost, []string{"api", "projects", projectID, "assets"}, mw.FormDataContentType(), pr)
	if err2 != nil {
		log.Errorf("cms: upload asset: failed to upload an asset with multipart: %s", err2)
		return "", fmt.Errorf("failed to upload an asset with multipart: %w", err2)
	}

	defer func() { _ = b.Close() }()

	body, err2 := io.ReadAll(b)
	if err2 != nil {
		return "", fmt.Errorf("failed to read body: %w", err2)
	}

	type res struct {
		ID string `json:"id"`
	}

	r := &res{}
	if err2 := json.Unmarshal(body, &r); err2 != nil {
		return "", fmt.Errorf("failed to parse body: %w", err2)
	}

	return r.ID, nil
}

func (c *CMS) Asset(ctx context.Context, assetID string) (*Asset, error) {
	b, err := c.send(ctx, http.MethodGet, []string{"api", "assets", assetID}, "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get an asset: %w", err)
	}
	defer func() { _ = b.Close() }()

	a := &Asset{}
	if err := json.NewDecoder(b).Decode(a); err != nil {
		return nil, fmt.Errorf("failed to parse an asset: %w", err)
	}

	return a, nil
}

func (c *CMS) CommentToItem(ctx context.Context, itemID, content string) error {
	rb := map[string]string{
		"content": content,
	}

	b, err := c.send(ctx, http.MethodPost, []string{"api", "items", itemID, "comments"}, "", rb)
	if err != nil {
		return fmt.Errorf("failed to comment to item %s: %w", itemID, err)
	}
	defer func() { _ = b.Close() }()

	return nil
}

func (c *CMS) CommentToAsset(ctx context.Context, assetID, content string) error {
	rb := map[string]string{
		"content": content,
	}

	b, err := c.send(ctx, http.MethodPost, []string{"api", "assets", assetID, "comments"}, "", rb)
	if err != nil {
		return fmt.Errorf("failed to comment to asset %s: %w", assetID, err)
	}
	defer func() { _ = b.Close() }()

	return nil
}

func (c *CMS) send(ctx context.Context, m string, p []string, ct string, body any) (io.ReadCloser, error) {
	req, err := c.request(ctx, m, p, ct, body)
	if err != nil {
		return nil, err
	}

	log.Infof("CMS: request: %s %s body=%+v", req.Method, req.URL, body)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if res.StatusCode >= 300 {
		defer func() {
			_ = res.Body.Close()
		}()

		if res.StatusCode == http.StatusNotFound {
			return nil, rerror.ErrNotFound
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read body: %w", err)
		}

		return nil, fmt.Errorf("failed to request: code=%d, body=%s", res.StatusCode, b)
	}

	return res.Body, nil
}

func (c *CMS) request(ctx context.Context, m string, p []string, ct string, body any) (*http.Request, error) {
	if m != "GET" && ct == "" {
		ct = "application/json"
	}

	u := c.base.JoinPath(p...)
	var b io.Reader

	if m == "POST" || m == "PUT" || m == "PATCH" {
		if ct == "application/json" && body != nil {
			bb, err := json.Marshal(body)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal JSON: %w", err)
			}
			b = bytes.NewReader(bb)
		} else if strings.HasPrefix(ct, "multipart/form-data") {
			if bb, ok := body.(io.Reader); ok {
				b = bb
			}
		}
	} else if q, ok := body.(map[string][]string); ok {
		v := url.Values(q)
		u.RawQuery = v.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, m, u.String(), b)
	if err != nil {
		return nil, fmt.Errorf("failed to init request: %w", err)
	}
	if ct != "" {
		req.Header.Set("Content-Type", ct)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	return req, nil
}
