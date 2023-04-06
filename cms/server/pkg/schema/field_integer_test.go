package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewInteger(t *testing.T) {
	got, err := NewInteger(lo.ToPtr(int64(1)), lo.ToPtr(int64(2)))
	assert.Equal(t, &FieldInteger{
		min: lo.ToPtr(int64(1)),
		max: lo.ToPtr(int64(2)),
	}, got)
	assert.NoError(t, err)
	got, err = NewInteger(lo.ToPtr(int64(3)), lo.ToPtr(int64(2)))
	assert.Nil(t, got)
	assert.Equal(t, ErrInvalidMinMax, err)
}

func TestFieldInteger_Type(t *testing.T) {
	assert.Equal(t, value.TypeInteger, (&FieldInteger{}).Type())
}

func TestFieldInteger_Clone(t *testing.T) {
	assert.Nil(t, (*FieldInteger)(nil).Clone())
	assert.Equal(t, &FieldInteger{
		min: lo.ToPtr(int64(1)),
		max: lo.ToPtr(int64(2)),
	}, (&FieldInteger{
		min: lo.ToPtr(int64(1)),
		max: lo.ToPtr(int64(2)),
	}).Clone())
}

func TestFieldInteger_Validate(t *testing.T) {
	assert.NoError(t, (&FieldInteger{}).Validate(value.TypeInteger.Value(1)))
	assert.Equal(t, ErrInvalidValue, (&FieldInteger{}).Validate(value.TypeText.Value("")))
}
