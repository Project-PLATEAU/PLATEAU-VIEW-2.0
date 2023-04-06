package schema

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
)

type FieldID = id.FieldID
type WorkspaceID = id.WorkspaceID

var NewFieldID = id.NewFieldID
var MustFieldID = id.MustFieldID
var FieldIDFrom = id.FieldIDFrom
var FieldIDFromRef = id.FieldIDFromRef
var ErrInvalidFieldID = id.ErrInvalidID

type ID = id.SchemaID
type ProjectID = id.ProjectID

var NewID = id.NewSchemaID
var MustID = id.MustSchemaID
var IDFrom = id.SchemaIDFrom
var IDFromRef = id.SchemaIDFromRef
var ErrInvalidID = id.ErrInvalidID
