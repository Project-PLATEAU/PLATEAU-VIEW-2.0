package item

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/stretchr/testify/assert"
)

func TestNewField(t *testing.T) {
	f := id.NewFieldID()
	assert.Nil(t, NewField(f, nil))
	assert.Equal(t, &Field{
		field: f,
		value: value.TypeBool.Value(true).AsMultiple(),
	}, NewField(f, value.TypeBool.Value(true).AsMultiple()))
}
