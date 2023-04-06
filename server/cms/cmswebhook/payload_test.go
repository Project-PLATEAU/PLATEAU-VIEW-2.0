package cmswebhook

import (
	"encoding/json"
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/stretchr/testify/assert"
)

func TestPayload_UnmarshalJSON(t *testing.T) {
	p := Payload{}
	assert.NoError(t, json.Unmarshal([]byte(`{"type":"item.update","data":{"item":{"id":"i"}}}`), &p))
	assert.Equal(t, Payload{
		Type: "item.update",
		ItemData: &ItemData{
			Item: &cms.Item{
				ID: "i",
			},
		},
	}, p)

	p = Payload{}
	assert.NoError(t, json.Unmarshal([]byte(`{"type":"asset.decompress","data":{"id":"a"}}`), &p))
	assert.Equal(t, Payload{
		Type: "asset.decompress",
		AssetData: &AssetData{
			ID: "a",
		},
	}, p)
}
