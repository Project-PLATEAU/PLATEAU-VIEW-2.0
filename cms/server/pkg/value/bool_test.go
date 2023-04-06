package value

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func Test_propertyBool_ToValue(t *testing.T) {
	tests := []struct {
		name  string
		args  []any
		want1 any
		want2 bool
	}{
		{
			name: "true",
			args: []any{
				true, "true", "TRUE", "True", "T", "t", "1", 1,
				lo.ToPtr(true), lo.ToPtr("true"), lo.ToPtr("TRUE"), lo.ToPtr("True"),
				lo.ToPtr("T"), lo.ToPtr("t"), lo.ToPtr("1"), lo.ToPtr(1),
			},
			want1: true,
			want2: true,
		},
		{
			name: "false",
			args: []any{
				false, "false", "FALSE", "False", "F", "f", "0", 0,
				lo.ToPtr(false), lo.ToPtr("false"), lo.ToPtr("FALSE"), lo.ToPtr("False"),
				lo.ToPtr("F"), lo.ToPtr("f"), lo.ToPtr("0"), lo.ToPtr(0),
			},
			want1: false,
			want2: true,
		},
		{
			name:  "nil",
			args:  []any{"foo", (*bool)(nil), (*string)(nil), nil},
			want1: nil,
			want2: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &propertyBool{}
			for i, v := range tt.args {
				got1, got2 := p.ToValue(v)
				assert.Equal(t, tt.want1, got1, "test %d", i)
				assert.Equal(t, tt.want2, got2, "test %d", i)
			}
		})
	}
}

func Test_propertyBool_ToInterface(t *testing.T) {
	v := true
	tt, ok := (&propertyBool{}).ToInterface(v)
	assert.Equal(t, v, tt)
	assert.Equal(t, true, ok)
}

func Test_propertyBool_IsEmpty(t *testing.T) {
	assert.False(t, (&propertyBool{}).IsEmpty(false))
	assert.False(t, (&propertyBool{}).IsEmpty(true))
}

func Test_propertyBool_Validate(t *testing.T) {
	assert.True(t, (&propertyBool{}).Validate(true))
	assert.False(t, (&propertyBool{}).Validate("a"))
}

func TestValue_ValueBool(t *testing.T) {
	var v *Value = nil
	res, ok := v.ValueBool()
	assert.Equal(t, false, res)
	assert.False(t, ok)

	v = &Value{
		t: TypeBool,
		v: nil,
		p: nil,
	}

	res, ok = v.ValueBool()
	assert.Equal(t, false, res)
	assert.False(t, ok)

	v = &Value{
		t: TypeBool,
		v: true,
		p: nil,
	}

	res, ok = v.ValueBool()
	assert.Equal(t, true, res)
	assert.True(t, ok)
}

func TestValue_ValuesBool(t *testing.T) {
	var v *Multiple = nil
	res, ok := v.ValuesBool()
	assert.Nil(t, res)
	assert.False(t, ok)

	v = &Multiple{
		t: TypeBool,
		v: nil,
	}
	res, ok = v.ValuesBool()
	assert.Equal(t, []Bool{}, res)
	assert.True(t, ok)

	v = &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, false)},
	}

	res, ok = v.ValuesBool()
	assert.Equal(t, []Bool{true, false}, res)
	assert.True(t, ok)

	v = &Multiple{
		t: TypeBool,
		v: []*Value{New(TypeBool, true), New(TypeBool, "test")},
	}

	res, ok = v.ValuesBool()
	assert.Nil(t, res)
	assert.False(t, ok)
}
