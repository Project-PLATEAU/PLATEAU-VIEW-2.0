package project

import (
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var tb = New()
	assert.NotNil(t, tb)
}

func TestBuilder_ID(t *testing.T) {
	var tb = New()
	res := tb.ID(NewID()).MustBuild()
	assert.NotNil(t, res.ID())
}

func TestBuilder_Name(t *testing.T) {
	var tb = New().NewID()
	res := tb.Name("foo").MustBuild()
	assert.Equal(t, "foo", res.Name())
}

func TestBuilder_NewID(t *testing.T) {
	var tb = New()
	res := tb.NewID().MustBuild()
	assert.NotNil(t, res.ID())
}

func TestBuilder_Alias(t *testing.T) {
	var tb = New().NewID()
	res := tb.Alias("xxxxx").MustBuild()
	assert.Equal(t, "xxxxx", res.Alias())
}

func TestBuilder_Description(t *testing.T) {
	var tb = New().NewID()
	res := tb.Description("desc").MustBuild()
	assert.Equal(t, "desc", res.Description())
}

func TestBuilder_ImageURL(t *testing.T) {
	tests := []struct {
		name        string
		image       *url.URL
		expectedNil bool
	}{
		{
			name:        "image not nil",
			image:       &url.URL{},
			expectedNil: false,
		},
		{
			name:        "image is nil",
			image:       nil,
			expectedNil: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tb := New().NewID()
			res := tb.ImageURL(tt.image).MustBuild()
			if res.imageURL == nil {
				assert.True(t, tt.expectedNil)
			} else {
				assert.False(t, tt.expectedNil)
			}
		})
	}
}

func TestBuilder_Team(t *testing.T) {
	var tb = New().NewID()
	res := tb.Workspace(NewWorkspaceID()).MustBuild()
	assert.NotNil(t, res.Workspace())
}

func TestBuilder_UpdatedAt(t *testing.T) {
	var tb = New().NewID()
	d := time.Date(1900, 1, 1, 00, 00, 0, 1, time.UTC)
	res := tb.UpdatedAt(d).MustBuild()
	assert.True(t, reflect.DeepEqual(res.UpdatedAt(), d))
}

func TestBuilder_Publication(t *testing.T) {
	var tb = New().NewID()
	p := &Publication{}
	res := tb.Publication(p)
	assert.Equal(t, &Builder{
		p: &Project{id: tb.p.id, publication: p},
	}, res)
}

func TestBuilder_Build(t *testing.T) {
	d := time.Date(1900, 1, 1, 00, 00, 0, 1, time.UTC)
	i, _ := url.Parse("ttt://xxx.aa/")
	pid := NewID()
	tid := NewWorkspaceID()

	type args struct {
		name, description string
		alias             string
		id                ID
		updatedAt         time.Time
		imageURL          *url.URL
		team              WorkspaceID
	}

	tests := []struct {
		name     string
		args     args
		expected *Project
		err      error
	}{
		{
			name: "build normal project",
			args: args{
				name:        "xxx.aaa",
				description: "ddd",
				alias:       "aaaaa",
				id:          pid,
				updatedAt:   d,
				imageURL:    i,
				team:        tid,
			},
			expected: &Project{
				id:          pid,
				description: "ddd",
				name:        "xxx.aaa",
				alias:       "aaaaa",
				updatedAt:   d,
				imageURL:    i,
				workspaceID: tid,
			},
		},
		{
			name: "zero updated at",
			args: args{
				id: pid,
			},
			expected: &Project{
				id:        pid,
				updatedAt: pid.Timestamp(),
			},
		},
		{
			name: "failed invalid id",
			err:  ErrInvalidID,
		},
		{
			name: "failed invalid alias",
			args: args{
				id:    NewID(),
				alias: "xxx.aaa",
			},
			expected: nil,
			err:      ErrInvalidAlias,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p, err := New().
				ID(tt.args.id).
				UpdatedAt(tt.args.updatedAt).
				Workspace(tt.args.team).
				ImageURL(tt.args.imageURL).
				Name(tt.args.name).
				Alias(tt.args.alias).
				UpdatedAt(tt.args.updatedAt).
				Description(tt.args.description).
				Build()

			if tt.err == nil {
				assert.Equal(t, tt.expected, p)
			} else {
				assert.Equal(t, tt.err, err)
			}
		})
	}
}

func TestBuilder_MustBuild(t *testing.T) {
	d := time.Date(1900, 1, 1, 00, 00, 0, 1, time.UTC)
	i, _ := url.Parse("ttt://xxx.aa/")
	pid := NewID()
	tid := NewWorkspaceID()

	type args struct {
		name, description string
		alias             string
		id                ID
		updatedAt         time.Time
		imageURL          *url.URL
		team              WorkspaceID
	}

	tests := []struct {
		name     string
		args     args
		expected *Project
		err      error
	}{
		{
			name: "build normal project",
			args: args{
				name:        "xxx.aaa",
				description: "ddd",
				alias:       "aaaaa",
				id:          pid,
				updatedAt:   d,
				imageURL:    i,
				team:        tid,
			},
			expected: &Project{
				id:          pid,
				description: "ddd",
				name:        "xxx.aaa",
				alias:       "aaaaa",
				updatedAt:   d,
				imageURL:    i,
				workspaceID: tid,
			},
		},
		{
			name: "zero updated at",
			args: args{
				id: pid,
			},
			expected: &Project{
				id:        pid,
				updatedAt: pid.Timestamp(),
			},
		},
		{
			name: "failed invalid id",
			err:  ErrInvalidID,
		},
		{
			name: "failed invalid alias",
			args: args{
				id:    NewID(),
				alias: "xxx.aaa",
			},
			err: ErrInvalidAlias,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			build := func() *Project {
				t.Helper()
				return New().
					ID(tt.args.id).
					UpdatedAt(tt.args.updatedAt).
					Workspace(tt.args.team).
					ImageURL(tt.args.imageURL).
					Name(tt.args.name).
					Alias(tt.args.alias).
					UpdatedAt(tt.args.updatedAt).
					Description(tt.args.description).
					MustBuild()
			}

			if tt.err != nil {
				assert.PanicsWithValue(t, tt.err, func() { _ = build() })
			} else {
				assert.Equal(t, tt.expected, build())
			}
		})
	}
}
