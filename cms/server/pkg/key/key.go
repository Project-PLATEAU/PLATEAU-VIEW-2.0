package key

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
)

// TODO: completely delete this "key" package
type Key = id.Key

var New = id.NewKey
var Random = id.RandomKey
