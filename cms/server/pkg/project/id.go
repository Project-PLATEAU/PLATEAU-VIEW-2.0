package project

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/samber/lo"
)

type ID = id.ProjectID
type WorkspaceID = id.WorkspaceID

type IDList = id.ProjectIDList

var NewID = id.NewProjectID
var NewWorkspaceID = id.NewWorkspaceID

var MustID = id.MustProjectID
var MustWorkspaceID = id.MustWorkspaceID

var IDFrom = id.ProjectIDFrom
var WorkspaceIDFrom = id.WorkspaceIDFrom

var IDFromRef = id.ProjectIDFromRef
var WorkspaceIDFromRef = id.WorkspaceIDFromRef

var ErrInvalidID = id.ErrInvalidID

type IDOrAlias string

func (i IDOrAlias) ID() *ID {
	return IDFromRef(lo.ToPtr(string(i)))
}

func (i IDOrAlias) Alias() *string {
	if string(i) != "" && i.ID() == nil {
		return lo.ToPtr(string(i))
	}
	return nil
}
