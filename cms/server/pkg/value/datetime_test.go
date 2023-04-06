package value

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func Test_propertyDateTime_ToValue(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tests := []struct {
		name  string
		args  []any
		want1 any
		want2 bool
	}{
		{
			name: "time",
			args: []any{
				now, now.Format(time.RFC3339), now.Format(time.RFC3339Nano),
			},
			want1: now,
			want2: true,
		},
		{
			name:  "integer",
			args:  []any{now.Unix(), float64(now.Unix()), json.Number(fmt.Sprintf("%d", now.Unix()))},
			want1: now,
			want2: true,
		},
		{
			name:  "nil",
			args:  []any{"foo", (*float64)(nil), (*string)(nil), (*int)(nil), (*json.Number)(nil), nil},
			want1: nil,
			want2: false,
		},
		{
			name:  "bool",
			args:  []any{true, false, lo.ToPtr(true), lo.ToPtr(false)},
			want1: nil,
			want2: false,
		},
		{
			name: "pointers",
			args: []any{
				&now, lo.ToPtr(now.Format(time.RFC3339)), lo.ToPtr(now.Format(time.RFC3339Nano)),
				lo.ToPtr(now.Unix()), lo.ToPtr(float64(now.Unix())), lo.ToPtr(json.Number(fmt.Sprintf("%d", now.Unix()))),
			},
			want1: now,
			want2: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &propertyDateTime{}
			for i, v := range tt.args {
				got1, got2 := p.ToValue(v)
				if tt.want1 != nil {
					assert.Equal(t, tt.want1.(time.Time).Unix(), got1.(time.Time).Unix(), "test %d", i)
				} else {
					assert.Nil(t, got1, "test %d", i)
				}
				assert.Equal(t, tt.want2, got2, "test %d", i)
			}
		})
	}
}

func Test_propertyDateTime_ToInterface(t *testing.T) {
	v := time.Now()
	tt, ok := (&propertyDateTime{}).ToInterface(v)
	assert.Equal(t, v.Format(time.RFC3339), tt)
	assert.Equal(t, true, ok)
}

func Test_propertyDateTime_IsEmpty(t *testing.T) {
	assert.True(t, (&propertyDateTime{}).IsEmpty(time.Time{}))
	assert.False(t, (&propertyDateTime{}).IsEmpty(time.Now()))
}

func Test_propertyDateTime_Validate(t *testing.T) {
	assert.True(t, (&propertyDateTime{}).Validate(time.Now()))
	assert.False(t, (&propertyDateTime{}).Validate("a"))
}

func Test_propertyDateTime_Equal(t *testing.T) {
	now := time.Now()
	p := &propertyDateTime{}
	assert.True(t, (&propertyDateTime{}).Equal(now, lo.Must(p.ToValue(&now))))
	assert.True(t, (&propertyDateTime{}).Equal(now, lo.Must(p.ToValue(now))))
	assert.False(t, (&propertyDateTime{}).Equal(now, now.Add(2*time.Millisecond)))
	assert.False(t, (&propertyDateTime{}).Equal(now, now.Add(2*time.Millisecond)))
}

func TestValue_ValueDateTime(t *testing.T) {
	var v *Value = nil
	res, ok := v.ValueDateTime()
	assert.Equal(t, DateTime{}, res)
	assert.False(t, ok)

	v = &Value{
		t: TypeDateTime,
		v: nil,
		p: nil,
	}

	res, ok = v.ValueDateTime()
	assert.Equal(t, DateTime{}, res)
	assert.False(t, ok)

	now := time.Now()
	v = &Value{
		t: TypeDateTime,
		v: now,
		p: nil,
	}

	res, ok = v.ValueDateTime()
	assert.Equal(t, now, res)
	assert.True(t, ok)
}

func TestValue_ValuesDateTime(t *testing.T) {
	var v *Multiple = nil
	res, ok := v.ValuesDateTime()
	assert.Nil(t, res)
	assert.False(t, ok)

	v = &Multiple{
		t: TypeDateTime,
		v: nil,
	}
	res, ok = v.ValuesDateTime()
	assert.Equal(t, []DateTime{}, res)
	assert.True(t, ok)

	now1 := time.Now()
	now2 := time.Now()
	v = &Multiple{
		t: TypeDateTime,
		v: []*Value{New(TypeDateTime, now1), New(TypeDateTime, now2)},
	}

	res, ok = v.ValuesDateTime()
	assert.Equal(t, []DateTime{now1, now2}, res)
	assert.True(t, ok)

	v = &Multiple{
		t: TypeDateTime,
		v: []*Value{New(TypeDateTime, now1), New(TypeInteger, 1)},
	}

	res, ok = v.ValuesDateTime()
	assert.Nil(t, res)
	assert.False(t, ok)
}
