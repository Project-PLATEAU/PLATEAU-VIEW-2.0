package value

import (
	"net/url"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func Test_propertyURL_ToValue(t *testing.T) {
	u, _ := url.Parse("https://example.com")

	tests := []struct {
		name  string
		args  []any
		want1 any
		want2 bool
	}{
		{
			name:  "string",
			args:  []any{"https://example.com", lo.ToPtr("https://example.com")},
			want1: u,
			want2: true,
		},
		{
			name:  "string empty",
			args:  []any{""},
			want1: nil,
			want2: false,
		},
		{
			name:  "url",
			args:  []any{u, *u},
			want1: u,
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
			p := &propertyURL{}
			for i, v := range tt.args {
				got1, got2 := p.ToValue(v)
				assert.Equal(t, tt.want1, got1, "test %d", i)
				assert.Equal(t, tt.want2, got2, "test %d", i)
			}
		})
	}
}

func Test_propertyURL_ToInterface(t *testing.T) {
	v := lo.Must(url.Parse("https://example.com"))
	tt, ok := (&propertyURL{}).ToInterface(v)
	assert.Equal(t, v.String(), tt)
	assert.Equal(t, true, ok)
}

func Test_propertyURL_IsEmpty(t *testing.T) {
	assert.True(t, (&propertyURL{}).IsEmpty(&url.URL{}))
	assert.False(t, (&propertyURL{}).IsEmpty(&url.URL{Path: "a"}))
}

func Test_propertyURL_Validate(t *testing.T) {
	v := lo.Must(url.Parse("https://example.com"))
	assert.True(t, (&propertyURL{}).Validate(v))
}
