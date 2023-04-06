package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/stretchr/testify/assert"
)

func TestNewBool(t *testing.T) {
	assert.Equal(t, &FieldBool{}, NewBool())
}

func TestFieldBool_Type(t *testing.T) {
	assert.Equal(t, value.TypeBool, (&FieldBool{}).Type())
}

func TestFieldBool_Clone(t *testing.T) {
	assert.Nil(t, (*FieldBool)(nil).Clone())
	assert.Equal(t, &FieldBool{}, (&FieldBool{}).Clone())
}

func TestFieldBool_Validate(t *testing.T) {
	assert.NoError(t, (&FieldBool{}).Validate(value.TypeBool.Value(true)))
	assert.Equal(t, ErrInvalidValue, (&FieldBool{}).Validate(value.TypeText.Value("")))
}
