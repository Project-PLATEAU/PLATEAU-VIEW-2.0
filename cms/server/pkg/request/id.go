package request

import "github.com/reearth/reearth-cms/server/pkg/id"

type ID = id.RequestID
type WorkspaceID = id.WorkspaceID
type ProjectID = id.ProjectID
type ItemID = id.ItemID
type UserID = id.UserID
type UserIDList = id.UserIDList
type ThreadID = id.ThreadID

var NewID = id.NewRequestID
var NewWorkspaceID = id.NewWorkspaceID
var NewProjectID = id.NewProjectID
var NewThreadID = id.NewThreadID
var NewUserID = id.NewUserID
var NewItemID = id.NewItemID
var MustID = id.MustRequestID
var IDFrom = id.RequestIDFrom
var IDFromRef = id.RequestIDFromRef

var ErrInvalidID = id.ErrInvalidID
