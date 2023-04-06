package dataconv

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/eukarya-inc/reearth-plateauview/server/cms/cmswebhook"
	"github.com/jarcoal/httpmock"
	"github.com/reearth/reearthx/rerror"
	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder("GET", borderURL, httpmock.NewStringResponder(http.StatusOK, border))
	c := &cmsMock{}

	// case1: skip conv
	err := webhookHandler(context.Background(), &cmswebhook.Payload{
		Type: cmswebhook.EventItemUpdate,
		ItemData: &cmswebhook.ItemData{
			Item: &cms.Item{
				ID: "xxx",
				Fields: Item{
					Type:       "行政界",
					DataFormat: "GeoJSON",
					Data:       "aaa",
					DataConv:   "変換しない",
				}.Fields(),
			},
			Model: &cms.Model{Key: "dataset"},
			Schema: &cms.Schema{
				ProjectID: "project",
			},
		},
		Operator: cmswebhook.Operator{User: &cmswebhook.User{}},
	}, Config{}, c)

	assert.NoError(t, err)
	assert.Nil(t, c.i)

	// case2: normal
	err = webhookHandler(context.Background(), &cmswebhook.Payload{
		Type: cmswebhook.EventItemUpdate,
		ItemData: &cmswebhook.ItemData{
			Item: &cms.Item{
				ID: "xxx",
				Fields: Item{
					Type:       "行政界",
					DataFormat: "GeoJSON",
					Data:       "aaa",
				}.Fields(),
			},
			Model: &cms.Model{Key: "dataset"},
			Schema: &cms.Schema{
				ProjectID: "project",
			},
		},
		Operator: cmswebhook.Operator{User: &cmswebhook.User{}},
	}, Config{}, c)

	assert.NoError(t, err)
	assert.Equal(t, &cms.Item{
		ID: "xxx",
		Fields: Item{
			DataFormat: "CZML",
			Data:       "asset",
		}.Fields(),
	}, c.i)
}

var borderURL = fmt.Sprintf("http://example.com/%s.geojson", borderName)

type cmsMock struct {
	cms.Interface
	i *cms.Item
}

func (c *cmsMock) UpdateItem(ctx context.Context, itemID string, fields []cms.Field) (*cms.Item, error) {
	c.i = &cms.Item{
		ID:     itemID,
		Fields: fields,
	}
	return nil, nil
}

func (c *cmsMock) Asset(ctx context.Context, id string) (*cms.Asset, error) {
	if id == "aaa" {
		return &cms.Asset{
			URL: borderURL,
		}, nil
	}
	return nil, rerror.ErrNotFound
}

func (c *cmsMock) UploadAssetDirectly(ctx context.Context, projectID, name string, data io.Reader) (string, error) {
	return "asset", nil
}
