package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewTextArea(t *testing.T) {
	assert.Equal(t, &FieldTextArea{s: &FieldString{t: value.TypeTextArea, maxLength: lo.ToPtr(1)}}, NewTextArea(lo.ToPtr(1)))
}

func TestFieldTextArea_Type(t *testing.T) {
	assert.Equal(t, value.TypeTextArea, (&FieldTextArea{s: &FieldString{t: value.TypeTextArea}}).Type())
}

func TestFieldTextArea_Clone(t *testing.T) {
	assert.Nil(t, (*FieldTextArea)(nil).Clone())
	assert.Equal(t, &FieldTextArea{}, (&FieldTextArea{}).Clone())
}

func TestFieldTextArea_Validate(t *testing.T) {
	assert.NoError(t, (&FieldTextArea{s: &FieldString{t: value.TypeTextArea}}).Validate(value.TypeTextArea.Value("aaa")))
	assert.Equal(t, ErrInvalidValue, (&FieldTextArea{s: &FieldString{t: value.TypeTextArea}}).Validate(value.TypeText.Value("")))
}
