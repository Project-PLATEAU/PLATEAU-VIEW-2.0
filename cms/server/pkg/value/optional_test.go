package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOptional(t *testing.T) {
	assert.Nil(t, NewOptional(TypeAsset, TypeBool.Value(true)))
	assert.Equal(t, &Optional{t: TypeBool}, NewOptional(TypeBool, nil))
	assert.Equal(t, &Optional{t: TypeBool, v: TypeBool.Value(true)}, NewOptional(TypeBool, TypeBool.Value(true)))
}

func TestOptionalFrom(t *testing.T) {
	type args struct {
		v *Value
	}

	tests := []struct {
		name string
		args args
		want *Optional
	}{
		{
			name: "default type",
			args: args{
				v: TypeText.ValueFrom("foo", nil),
			},
			want: &Optional{t: TypeText, v: TypeText.ValueFrom("foo", nil)},
		},
		{
			name: "custom type",
			args: args{
				v: &Value{t: Type("foo"), v: 1},
			},
			want: &Optional{t: Type("foo"), v: &Value{t: Type("foo"), v: 1}},
		},
		{
			name: "invalid value",
			args: args{
				v: &Value{v: "string"},
			},
			want: nil,
		},
		{
			name: "nil value",
			args: args{},
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, OptionalFrom(tt.args.v))
		})
	}
}

func TestOptional_Type(t *testing.T) {
	tests := []struct {
		name  string
		value *Optional
		want  Type
	}{
		{
			name:  "ok",
			value: &Optional{t: Type("foo")},
			want:  Type("foo"),
		},
		{
			name:  "empty",
			value: &Optional{},
			want:  TypeUnknown,
		},
		{
			name:  "nil",
			value: nil,
			want:  TypeUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.value.Type())
		})
	}
}

func TestOptional_Value(t *testing.T) {
	tests := []struct {
		name  string
		value *Optional
		want  *Value
	}{
		{
			name:  "ok",
			value: &Optional{t: TypeText, v: &Value{t: TypeText, v: "foobar"}},
			want:  &Value{t: TypeText, v: "foobar"},
		},
		{
			name:  "empty",
			value: &Optional{},
			want:  nil,
		},
		{
			name:  "nil",
			value: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := tt.value.Value()
			assert.Equal(t, tt.want, res)
			if res != nil {
				assert.NotSame(t, tt.want, res)
			}
		})
	}
}

func TestOptional_TypeAndValue(t *testing.T) {
	tests := []struct {
		name  string
		value *Optional
		wantt Type
		wantv *Value
	}{
		{
			name:  "ok",
			value: &Optional{t: TypeText, v: &Value{t: TypeText, v: "foobar"}},
			wantt: TypeText,
			wantv: &Value{t: TypeText, v: "foobar"},
		},
		{
			name:  "empty",
			value: &Optional{},
			wantt: TypeUnknown,
			wantv: nil,
		},
		{
			name:  "nil",
			value: nil,
			wantt: TypeUnknown,
			wantv: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ty, tv := tt.value.TypeAndValue()
			assert.Equal(t, tt.wantt, ty)
			assert.Equal(t, tt.wantv, tv)
			if tv != nil {
				assert.NotSame(t, tt.wantv, tv)
			}
		})
	}
}

func TestOptional_SetValue(t *testing.T) {
	type args struct {
		v *Value
	}

	tests := []struct {
		name    string
		value   *Optional
		args    args
		invalid bool
	}{
		{
			name: "set",
			value: &Optional{
				t: TypeText,
				v: &Value{t: TypeText, v: "foobar"},
			},
			args: args{v: &Value{t: TypeText, v: "bar"}},
		},
		{
			name: "set to nil",
			value: &Optional{
				t: TypeText,
			},
			args: args{v: &Value{t: TypeText, v: "bar"}},
		},
		{
			name: "invalid value",
			value: &Optional{
				t: TypeNumber,
				v: &Value{t: TypeNumber, v: 1},
			},
			args:    args{v: &Value{t: TypeText, v: "bar"}},
			invalid: true,
		},
		{
			name: "nil value",
			value: &Optional{
				t: TypeNumber,
				v: &Value{t: TypeNumber, v: 1},
			},
		},
		{
			name:    "empty",
			value:   &Optional{},
			args:    args{v: &Value{t: TypeText, v: "bar"}},
			invalid: true,
		},
		{
			name: "nil",
			args: args{v: &Value{t: TypeText, v: "bar"}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var v *Value
			if tt.value != nil {
				v = tt.value.v
			}

			tt.value.SetValue(tt.args.v)

			if tt.value != nil {
				if tt.invalid {
					assert.Same(t, v, tt.value.v)
				} else {
					assert.Equal(t, tt.args.v, tt.value.v)
					if tt.args.v != nil {
						assert.NotSame(t, tt.args.v, tt.value.v)
					}
				}
			}
		})
	}
}

func TestOptional_Clone(t *testing.T) {
	tests := []struct {
		name   string
		target *Optional
	}{
		{
			name:   "ok",
			target: &Optional{t: TypeText, v: TypeText.ValueFrom("foo", nil)},
		},
		{
			name:   "empty",
			target: &Optional{},
		},
		{
			name:   "nil",
			target: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := tt.target.Clone()
			assert.Equal(t, tt.target, res)
			if tt.target != nil {
				assert.NotSame(t, tt.target, res)
			}
		})
	}
}

func TestOptional_Cast(t *testing.T) {
	type args struct {
		t Type
	}

	tests := []struct {
		name   string
		target *Optional
		args   args
		want   *Optional
	}{
		{
			name:   "diff type",
			target: &Optional{t: TypeNumber, v: TypeNumber.ValueFrom(1.1, nil)},
			args:   args{t: TypeText},
			want:   &Optional{t: TypeText, v: TypeText.ValueFrom("1.1", nil)},
		},
		{
			name:   "same type",
			target: &Optional{t: TypeNumber, v: TypeNumber.ValueFrom(1.1, nil)},
			args:   args{t: TypeNumber},
			want:   &Optional{t: TypeNumber, v: TypeNumber.ValueFrom(1.1, nil)},
		},
		{
			name:   "nil value",
			target: &Optional{t: TypeNumber},
			args:   args{t: TypeText},
			want:   &Optional{t: TypeText},
		},
		{
			name:   "failed to cast",
			target: &Optional{t: TypeBool, v: TypeBool.ValueFrom(true, nil)},
			args:   args{t: TypeDateTime},
			want:   &Optional{t: TypeDateTime},
		},
		{
			name:   "empty",
			target: &Optional{},
			args:   args{t: TypeText},
			want:   nil,
		},
		{
			name:   "nil",
			target: nil,
			args:   args{t: TypeText},
			want:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.target.Cast(tt.args.t))
		})
	}
}

func TestOptional_IsSome_IsNone_IsEmpty(t *testing.T) {
	v := (*Optional)(nil)
	assert.False(t, v.IsSome())
	assert.True(t, v.IsNone())
	assert.True(t, v.IsEmpty())

	v = &Optional{t: TypeText}
	assert.False(t, v.IsSome())
	assert.True(t, v.IsNone())
	assert.True(t, v.IsEmpty())

	v = &Optional{t: TypeText, v: &Value{t: TypeText, v: ""}}
	assert.True(t, v.IsSome())
	assert.False(t, v.IsNone())
	assert.True(t, v.IsEmpty())

	v = &Optional{t: TypeText, v: &Value{t: TypeText, v: "a"}}
	assert.True(t, v.IsSome())
	assert.False(t, v.IsNone())
	assert.False(t, v.IsEmpty())
}
