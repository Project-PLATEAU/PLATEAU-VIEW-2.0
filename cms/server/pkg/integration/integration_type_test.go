package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeFrom(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Type
	}{
		{
			name:  "public",
			input: "public",
			want:  TypePublic,
		},
		{
			name:  "private",
			input: "private",
			want:  TypePrivate,
		},
		{
			name:  "other",
			input: "xyz",
			want:  TypePrivate,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, TypeFrom(tt.input))
		})
	}
}
