package ckan

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/vincent-petithory/dataurl"
)

var _ Interface = (*Ckan)(nil)

func TestCkan(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	mockCkan(t)

	ctx := context.Background()
	ckan, err := New("https://www.geospatial.jp/ckan", "TOKEN")
	assert.NoError(t, err)

	p, err := ckan.ShowPackage(ctx, "plateau-tokyo23ku")
	assert.NoError(t, err)
	assert.Equal(t, Package{
		ID:       "xxx",
		Name:     "plateau-tokyo23ku",
		OwnerOrg: "yyy",
		Resources: []Resource{
			{ID: "a", URL: "https://example.com", PackageID: "xxx"},
		},
	}, p)

	s, err := ckan.SearchPackageByName(ctx, "plateau-tokyo23ku")
	assert.NoError(t, err)
	assert.Equal(t, List[Package]{
		Count: 100,
		Sort:  "hoge",
		Results: []Package{
			{
				ID:       "xxx",
				Name:     "plateau-tokyo23ku",
				OwnerOrg: "yyy",
				Resources: []Resource{
					{ID: "a", URL: "https://example.com", PackageID: "xxx"},
				},
			},
		},
	}, s)

	p, err = ckan.CreatePackage(ctx, Package{
		Name:     "plateau-tokyo23ku",
		OwnerOrg: "yyy",
	})
	assert.NoError(t, err)
	assert.Equal(t, Package{
		ID:       "xxx",
		Name:     "plateau-tokyo23ku",
		OwnerOrg: "yyy",
	}, p)

	p, err = ckan.PatchPackage(ctx, Package{
		ID:       "xxx",
		Name:     "plateau-tokyo23ku!",
		OwnerOrg: "yyy",
	})
	assert.NoError(t, err)
	assert.Equal(t, Package{
		ID:       "xxx",
		Name:     "plateau-tokyo23ku!",
		OwnerOrg: "yyy",
	}, p)

	r, err := ckan.CreateResource(ctx, Resource{
		Name: "aaa",
		URL:  "https://example.com",
	})
	assert.NoError(t, err)
	assert.Equal(t, Resource{
		ID:   "a",
		Name: "aaa",
		URL:  "https://example.com",
	}, r)

	r, err = ckan.PatchResource(ctx, Resource{
		ID:   "a",
		Name: "aaa!",
		URL:  "https://example.com",
	})
	assert.NoError(t, err)
	assert.Equal(t, Resource{
		ID:   "a",
		Name: "aaa!",
		URL:  "https://example.com",
	}, r)

	data := []byte("hello!")
	r, err = ckan.UploadResource(ctx, Resource{
		ID:          "aid",
		PackageID:   "pkgid",
		Description: "desc",
		Format:      "format",
		Name:        "name",
	}, "a.txt", data)
	assert.NoError(t, err)
	assert.Equal(t, Resource{
		ID:          "aid",
		PackageID:   "pkgid",
		Description: "desc",
		Format:      "format",
		Name:        "name",
		URL:         dataurl.New(data, http.DetectContentType(data)).String(),
	}, r)

	ckan.token = "xxx"
	r, err = ckan.PatchResource(ctx, Resource{
		ID:   "a",
		Name: "aaa!",
		URL:  "https://example.com",
	})
	assert.ErrorContains(t, err, "failed to patch a resource: status code 401: ")
	assert.Empty(t, r)
}

