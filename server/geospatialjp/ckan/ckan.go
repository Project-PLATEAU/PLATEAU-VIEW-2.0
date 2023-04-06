package ckan

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/reearth/reearthx/log"
)

type Interface interface {
	ShowPackage(ctx context.Context, id string) (Package, error)
	SearchPackageByName(ctx context.Context, name string) (List[Package], error)
	CreatePackage(ctx context.Context, pkg Package) (Package, error)
	PatchPackage(ctx context.Context, pkg Package) (Package, error)
	SavePackage(ctx context.Context, pkg Package) (Package, error)
	CreateResource(ctx context.Context, resource Resource) (Resource, error)
	PatchResource(ctx context.Context, resource Resource) (Resource, error)
	UploadResource(ctx context.Context, resource Resource, filename string, data []byte) (Resource, error)
	SaveResource(ctx context.Context, resource Resource) (Resource, error)
}

type Ckan struct {
	base   *url.URL
	token  string
	client *http.Client
}

type Response[T any] struct {
	Help    string `json:"help,omitempty"`
	Success bool   `json:"success,omitempty"`
	Error   *Error `json:"error,omitempty"`
	Result  T      `json:"result,omitempty"`
}

type List[T any] struct {
	Count   int    `json:"count,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Results []T    `json:"results,omitempty"`
	// facets
	// search_facets
}

func (l List[T]) IsEmpty() bool {
	return len(l.Results) == 0
}

type Error struct {
	Message string `json:"message,omitempty"`
	Type    string `json:"__type,omitempty"`
}

type Package struct {
	ID string `json:"id,omitempty"`
	// The name of the new dataset, must be between 2 and 100 characters long and
	// contain only lowercase alphanumeric characters, - and _, e.g. 'warandpeace'
	Name string `json:"name,omitempty"`
	// The title of the dataset (optional, default: same as name)
	Title string `json:"title,omitempty"`
	// If True creates a private dataset
	Private bool `json:"private,omitempty"`
	// The name of the dataset's author (optional)
	Author string `json:"author,omitempty"`
	// The email address of the dataset's author (optional)
	AuthorEmail string `json:"author_email,omitempty"`
	// The name of the dataset's maintainer (optional)
	Maintainer string `json:"maintainer,omitempty"`
	// The email address of the dataset's maintainer (optional)
	MaintainerEmail string `json:"maintainer_email,omitempty"`
	// The id of the dataset's license, see `~ckan.logic.action.get.license_list` for available values (optional)
	LicenseID string `json:"license_id,omitempty"`
	// A description of the dataset (optional)
	Notes string `json:"notes,omitempty"`
	// A URL for the dataset's source (optional)
	URL string `json:"url,omitempty"`
	// (optional)
	Version string `json:"version,omitempty"`
	// The current state of the dataset, e.g. 'active' or
	// 'deleted', only active datasets show up in search results and
	// other lists of datasets, this parameter will be ignored if you are not
	// authorized to change the state of the dataset (optional, default:
	// 'active')
	State string `json:"state,omitempty"`
	// The type of the dataset (optional),
	// `~ckan.plugins.interfaces.IDatasetForm` plugins
	// associate themselves with different dataset types and provide custom
	// dataset handling behaviour for these types
	Type string `json:"type,omitempty"`
	// The dataset's resources, see
	// `resource_create` for the format of resource dictionaries (optional)
	Resources []Resource `json:"resources,omitempty"`
	// The dataset's tags, see `tag_create` for the format
	// of tag dictionaries (optional)
	Tags []Tag `json:"tags,omitempty"`

	// The dataset's extras (optional), extras are arbitrary
	// (key: value) metadata items that can be added to datasets, each extra
	// dictionary should have keys 'key' (a string), 'value' (a string)
	// Extras map[string]string `json:"extras,omitempty"`

	// relationships_as_object: See `package_relationship_create`
	// for the format of relationship dictionaries (optional)
	//  relationships_as_object: list of relationship dictionaries
	// relationships_as_object []? `json:"relationships_as_object,omitempty"`

	// relationships_as_subject: See `package_relationship_create`
	// for the format of relationship dictionaries (optional)
	//  relationships_as_subject: list of relationship dictionaries
	// relationships_as_subject []? `json:"relationships_as_subject,omitempty"`

	// groups: The groups to which the dataset belongs (optional), each
	// group dictionary should have one or more of the following keys which
	// identify an existing group:
	// 'id' (the id of the group, string), or 'name' (the name of the
	// group, string),  to see which groups exist
	// call `~ckan.logic.action.get.group_list`
	//  groups: list of dictionaries
	// groups []? `json:"groups,omitempty"`

	// The id of the dataset's owning organization, see
	// `~ckan.logic.action.get.organization_list` or
	// `~ckan.logic.action.get.organization_list_for_user` for
	// available values. This parameter can be made optional if the config
	// option `ckan.auth.create_unowned_dataset` is set to True.
	// Note: name is also available instead of ID on creating a package
	OwnerOrg string `json:"owner_org,omitempty"`

	// geospatialjp
	Restriction      string `json:"restriction,omitempty"`
	LicenseAgreement string `json:"license_agreement,omitempty"`
	RegisterdDate    string `json:"registerd_date,omitempty"`
	Fee              string `json:"fee,omitempty"`
	Charge           string `json:"charge,omitempty"`
	Area             string `json:"area,omitempty"`
	Quality          string `json:"quality,omitempty"`
	Emergency        string `json:"emergency,omitempty"`
	LicenseTitle     string `json:"license_title,omitempty"`
	ThumbnailURL     string `json:"thumbnail_url,omitempty"`
	LicenseURL       string `json:"license_url,omitempty"`
	Spatial          string `json:"spatial,omitempty"`
}

type Resource struct {
	ID string `json:"id,omitempty"`
	// id of package that the resource should be added to.
	PackageID string `json:"package_id,omitempty"`
	// url of resource
	URL              string `json:"url,omitempty"`
	URLType          string `json:"url_type,omitempty"`
	RevisionID       string `json:"revision_id,omitempty"`
	Description      string `json:"description,omitempty"`
	Format           string `json:"format,omitempty"`
	Hash             string `json:"hash,omitempty"`
	Name             string `json:"name,omitempty"`
	ResourceType     string `json:"resource_type,omitempty"`
	Mimetype         string `json:"mimetype,omitempty"`
	MimetypeInner    string `json:"mimetype_inner,omitempty"`
	CacheUrl         string `json:"cache_url,omitempty"`
	Size             int    `json:"size,omitempty"`
	Created          string `json:"created,omitempty"`
	LastModified     string `json:"last_modified,omitempty"`
	CacheLastUpdated string `json:"cache_last_updated,omitempty"`
}

func (r Resource) WriteMultipart(m *multipart.Writer) error {
	if r.ID != "" {
		if err := m.WriteField("id", r.ID); err != nil {
			return err
		}
	}
	if r.PackageID != "" {
		if err := m.WriteField("package_id", r.PackageID); err != nil {
			return err
		}
	}
	if r.URL != "" {
		if err := m.WriteField("url", r.URL); err != nil {
			return err
		}
	}
	if r.Description != "" {
		if err := m.WriteField("description", r.Description); err != nil {
			return err
		}
	}
	if r.Format != "" {
		if err := m.WriteField("format", r.Format); err != nil {
			return err
		}
	}
	if r.Name != "" {
		if err := m.WriteField("name", r.Name); err != nil {
			return err
		}
	}
	if r.Mimetype != "" {
		if err := m.WriteField("mimetype", r.Mimetype); err != nil {
			return err
		}
	}
	return nil
}

type Tag struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
	State        string `json:"state,omitempty"`
	VocabularyID string `json:"vocabulary_id,omitempty"`
}

func New(base, token string) (*Ckan, error) {
	b, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	return &Ckan{base: b, token: token, client: http.DefaultClient}, nil
}

func (c *Ckan) ShowPackage(ctx context.Context, id string) (Package, error) {
	res := Response[Package]{}
	err := c.send(ctx, "GET", []string{"api", "3", "action", "package_show"}, map[string]string{
		"id": id,
	}, "", 0, nil, &res)
	if err != nil {
		return res.Result, fmt.Errorf("failed to get a package: %w", err)
	}

	return res.Result, nil
}

func (c *Ckan) SearchPackageByName(ctx context.Context, name string) (List[Package], error) {
	res := Response[List[Package]]{}
	err := c.send(ctx, "GET", []string{"api", "3", "action", "package_search"}, map[string]string{
		"q": fmt.Sprintf("name:%s", name),
	}, "", 0, nil, &res)
	if err != nil {
		return res.Result, fmt.Errorf("failed to get a package: %w", err)
	}

	return res.Result, nil
}

func (c *Ckan) CreatePackage(ctx context.Context, pkg Package) (Package, error) {
	res := Response[Package]{}

	b, err := json.Marshal(pkg)
	if err != nil {
		return res.Result, fmt.Errorf("failed to marshal: %w", err)
	}

	if err = c.send(ctx, "POST", []string{"api", "3", "action", "package_create"}, nil, "", 0, bytes.NewReader(b), &res); err != nil {
		return res.Result, fmt.Errorf("failed to create a package: %w", err)
	}

	return res.Result, nil
}

func (c *Ckan) PatchPackage(ctx context.Context, pkg Package) (Package, error) {
	res := Response[Package]{}

	b, err := json.Marshal(pkg)
	if err != nil {
		return res.Result, fmt.Errorf("failed to marshal: %w", err)
	}

	err = c.send(ctx, "POST", []string{"api", "3", "action", "package_patch"}, nil, "", 0, bytes.NewReader(b), &res)
	if err != nil {
		return res.Result, fmt.Errorf("failed to patch a package: %w", err)
	}

	return res.Result, nil
}

func (c *Ckan) SavePackage(ctx context.Context, pkg Package) (Package, error) {
	if pkg.ID == "" {
		return c.CreatePackage(ctx, pkg)
	}
	return c.PatchPackage(ctx, pkg)
}

func (c *Ckan) CreateResource(ctx context.Context, resource Resource) (Resource, error) {
	res := Response[Resource]{}

	b, err := json.Marshal(resource)
	if err != nil {
		return res.Result, fmt.Errorf("failed to marshal: %w", err)
	}

	err = c.send(ctx, "POST", []string{"api", "3", "action", "resource_create"}, nil, "", 0, bytes.NewReader(b), &res)
	if err != nil {
		return res.Result, fmt.Errorf("failed to create a resource: %w", err)
	}

	return res.Result, nil
}

func (c *Ckan) PatchResource(ctx context.Context, resource Resource) (Resource, error) {
	res := Response[Resource]{}

	b, err := json.Marshal(resource)
	if err != nil {
		return res.Result, fmt.Errorf("failed to marshal: %w", err)
	}

	err = c.send(ctx, "POST", []string{"api", "3", "action", "resource_patch"}, nil, "", 0, bytes.NewReader(b), &res)
	if err != nil {
		return res.Result, fmt.Errorf("failed to patch a resource: %w", err)
	}

	return res.Result, nil
}

func (c *Ckan) UploadResource(ctx context.Context, resource Resource, filename string, data []byte) (Resource, error) {
	if filename == "" {
		return Resource{}, errors.New("filename missing")
	}

	res := Response[Resource]{}

	b := bytes.NewBuffer(nil)
	mw := multipart.NewWriter(b)
	if err := resource.WriteMultipart(mw); err != nil {
		return res.Result, fmt.Errorf("failed to upload a resource: %w", err)
	}
	fw, err := mw.CreateFormFile("upload", filename)
	if err != nil {
		return res.Result, fmt.Errorf("failed to upload a resource: %w", err)
	}
	if _, err = fw.Write(data); err != nil {
		return res.Result, fmt.Errorf("failed to upload a resource: %w", err)
	}
	mw.Close()

	var act string
	if resource.ID == "" {
		act = "create"
	} else {
		act = "patch"
	}

	if err := c.send(ctx, "POST", []string{"api", "3", "action", "resource_" + act}, nil, mw.FormDataContentType(), b.Len(), b, &res); err != nil {
		return res.Result, fmt.Errorf("failed to upload a resource: %w", err)
	}

	return res.Result, nil
}

func (c *Ckan) SaveResource(ctx context.Context, resource Resource) (Resource, error) {
	if resource.ID == "" {
		return c.CreateResource(ctx, resource)
	}
	return c.PatchResource(ctx, resource)
}

func (c *Ckan) send(
	ctx context.Context,
	method string,
	path []string,
	queries map[string]string,
	contentType string,
	contentLength int,
	body io.Reader,
	result any,
) error {
	u := c.base.JoinPath(path...)
	if queries != nil {
		q := u.Query()
		for k, v := range queries {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if contentType == "" {
		contentType = "application/json"
	} else if strings.HasPrefix(contentType, "multipart/form-data") {
		req.Header.Set("Content-Length", strconv.Itoa(contentLength))
	}
	req.Header.Set("Content-Type", contentType)
	if c.token != "" {
		req.Header.Set("X-CKAN-API-Key", c.token)
	}

	log.Infof("ckan: send: %s %s", method, u)
	res, err := c.client.Do(req)
	if err != nil {
		log.Errorf("ckan: send error: %s", err)
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer func() { _ = res.Body.Close() }()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("ckan: result (%d): failed to read response body")
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode != 200 {
		var msg any = b
		var eres Response[any]
		if err := json.Unmarshal(b, &eres); err == nil && eres.Error.Message != "" {
			msg = eres.Error.Message
		}

		log.Infof("ckan: result (%d): %s", res.StatusCode, msg)
		return fmt.Errorf("status code %d: %s", res.StatusCode, msg)
	}

	if result != nil {
		if err := json.Unmarshal(b, result); err != nil {
			return fmt.Errorf("failed to parse JSON: %w", err)
		}
	}

	log.Debugf("ckan: ok: %s", b)

	return nil
}
