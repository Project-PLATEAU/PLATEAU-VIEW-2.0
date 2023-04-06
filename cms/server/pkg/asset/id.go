package asset

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
)

type ID = id.AssetID
type IDList = id.AssetIDList
type ProjectID = id.ProjectID
type UserID = id.UserID
type ThreadID = id.ThreadID
type IntegrationID = id.IntegrationID

var NewID = id.NewAssetID
var NewProjectID = id.NewProjectID
var NewUserID = id.NewUserID
var NewThreadID = id.NewThreadID
var NewIntegrationID = id.NewIntegrationID

var MustID = id.MustAssetID
var MustProjectID = id.MustProjectID
var MustUserID = id.MustUserID
var MustThreadID = id.MustThreadID

var IDFrom = id.AssetIDFrom
var ProjectIDFrom = id.ProjectIDFrom
var UserIDFrom = id.UserIDFrom
var ThreadIDFrom = id.ThreadIDFrom

var IDFromRef = id.AssetIDFromRef
var ProjectIDFromRef = id.ProjectIDFromRef
var UserIDFromRef = id.UserIDFromRef
var ThreadIDFromRef = id.ThreadIDFromRef

var ErrInvalidID = id.ErrInvalidID
