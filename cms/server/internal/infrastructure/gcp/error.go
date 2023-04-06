package gcp

import (
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var (
	ErrMissignConfig = rerror.NewE(i18n.T("missing required config"))
)
