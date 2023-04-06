package project

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckAliasPattern(t *testing.T) {
	testCase := []struct {
		name, alias string
		expexted    bool
	}{
		{
			name:     "accepted regex",
			alias:    "xxxxx",
			expexted: true,
		},
		{
			name:     "refused regex",
			alias:    "xxx",
			expexted: false,
		},
	}

	for _, tt := range testCase {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expexted, CheckAliasPattern(tt.alias))
		})
	}
}

func TestProject_SetUpdatedAt(t *testing.T) {
	p := &Project{}
	p.SetUpdatedAt(time.Date(1900, 1, 1, 00, 00, 1, 1, time.UTC))
	assert.Equal(t, time.Date(1900, 1, 1, 00, 00, 1, 1, time.UTC), p.UpdatedAt())
}

func TestProject_SetImageURL(t *testing.T) {
	testCase := []struct {
		name        string
		image       *url.URL
		p           *Project
		expectedNil bool
	}{
		{
			name:        "nil image",
			image:       nil,
			p:           &Project{},
			expectedNil: true,
		},
		{
			name:        "set new image",
			image:       &url.URL{},
			p:           &Project{},
			expectedNil: false,
		},
	}

	for _, tt := range testCase {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.p.SetImageURL(tt.image)
			if tt.expectedNil {
				assert.Nil(t, tt.p.ImageURL())
			} else {
				assert.NotNil(t, tt.p.ImageURL())
			}
		})
	}
}

func TestProject_UpdateName(t *testing.T) {
	p := &Project{}
	p.UpdateName("foo")
	assert.Equal(t, "foo", p.Name())
}

func TestProject_UpdateDescription(t *testing.T) {
	p := &Project{}
	p.UpdateDescription("aaa")
	assert.Equal(t, "aaa", p.Description())
}

func TestProject_UpdateTeam(t *testing.T) {
	p := &Project{}
	p.UpdateTeam(NewWorkspaceID())
	assert.NotNil(t, p.Workspace())
}

func TestProject_UpdateAlias(t *testing.T) {
	tests := []struct {
		name, a  string
		expected string
		err      error
	}{
		{
			name:     "accepted alias",
			a:        "xxxxx",
			expected: "xxxxx",
			err:      nil,
		},
		{
			name:     "fail: invalid alias",
			a:        "xxx",
			expected: "",
			err:      ErrInvalidAlias,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := &Project{}
			err := p.UpdateAlias(tt.a)
			if tt.err == nil {
				assert.Equal(t, tt.expected, p.Alias())
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func TestProject_Clone(t *testing.T) {
	pub := &Publication{}
	p := New().NewID().Name("a").Publication(pub).MustBuild()

	got := p.Clone()
	assert.Equal(t, p, got)
	assert.NotSame(t, p, got)
	assert.NotSame(t, p, got.publication)
	assert.Nil(t, (*Project)(nil).Clone())
}
