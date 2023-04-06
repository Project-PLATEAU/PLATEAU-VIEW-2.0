package operator

import "github.com/reearth/reearth-cms/server/pkg/id"

type ID = id.EventID
type UserID = id.UserID
type IntegrationID = id.IntegrationID

var ErrInvalidID = id.ErrInvalidID
var NewIntegrationID = id.NewIntegrationID
