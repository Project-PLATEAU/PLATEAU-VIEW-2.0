package mongodoc

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/stretchr/testify/assert"
)

func TestNewValue(t *testing.T) {
	assert.Nil(t, NewValue(nil))
	assert.Equal(t, &ValueDocument{
		T: "bool",
		V: true,
	}, NewValue(value.TypeBool.Value(true)))
}

func TestNewOptionalValue(t *testing.T) {
	assert.Nil(t, NewOptionalValue(nil))
	assert.Equal(t, &ValueDocument{
		T: "bool",
	}, NewOptionalValue(value.TypeBool.None()))
	assert.Equal(t, &ValueDocument{
		T: "bool",
		V: true,
	}, NewOptionalValue(value.TypeBool.Value(true).Some()))
}

func TestNewMultipleValue(t *testing.T) {
	assert.Nil(t, NewMultipleValue(nil))
	assert.Equal(t, &ValueDocument{
		T: "bool",
		V: []any{},
	}, NewMultipleValue(value.MultipleFrom(value.TypeBool, nil)))
	assert.Equal(t, &ValueDocument{
		T: "bool",
		V: []any{true},
	}, NewMultipleValue(value.MultipleFrom(value.TypeBool, []*value.Value{value.TypeBool.Value(true)})))
}
