package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewText(t *testing.T) {
	assert.Equal(t, &FieldText{s: &FieldString{t: value.TypeText, maxLength: lo.ToPtr(1)}}, NewText(lo.ToPtr(1)))
}

func TestFieldText_Type(t *testing.T) {
	assert.Equal(t, value.TypeText, (&FieldText{s: &FieldString{t: value.TypeText}}).Type())
}

func TestFieldText_Clone(t *testing.T) {
	assert.Nil(t, (*FieldText)(nil).Clone())
	assert.Equal(t, &FieldText{}, (&FieldText{}).Clone())
}

func TestFieldText_Validate(t *testing.T) {
	assert.NoError(t, (&FieldText{s: &FieldString{t: value.TypeText}}).Validate(value.TypeText.Value("aaa")))
	assert.Equal(t, ErrInvalidValue, (&FieldText{s: &FieldString{t: value.TypeText}}).Validate(value.TypeTextArea.Value("")))
}
