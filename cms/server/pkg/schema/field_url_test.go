package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/stretchr/testify/assert"
)

func TestNewURL(t *testing.T) {
	assert.Equal(t, &FieldURL{}, NewURL())
}

func TestFieldURL_Type(t *testing.T) {
	assert.Equal(t, value.TypeURL, (&FieldURL{}).Type())
}

func TestFieldURL_Clone(t *testing.T) {
	assert.Nil(t, (*FieldURL)(nil).Clone())
	assert.Equal(t, &FieldURL{}, (&FieldURL{}).Clone())
}

func TestFieldURL_Validate(t *testing.T) {
	assert.NoError(t, (&FieldURL{}).Validate(value.TypeURL.Value("https://example.com")))
	assert.Equal(t, ErrInvalidValue, (&FieldURL{}).Validate(value.TypeText.Value("")))
}
