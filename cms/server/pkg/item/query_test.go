package item

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/stretchr/testify/assert"
)

func TestNewQuery(t *testing.T) {
	pid := id.NewProjectID()
	str := "foo"
	type args struct {
		project id.ProjectID
		schema  *id.SchemaID
		q       string
		ref     *version.Ref
	}
	tests := []struct {
		name string
		args args
		want *Query
	}{
		{
			name: "",
			args: args{
				project: pid,
				q:       str,
				ref:     version.Public.Ref(),
			},
			want: &Query{
				project: pid,
				q:       "foo",
				ref:     version.Public.Ref(),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := NewQuery(tc.args.project, tc.args.schema, tc.args.q, tc.args.ref)
			assert.Equal(tt, tc.want, got)
		})
	}
}

func TestQuery_Project(t *testing.T) {
	pid := id.NewProjectID()
	q := &Query{
		project: pid,
	}
	assert.Equal(t, pid, q.Project())
}

func TestQuery_Q(t *testing.T) {
	str := "foo"
	q := &Query{
		q: str,
	}
	assert.Equal(t, str, q.Q())
}

func TestQuery_Ref(t *testing.T) {
	q := &Query{
		ref: version.Public.Ref(),
	}
	assert.Equal(t, version.Public.Ref(), q.Ref())
}
