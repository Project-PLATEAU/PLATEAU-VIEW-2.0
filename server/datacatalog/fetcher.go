package datacatalog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
)

const timeoutSecond int64 = 20
const ModelPlateau = "plateau"
const ModelUsecase = "usecase"
const ModelDataset = "dataset"

type Fetcher struct {
	c    *http.Client
	base *url.URL
}

func NewFetcher(c *http.Client, cmabse string) (*Fetcher, error) {
	u, err := url.Parse(cmabse)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "api", "p")

	if c == nil {
		c = http.DefaultClient
	}

	return &Fetcher{
		c:    c,
		base: u,
	}, nil
}

func (f *Fetcher) Clone() *Fetcher {
	if f == nil {
		return nil
	}
	return &Fetcher{
		c:    f.c,
		base: util.CloneRef(f.base),
	}
}

func (f *Fetcher) Do(ctx context.Context, project string) (ResponseAll, error) {
	f1, f2, f3 := f.Clone(), f.Clone(), f.Clone()

	res1 := lo.Async2(func() (ResponseAll, error) {
		return f1.all(ctx, project, ModelPlateau)
	})
	res2 := lo.Async2(func() (ResponseAll, error) {
		return f2.all(ctx, project, ModelUsecase)
	})
	res3 := lo.Async2(func() (ResponseAll, error) {
		return f3.all(ctx, project, ModelDataset)
	})

	notFound := 0
	r := ResponseAll{}

	if res := <-res1; res.B != nil {
		if errors.Is(res.B, rerror.ErrNotFound) {
			notFound++
		} else {
			return ResponseAll{}, res.B
		}
	} else {
		r.Plateau = append(r.Plateau, res.A.Plateau...)
		r.Usecase = append(r.Usecase, res.A.Usecase...)
	}

	if res := <-res2; res.B != nil {
		if errors.Is(res.B, rerror.ErrNotFound) {
			notFound++
		} else {
			return ResponseAll{}, res.B
		}
	} else {
		r.Plateau = append(r.Plateau, res.A.Plateau...)
		r.Usecase = append(r.Usecase, res.A.Usecase...)
	}

	if res := <-res3; res.B != nil {
		if errors.Is(res.B, rerror.ErrNotFound) {
			notFound++
		} else {
			return ResponseAll{}, res.B
		}
	} else {
		r.Plateau = append(r.Plateau, res.A.Plateau...)
		r.Usecase = append(r.Usecase, res.A.Usecase...)
	}

	if notFound == 3 {
		return r, rerror.ErrNotFound
	}
	return r, nil
}

func (f *Fetcher) all(ctx context.Context, project, model string) (resp ResponseAll, err error) {
	for p := 1; ; p++ {
		r, err := f.get(ctx, project, model, p, 0)
		if err != nil {
			return ResponseAll{}, err
		}

		resp.Plateau = append(resp.Plateau, r.Plateau...)
		resp.Usecase = append(resp.Usecase, r.Usecase...)
		if !r.HasNext() {
			break
		}
	}
	return
}

func (f *Fetcher) get(ctx context.Context, project, model string, page, perPage int) (r response, err error) {
	if f.c == nil {
		f.c = http.DefaultClient
	}

	r.Model = model
	if perPage == 0 {
		perPage = 100
	}

	u := f.url(project, model, page, perPage)
	log.Infof("datacatalog: get: %s", u)

	ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSecond)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx2, "GET", u, nil)
	if err != nil {
		return
	}

	res, err := f.c.Do(req)
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
			err = fmt.Errorf("status code: %d", res.StatusCode)
		}
		return
	}

	err = json.NewDecoder(res.Body).Decode(&r)
	r.Page = page
	r.PerPage = perPage
	return
}

func (f *Fetcher) url(project, model string, page, perPage int) string {
	u := util.CloneRef(f.base)
	u.Path = path.Join(u.Path, project, model)
	u.RawQuery = url.Values{
		"page":     []string{strconv.Itoa(page)},
		"per_page": []string{strconv.Itoa(perPage)},
	}.Encode()
	return u.String()
}

type response responseInternal

type responseInternal struct {
	Model      string          `json:"-"`
	Results    json.RawMessage `json:"results"`
	Plateau    []PlateauItem   `json:"-"`
	Usecase    []UsecaseItem   `json:"-"`
	Page       int             `json:"page"`
	PerPage    int             `json:"perPage"`
	TotalCount int             `json:"totalCount"`
}

func (r *response) UnmarshalJSON(data []byte) error {
	r2 := responseInternal{}
	if err := json.Unmarshal(data, &r2); err != nil {
		return err
	}

	if r.Model == ModelPlateau {
		if err := json.Unmarshal(r2.Results, &r2.Plateau); err != nil {
			return err
		}
	} else if r.Model == ModelUsecase || r.Model == ModelDataset {
		if err := json.Unmarshal(r2.Results, &r2.Usecase); err != nil {
			return err
		}
	}

	if r.Model == ModelUsecase {
		r2.Usecase = lo.Map(r2.Usecase, func(r UsecaseItem, _ int) UsecaseItem {
			r.Type = "ユースケース"
			return r
		})
	}

	r2.Results = nil
	*r = response(r2)
	return nil
}

func (r response) HasNext() bool {
	if r.PerPage == 0 {
		return false
	}
	return r.TotalCount > r.Page*r.PerPage
}

func (r response) DataCatalogs() []DataCatalogItem {
	if r.Plateau != nil {
		return lo.FlatMap(r.Plateau, func(i PlateauItem, _ int) []DataCatalogItem {
			return i.DataCatalogItems()
		})
	}
	if r.Usecase != nil {
		return lo.FlatMap(r.Usecase, func(i UsecaseItem, _ int) []DataCatalogItem {
			return i.DataCatalogs()
		})
	}
	return nil
}
