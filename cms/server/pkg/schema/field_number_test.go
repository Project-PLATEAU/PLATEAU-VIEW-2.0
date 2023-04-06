package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewNumber(t *testing.T) {
	got, err := NewNumber(lo.ToPtr(float64(1)), lo.ToPtr(float64(2)))
	assert.Equal(t, &FieldNumber{
		min: lo.ToPtr(float64(1)),
		max: lo.ToPtr(float64(2)),
	}, got)
	assert.NoError(t, err)
	got, err = NewNumber(lo.ToPtr(float64(3)), lo.ToPtr(float64(2)))
	assert.Nil(t, got)
	assert.Equal(t, ErrInvalidMinMax, err)
}

func TestFieldNumber_Type(t *testing.T) {
	assert.Equal(t, value.TypeNumber, (&FieldNumber{}).Type())
}

func TestFieldNumber_Clone(t *testing.T) {
	assert.Nil(t, (*FieldNumber)(nil).Clone())
	assert.Equal(t, &FieldNumber{
		min: lo.ToPtr(float64(1)),
		max: lo.ToPtr(float64(2)),
	}, (&FieldNumber{
		min: lo.ToPtr(float64(1)),
		max: lo.ToPtr(float64(2)),
	}).Clone())
}

func TestFieldNumber_Validate(t *testing.T) {
	assert.NoError(t, (&FieldNumber{}).Validate(value.TypeNumber.Value(1)))
	assert.ErrorContains(t,
		(&FieldNumber{min: lo.ToPtr(float64(1.1))}).Validate(value.TypeNumber.Value(1.01)),
		"value should be larger than 1.100000")
	assert.ErrorContains(t,
		(&FieldNumber{max: lo.ToPtr(float64(1.1))}).Validate(value.TypeNumber.Value(1.11)),
		"value should be smaller than 1.100000")
	assert.Equal(t, ErrInvalidValue, (&FieldNumber{}).Validate(value.TypeText.Value("")))
}
