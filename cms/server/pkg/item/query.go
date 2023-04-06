package item

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/util"
)

type Query struct {
	project id.ProjectID
	schema  *id.SchemaID
	q       string
	ref     *version.Ref
}

func NewQuery(project id.ProjectID, schema *id.SchemaID, q string, ref *version.Ref) *Query {
	return &Query{
		project: project,
		schema:  schema,
		q:       q,
		ref:     ref,
	}
}

// Q returns keywords for search
func (q *Query) Q() string {
	return q.q
}

func (q *Query) Project() id.ProjectID {
	return q.project
}

func (q *Query) Schema() *id.SchemaID {
	return q.schema
}

func (q *Query) Ref() *version.Ref {
	return util.CloneRef(q.ref)
}
