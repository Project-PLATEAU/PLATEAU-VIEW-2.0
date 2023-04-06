package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewString(t *testing.T) {
	assert.Equal(t, &FieldString{t: value.TypeText, maxLength: lo.ToPtr(1)}, NewString(value.TypeText, lo.ToPtr(1)))
}

func TestFieldString_Type(t *testing.T) {
	assert.Equal(t, value.TypeText, (&FieldString{t: value.TypeText}).Type())
}

func TestFieldString_Clone(t *testing.T) {
	assert.Nil(t, (*FieldString)(nil).Clone())
	assert.Equal(t, &FieldString{t: value.TypeText, maxLength: lo.ToPtr(1)}, (&FieldString{t: value.TypeText, maxLength: lo.ToPtr(1)}).Clone())
}

func TestFieldString_Validate(t *testing.T) {
	assert.ErrorContains(t,
		(&FieldString{t: value.TypeText, maxLength: lo.ToPtr(1)}).Validate(value.TypeText.Value("aaa")),
		"value has 3 characters, but it sholud be shorter than 1 characters")
	assert.ErrorContains(t,
		(&FieldString{t: value.TypeText, maxLength: lo.ToPtr(1)}).Validate(value.TypeText.Value("ああ")),
		"value has 2 characters, but it sholud be shorter than 1 characters")
	assert.NoError(t, (&FieldString{t: value.TypeText, maxLength: lo.ToPtr(4)}).Validate(value.TypeText.Value("あ")))
	assert.NoError(t, (&FieldString{t: value.TypeText}).Validate(value.TypeText.Value("aaa")))
	assert.Equal(t, ErrInvalidValue, (&FieldString{t: value.TypeText}).Validate(value.TypeNumber.Value(1)))
}
