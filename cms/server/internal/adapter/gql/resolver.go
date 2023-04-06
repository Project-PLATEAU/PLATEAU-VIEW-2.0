//go:generate go run github.com/99designs/gqlgen generate

package gql

import (
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var ErrNotImplemented = rerror.NewE(i18n.T("not implemented yet"))
var ErrUnauthorized = rerror.NewE(i18n.T("unauthorized"))

type Resolver struct {
}

func NewResolver() ResolverRoot {
	return &Resolver{}
}
