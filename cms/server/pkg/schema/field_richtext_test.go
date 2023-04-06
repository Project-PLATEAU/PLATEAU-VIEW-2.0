package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewRichText(t *testing.T) {
	assert.Equal(t, &FieldRichText{s: &FieldString{t: value.TypeRichText, maxLength: lo.ToPtr(1)}}, NewRichText(lo.ToPtr(1)))
}

func TestFieldRichText_Type(t *testing.T) {
	assert.Equal(t, value.TypeRichText, (&FieldRichText{s: &FieldString{t: value.TypeRichText}}).Type())
}

func TestFieldRichText_Clone(t *testing.T) {
	assert.Nil(t, (*FieldRichText)(nil).Clone())
	assert.Equal(t, &FieldRichText{}, (&FieldRichText{}).Clone())
}

func TestFieldRichText_Validate(t *testing.T) {
	assert.NoError(t, (&FieldRichText{s: &FieldString{t: value.TypeRichText}}).Validate(value.TypeRichText.Value("aaa")))
	assert.Equal(t, ErrInvalidValue, (&FieldRichText{s: &FieldString{t: value.TypeRichText}}).Validate(value.TypeText.Value("")))
}
