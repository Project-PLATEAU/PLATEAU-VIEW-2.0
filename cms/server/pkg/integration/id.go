package integration

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
)

type ID = id.IntegrationID
type WebhookID = id.WebhookID
type UserID = id.UserID
type ModelID = id.ModelID

var NewID = id.NewIntegrationID
var NewWebhookID = id.NewWebhookID
var MustID = id.MustIntegrationID
var IDFrom = id.IntegrationIDFrom
var IDFromRef = id.IntegrationIDFromRef
var ErrInvalidID = id.ErrInvalidID