func mockCkan(t *testing.T) {
	t.Helper()

	checkAuth := func(req *http.Request) (*http.Response, error) {
		if req.Header.Get("X-CKAN-API-Key") != "TOKEN" {
			return httpmock.NewJsonResponse(http.StatusUnauthorized, Response[any]{Error: &Error{Message: "error"}})
		}
		return nil, nil
	}

	httpmock.RegisterResponderWithQuery("GET", "https://www.geospatial.jp/ckan/api/3/action/package_show", "id=plateau-tokyo23ku", func(req *http.Request) (*http.Response, error) {
		if res, err := checkAuth(req); res != nil {
			return res, err
		}

		return httpmock.NewJsonResponse(http.StatusOK, Response[Package]{
			Success: true,
			Result: Package{
				ID:       "xxx",
				Name:     "plateau-tokyo23ku",
				OwnerOrg: "yyy",
				Resources: []Resource{
					{ID: "a", URL: "https://example.com", PackageID: "xxx"},
				},
			},
		})
	})

	httpmock.RegisterResponderWithQuery("GET", "https://www.geospatial.jp/ckan/api/3/action/package_search", "q=name:plateau-tokyo23ku", func(req *http.Request) (*http.Response, error) {
		if res, err := checkAuth(req); res != nil {
			return res, err
		}

		return httpmock.NewJsonResponse(http.StatusOK, Response[List[Package]]{
			Success: true,
			Result: List[Package]{
				Count: 100,
				Sort:  "hoge",
				Results: []Package{
					{
						ID:       "xxx",
						Name:     "plateau-tokyo23ku",
						OwnerOrg: "yyy",
						Resources: []Resource{
							{ID: "a", URL: "https://example.com", PackageID: "xxx"},
						},
					},
				},
			},
		})
	})

	httpmock.RegisterResponder("POST", "https://www.geospatial.jp/ckan/api/3/action/package_create", func(req *http.Request) (*http.Response, error) {
		if res, err := checkAuth(req); res != nil {
			return res, err
		}

		res := Package{}
		_ = json.NewDecoder(req.Body).Decode(&res)
		res.ID = "xxx"

		return httpmock.NewJsonResponse(http.StatusOK, Response[Package]{
			Result: res,
		})
	})

	httpmock.RegisterResponder("POST", "https://www.geospatial.jp/ckan/api/3/action/package_patch", func(req *http.Request) (*http.Response, error) {
		if res, err := checkAuth(req); res != nil {
			return res, err
		}

		res := Package{}
		_ = json.NewDecoder(req.Body).Decode(&res)
		if res.ID == "" {
			return httpmock.NewJsonResponse(http.StatusBadRequest, Response[any]{Error: &Error{Message: "id missing"}})
		}

		return httpmock.NewJsonResponse(http.StatusOK, Response[Package]{
			Result: res,
		})
	})

	httpmock.RegisterResponder("POST", "https://www.geospatial.jp/ckan/api/3/action/resource_create", func(req *http.Request) (*http.Response, error) {
		if res, err := checkAuth(req); res != nil {
			return res, err
		}

		res := Resource{}
		_ = json.NewDecoder(req.Body).Decode(&res)
		res.ID = "a"

		return httpmock.NewJsonResponse(http.StatusOK, Response[Resource]{
			Result: res,
		})
	})

	httpmock.RegisterResponder("POST", "https://www.geospatial.jp/ckan/api/3/action/resource_patch", func(req *http.Request) (*http.Response, error) {
		if res, err := checkAuth(req); res != nil {
			return res, err
		}

		res := Resource{}

		if req.Header.Get("Content-Type") == "application/json" {
			_ = json.NewDecoder(req.Body).Decode(&res)
			if res.ID == "" {
				return httpmock.NewJsonResponse(http.StatusBadRequest, Response[any]{Error: &Error{Message: "id missing"}})
			}
		} else {
			if req.Header.Get("Content-Length") == "" {
				return httpmock.NewJsonResponse(http.StatusBadRequest, Response[any]{Error: &Error{Message: "content length required"}})
			}

			res.ID = req.FormValue("id")
			res.PackageID = req.FormValue("package_id")
			res.Name = req.FormValue("name")
			res.Description = req.FormValue("description")
			res.Format = req.FormValue("format")
			res.Mimetype = req.FormValue("mimetype")

			if m, _, err := req.FormFile("upload"); err == nil {
				b := bytes.NewBuffer(nil)
				_, _ = io.Copy(b, m)
				bb := b.Bytes()
				res.URL = dataurl.New(bb, http.DetectContentType(bb)).String()
			}
		}

		return httpmock.NewJsonResponse(http.StatusOK, Response[Resource]{
			Result: res,
		})
	})
}
