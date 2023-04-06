package model

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/samber/lo"
)

type ID = id.ModelID
type ProjectID = id.ProjectID
type SchemaID = id.SchemaID

var NewID = id.NewModelID
var MustID = id.MustModelID
var IDFrom = id.ModelIDFrom
var IDFromRef = id.ModelIDFromRef
var ErrInvalidID = id.ErrInvalidID

type IDOrKey string

func (i IDOrKey) ID() *ID {
	return IDFromRef(lo.ToPtr(string(i)))
}

func (i IDOrKey) Key() *string {
	if i.ID() == nil {
		return lo.ToPtr(string(i))
	}
	return nil
}
