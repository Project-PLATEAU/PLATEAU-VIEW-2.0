package fs

import (
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

const (
	fileSizeLimit int64 = 10 * 1024 * 1024 * 1024 // 10GB
	assetDir            = "assets"
	defaultBase         = "http://localhost:8080/assets"
)

var (
	ErrInvalidBaseURL = rerror.NewE(i18n.T("invalid base URL"))
)
