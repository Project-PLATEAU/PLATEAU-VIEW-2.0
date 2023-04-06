package fme

import "net/url"

type MaxLODRequest struct {
	ID     string
	Target string
}

func (r MaxLODRequest) Query() url.Values {
	q := url.Values{}
	q.Set("id", r.ID)
	q.Set("url", r.Target)
	return q
}

func (r MaxLODRequest) Name() string {
	return "plateau2022-cms/maxlod-extract"
}
