package user

import "github.com/reearth/reearth-cms/server/pkg/id"

type ID = id.UserID
type WorkspaceID = id.WorkspaceID
type IntegrationID = id.IntegrationID

var NewID = id.NewUserID
var NewWorkspaceID = id.NewWorkspaceID

var MustID = id.MustUserID
var MustWorkspaceID = id.MustWorkspaceID

var IDFrom = id.UserIDFrom
var WorkspaceIDFrom = id.WorkspaceIDFrom

var IDFromRef = id.UserIDFromRef
var WorkspaceIDFromRef = id.WorkspaceIDFromRef

var ErrInvalidID = id.ErrInvalidID

type WorkspaceIDList = id.WorkspaceIDList
type IntegrationIDList = id.IntegrationIDList
