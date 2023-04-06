package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType_Default(t *testing.T) {
	tests := []struct {
		name string
		tr   Type
		want bool
	}{
		{
			name: "default",
			tr:   TypeText,
			want: true,
		},
		{
			name: "custom",
			tr:   Type("foo"),
			want: false,
		},
		{
			name: "unknown",
			tr:   TypeUnknown,
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.tr.Default())
		})
	}
}

func TestType_None(t *testing.T) {
	tests := []struct {
		name string
		tr   Type
		want *Optional
	}{
		{
			name: "default",
			tr:   TypeText,
			want: &Optional{t: TypeText},
		},
		{
			name: "custom",
			tr:   Type("foo"),
			want: &Optional{t: Type("foo")},
		},
		{
			name: "unknown",
			tr:   TypeUnknown,
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.tr.None())
		})
	}
}

func TestType_Value(t *testing.T) {
	type args struct {
		i any
	}

	tests := []struct {
		name string
		tr   Type
		args args
		want *Value
	}{
		{
			name: "default type",
			tr:   TypeText,
			args: args{
				i: "hoge",
			},
			want: &Value{t: TypeText, v: "hoge"},
		},
		{
			name: "nil",
			tr:   TypeText,
			args: args{},
			want: nil,
		},
		{
			name: "unknown type",
			tr:   TypeUnknown,
			args: args{
				i: "hoge",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.tr.Value(tt.args.i))
		})
	}
}

func TestType_ValueFrom(t *testing.T) {
	tpm := TypeRegistry{
		Type("foo"): &tpmock{},
	}

	type args struct {
		i any
		p TypeRegistry
	}

	tests := []struct {
		name string
		tr   Type
		args args
		want *Value
	}{
		{
			name: "default type",
			tr:   TypeText,
			args: args{
				i: "hoge",
			},
			want: &Value{t: TypeText, v: "hoge"},
		},
		{
			name: "custom type",
			tr:   Type("foo"),
			args: args{
				i: "hoge",
				p: tpm,
			},
			want: &Value{p: tpm, t: Type("foo"), v: "hogea"},
		},
		{
			name: "nil",
			tr:   TypeText,
			args: args{},
			want: nil,
		},
		{
			name: "unknown type",
			tr:   TypeUnknown,
			args: args{
				i: "hoge",
			},
			want: nil,
		},
		{
			name: "unknown type + custom type",
			tr:   Type("bar"),
			args: args{
				i: "hoge",
				p: tpm,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.tr.ValueFrom(tt.args.i, tt.args.p))
		})
	}
}
