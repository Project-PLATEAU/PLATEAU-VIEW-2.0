package event

import "github.com/reearth/reearth-cms/server/pkg/id"

type ID = id.EventID
type UserID = id.UserID
type IntegrationID = id.IntegrationID

var NewID = id.NewEventID
var MustID = id.MustEventID
var IDFrom = id.EventIDFrom
var IDFromRef = id.EventIDFromRef
var ErrInvalidID = id.ErrInvalidID
