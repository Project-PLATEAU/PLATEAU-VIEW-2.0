package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/stretchr/testify/assert"
)

func TestNewSelect(t *testing.T) {
	assert.Equal(t, &FieldSelect{values: []string{"a", "b"}}, NewSelect([]string{"a", "b", " a", "b"}))
}

func TestFieldSelect_Type(t *testing.T) {
	assert.Equal(t, value.TypeSelect, (&FieldSelect{}).Type())
}

func TestFieldSelect_Clone(t *testing.T) {
	assert.Nil(t, (*FieldSelect)(nil).Clone())
	assert.Equal(t, &FieldSelect{}, (&FieldSelect{}).Clone())
}

func TestFieldSelect_Validate(t *testing.T) {
	assert.NoError(t, (&FieldSelect{values: []string{"aaa"}}).Validate(value.TypeSelect.Value("aaa")))
	assert.Equal(t, ErrInvalidValue, (&FieldSelect{values: []string{"aa"}}).Validate(value.TypeSelect.Value("aaa")))
	assert.Equal(t, ErrInvalidValue, (&FieldSelect{}).Validate(value.TypeText.Value("")))
}
