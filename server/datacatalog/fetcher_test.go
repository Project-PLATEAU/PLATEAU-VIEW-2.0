package datacatalog

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

// func TestFetcher(t *testing.T) {
// 	f := lo.Must(NewFetcher(nil, ""))
// 	cmsres := lo.Must(f.Do(context.Background(), ""))
// 	item, _ := lo.Find(cmsres.Plateau, func(i PlateauItem) bool { return i.CityName == "" })
// 	res := item.GenItems(item.IntermediateItem())
// 	t.Log(string(lo.Must(json.MarshalIndent(res, "", "  "))))
// }

func TestFetcher_Do(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponderWithQuery("GET", "https://example.com/plateau", "page=1&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results":    []any{map[string]string{"id": "x"}},
		"totalCount": 1,
	})))
	httpmock.RegisterResponderWithQuery("GET", "https://example.com/usecase", "page=1&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results":    []any{map[string]string{"id": "y"}},
		"totalCount": 1,
	})))
	httpmock.RegisterResponderWithQuery("GET", "https://example.com/dataset", "page=1&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results":    []any{map[string]string{"id": "z"}},
		"totalCount": 1,
	})))

	ctx := context.Background()
	r, err := (&Fetcher{base: lo.Must(url.Parse("https://example.com"))}).Do(ctx, "")
	assert.Equal(t, ResponseAll{
		Plateau: []PlateauItem{{ID: "x"}},
		Usecase: []UsecaseItem{{ID: "y", Type: "ユースケース"}, {ID: "z"}},
	}, r)
	assert.NoError(t, err)
}

func TestFetcher_all(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponderWithQuery("GET", "https://example.com/plateau", "page=1&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results":    []any{map[string]string{"id": "x"}},
		"totalCount": 101,
	})))
	httpmock.RegisterResponderWithQuery("GET", "https://example.com/plateau", "page=2&per_page=100", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results":    []any{map[string]string{"id": "y"}},
		"totalCount": 101,
	})))

	ctx := context.Background()
	r, err := (&Fetcher{base: lo.Must(url.Parse("https://example.com"))}).all(ctx, "", "plateau")
	assert.Equal(t, ResponseAll{
		Plateau: []PlateauItem{
			{ID: "x"}, {ID: "y"},
		},
	}, r)
	assert.NoError(t, err)
}

func TestFetcher_get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder("GET", "https://example.com/plateau", lo.Must(httpmock.NewJsonResponder(http.StatusOK, map[string]any{
		"results":    []any{map[string]string{"id": "x"}},
		"totalCount": 100,
	})))

	ctx := context.Background()
	r, err := (&Fetcher{base: lo.Must(url.Parse("https://example.com"))}).get(ctx, "", "plateau", 1, 2)
	assert.Equal(t, response{
		Plateau:    []PlateauItem{{ID: "x"}},
		Page:       1,
		PerPage:    2,
		TotalCount: 100,
	}, r)
	assert.NoError(t, err)
}

func TestFetcher_url(t *testing.T) {
	assert.Equal(t, "https://example.com/b/a?page=1&per_page=2", (&Fetcher{base: lo.Must(url.Parse("https://example.com"))}).url("b", "a", 1, 2))
}

func TestResponse_UnmarshalJSON(t *testing.T) {
	assert := assert.New(t)
	r := response{
		Model: "plateau",
	}
	b := `{"results":[{"id":"xxx"}],"totalCount":100}`
	assert.NoError(json.Unmarshal([]byte(b), &r))
	assert.Equal(response{
		TotalCount: 100,
		Plateau: []PlateauItem{
			{ID: "xxx"},
		},
	}, r)
}

func TestResponse_HasNext(t *testing.T) {
	assert.True(t, response{Page: 1, PerPage: 50, TotalCount: 100}.HasNext())
	assert.False(t, response{Page: 2, PerPage: 50, TotalCount: 100}.HasNext())
	assert.True(t, response{Page: 1, PerPage: 10, TotalCount: 11}.HasNext())
	assert.False(t, response{Page: 2, PerPage: 10, TotalCount: 11}.HasNext())
}
