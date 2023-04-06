package value

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func Test_propertyInteger_ToValue(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name  string
		args  []any
		want1 any
		want2 bool
	}{
		{
			name: "zero",
			args: []any{
				0, 0.0, "0", json.Number("0"), json.Number("-0"),
				lo.ToPtr(0), lo.ToPtr(0.0), lo.ToPtr("0"), lo.ToPtr(json.Number("0")), lo.ToPtr(json.Number("-0")),
			},
			want1: int64(0),
			want2: true,
		},
		{
			name: "positive",
			args: []any{
				1.12, "1", "1.12", json.Number("1.12"),
				lo.ToPtr(1.12), lo.ToPtr("1.12"), lo.ToPtr(json.Number("1.12")),
			},
			want1: int64(1),
			want2: true,
		},
		{
			name: "negative",
			args: []any{
				-2.11, "-2", "-2.11", json.Number("-2.11"),
				lo.ToPtr(-2.11), lo.ToPtr("-2.11"), lo.ToPtr(json.Number("-2.11")),
			},
			want1: int64(-2),
			want2: true,
		},
		{
			name: "inf",
			args: []any{
				math.Inf(0), json.Number("Infinity"),
			},
			want1: int64(math.Inf(0)),
			want2: true,
		},
		{
			name: "negative inf",
			args: []any{
				math.Inf(-1), json.Number("-Infinity"),
			},
			want1: int64(math.Inf(-1)),
			want2: true,
		},
		{
			name:  "time",
			args:  []any{now, &now},
			want1: now.Unix(),
			want2: true,
		},
		{
			name:  "bool true",
			args:  []any{true, lo.ToPtr(true)},
			want1: int64(1),
			want2: true,
		},
		{
			name:  "bool false",
			args:  []any{false, lo.ToPtr(false)},
			want1: int64(0),
			want2: true,
		},
		{
			name:  "nil",
			args:  []any{"foo", (*float64)(nil), (*string)(nil), (*int)(nil), (*json.Number)(nil), nil, math.NaN()},
			want1: nil,
			want2: false,
		},
		{
			name: "types",
			args: []any{
				"123", lo.ToPtr("123"),
				123, lo.ToPtr(123),
				int(123), lo.ToPtr(int(123)),
				int8(123), lo.ToPtr(int8(123)),
				int16(123), lo.ToPtr(int16(123)),
				int32(123), lo.ToPtr(int32(123)),
				int64(123), lo.ToPtr(int64(123)),
				uint(123), lo.ToPtr(uint(123)),
				uint8(123), lo.ToPtr(uint8(123)),
				uint16(123), lo.ToPtr(uint16(123)),
				uint32(123), lo.ToPtr(uint32(123)),
				uint64(123), lo.ToPtr(uint64(123)),
				uintptr(123), lo.ToPtr(uintptr(123)),
				float32(123), lo.ToPtr(float32(123)),
				float64(123), lo.ToPtr(float64(123)),
			},
			want1: int64(123),
			want2: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &propertyInteger{}
			for i, v := range tt.args {
				got1, got2 := p.ToValue(v)
				assert.Equal(t, tt.want1, got1, "test %d", i)
				assert.Equal(t, tt.want2, got2, "test %d", i)
			}
		})
	}
}

func Test_propertyInteger_ToInterface(t *testing.T) {
	v := int64(1)
	tt, ok := (&propertyInteger{}).ToInterface(v)
	assert.Equal(t, v, tt)
	assert.Equal(t, true, ok)
}

func Test_propertyInteger_IsEmpty(t *testing.T) {
	assert.False(t, (&propertyInteger{}).IsEmpty(0))
}

func Test_propertyInteger_Validate(t *testing.T) {
	assert.True(t, (&propertyInteger{}).Validate(int64(1)))
	assert.False(t, (&propertyInteger{}).Validate("a"))
}
