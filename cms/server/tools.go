//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/reearth/reearthx/tools"
	_ "github.com/vektah/dataloaden"
)
