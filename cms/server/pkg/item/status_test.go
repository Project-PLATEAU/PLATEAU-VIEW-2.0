package item

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus_Wrap(t *testing.T) {

	tests := []struct {
		name  string
		s, s2 Status
		want  Status
	}{
		{
			name: "draft+draft=draft",
			s:    StatusDraft,
			s2:   StatusDraft,
			want: StatusDraft,
		},
		{
			name: "draft+review=review",
			s:    StatusDraft,
			s2:   StatusReview,
			want: StatusReview,
		},
		{
			name: "draft+public=public",
			s:    StatusDraft,
			s2:   StatusPublic,
			want: StatusPublic,
		},
		{
			name: "changed+public=publicdraft",
			s:    StatusChanged,
			s2:   StatusPublic,
			want: StatusPublicDraft,
		},
		{
			name: "draft+publicdraft=publicdraft",
			s:    StatusDraft,
			s2:   StatusPublicDraft,
			want: StatusPublicDraft,
		},
		{
			name: "review+public=publicreview",
			s:    StatusReview,
			s2:   StatusPublic,
			want: StatusPublicReview,
		},
		{
			name: "review+publicdraft=publicreview",
			s:    StatusReview,
			s2:   StatusPublicDraft,
			want: StatusPublicReview,
		},
		{
			name: "empty+public=public",
			s:    Status(0),
			s2:   StatusPublic,
			want: StatusPublic,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, tc.s.Wrap(tc.s2))
		})
	}
}
