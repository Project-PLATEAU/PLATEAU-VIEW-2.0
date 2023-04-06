package thread

import (
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var (
	ErrNoWorkspaceID       = rerror.NewE(i18n.T("workspace id is required"))
	ErrCommentAlreadyExist = rerror.NewE(i18n.T("comment already exist in this thread"))
	ErrCommentDoesNotExist = rerror.NewE(i18n.T("comment does not exist in this thread"))
)
