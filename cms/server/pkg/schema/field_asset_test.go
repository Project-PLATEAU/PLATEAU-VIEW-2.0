package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/stretchr/testify/assert"
)

func TestNewAsset(t *testing.T) {
	assert.Equal(t, &FieldAsset{}, NewAsset())
}

func TestFieldAsset_Type(t *testing.T) {
	assert.Equal(t, value.TypeAsset, (&FieldAsset{}).Type())
}

func TestFieldAsset_Clone(t *testing.T) {
	assert.Nil(t, (*FieldAsset)(nil).Clone())
	assert.Equal(t, &FieldAsset{}, (&FieldAsset{}).Clone())
}

func TestFieldAsset_Validate(t *testing.T) {
	aid := id.NewAssetID()
	assert.NoError(t, (&FieldAsset{}).Validate(value.TypeAsset.Value(aid)))
	assert.Equal(t, ErrInvalidValue, (&FieldAsset{}).Validate(value.TypeText.Value("")))
}
