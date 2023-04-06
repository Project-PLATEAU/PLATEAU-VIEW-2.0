package value

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/stretchr/testify/assert"
)

func Test_propertyReference_ToValue(t *testing.T) {
	a := id.NewItemID()

	tests := []struct {
		name  string
		args  []any
		want1 any
		want2 bool
	}{
		{
			name:  "string",
			args:  []any{a.String(), a.StringRef()},
			want1: a,
			want2: true,
		},
		{
			name:  "id",
			args:  []any{a, &a},
			want1: a,
			want2: true,
		},
		{
			name:  "nil",
			args:  []any{(*string)(nil), nil},
			want1: nil,
			want2: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := &propertyReference{}
			for i, v := range tt.args {
				got1, got2 := p.ToValue(v)
				assert.Equal(t, tt.want1, got1, "test %d", i)
				assert.Equal(t, tt.want2, got2, "test %d", i)
			}
		})
	}
}

func Test_propertyReference_ToInterface(t *testing.T) {
	a := id.NewItemID()
	tt, ok := (&propertyReference{}).ToInterface(a)
	assert.Equal(t, a.String(), tt)
	assert.Equal(t, true, ok)
}

func Test_propertyReference_IsEmpty(t *testing.T) {
	assert.True(t, (&propertyReference{}).IsEmpty(id.ItemID{}))
	assert.False(t, (&propertyReference{}).IsEmpty(id.NewItemID()))
}

func Test_propertyReference_Validate(t *testing.T) {
	a := id.NewItemID()
	assert.True(t, (&propertyReference{}).Validate(a))
}
