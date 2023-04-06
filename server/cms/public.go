package cms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
)

type PublicAPIListResponse[T any] struct {
	Results    []T `json:"results"`
	PerPage    int `json:"perPage"`
	Page       int `json:"page"`
	TotalCount int `json:"totalCount"`
}

func (r PublicAPIListResponse[T]) HasNext() bool {
	if r.PerPage == 0 {
		return false
	}
	return r.TotalCount > r.Page*r.PerPage
}

type PublicAsset struct {
	Type                    string   `json:"type,omitempty"`
	ID                      string   `json:"id,omitempty"`
	URL                     string   `json:"url,omitempty"`
	Files                   []string `json:"files,omitempty"`
	ContentType             string   `json:"contentType,omitempty"`
	ArchiveExtractionStatus string   `json:"archiveExtractionStatus,omitempty"`
}

func (a PublicAsset) IsExtractionDone() bool {
	return a.ArchiveExtractionStatus == "done"
}

type PublicAPIClient[T any] struct {
	c       *http.Client
	base    *url.URL
	project string
}

func NewPublicAPIClient[T any](c *http.Client, base, project string) (*PublicAPIClient[T], error) {
	if c == nil {
		c = http.DefaultClient
	}
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	return &PublicAPIClient[T]{
		c:       c,
		base:    u,
		project: project,
	}, nil
}

func (c *PublicAPIClient[T]) GetAllItems(ctx context.Context, model string) (res []T, err error) {
	perPage := 100
	for p := 1; ; p++ {
		r, err := c.GetItems(ctx, model, p, perPage)
		if err != nil {
			return nil, err
		}

		res = append(res, r.Results...)
		if !r.HasNext() {
			break
		}
	}
	return
}

func (c *PublicAPIClient[T]) GetItems(ctx context.Context, model string, page, perPage int) (_ *PublicAPIListResponse[T], err error) {
	u := util.CloneRef(c.base)
	u.Path = path.Join(u.Path, "api", "p", c.project, model)
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		q.Set("per_page", strconv.Itoa(perPage))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return
	}

	res, err := c.c.Do(req)
	if err != nil {
		return
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			err = rerror.ErrNotFound
		} else {
			err = fmt.Errorf("invalid status code: %d", res.StatusCode)
		}
		return
	}

	var r PublicAPIListResponse[T]
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	r.Page = page
	r.PerPage = perPage
	return &r, nil
}

func (c *PublicAPIClient[T]) GetItem(ctx context.Context, model string, id string) (_ T, err error) {
	u := util.CloneRef(c.base)
	u.Path = path.Join(u.Path, "api", "p", c.project, model, id)
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return
	}

	res, err := c.c.Do(req)
	if err != nil {
		return
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			err = rerror.ErrNotFound
		} else {
			err = fmt.Errorf("invalid status code: %d", res.StatusCode)
		}
		return
	}

	var r T
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return
	}

	return r, nil
}

func (c *PublicAPIClient[T]) GetAsset(ctx context.Context, id string) (*PublicAsset, error) {
	u := util.CloneRef(c.base)
	u.Path = path.Join(u.Path, "api", "p", c.project, "assets", id)
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			err = rerror.ErrNotFound
		} else {
			err = fmt.Errorf("invalid status code: %d", res.StatusCode)
		}
		return nil, err
	}

	var r PublicAsset
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r, nil
}
