package asset

import (
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var (
	ErrNoProjectID = rerror.NewE(i18n.T("projectID is required"))
	ErrZeroSize    = rerror.NewE(i18n.T("file size cannot be zero"))
	ErrNoUser      = rerror.NewE(i18n.T("createdBy is required"))
	ErrNoThread    = rerror.NewE(i18n.T("thread is required"))
	ErrNoUUID      = rerror.NewE(i18n.T("uuid is required"))
)
