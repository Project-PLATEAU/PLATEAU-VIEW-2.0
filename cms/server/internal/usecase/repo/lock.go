package repo

import (
	"context"

	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var (
	ErrFailedToLock  = rerror.NewE(i18n.T("failed to lock"))
	ErrAlreadyLocked = rerror.NewE(i18n.T("already locked"))
	ErrNotLocked     = rerror.NewE(i18n.T("not locked"))
)

type Lock interface {
	Lock(context.Context, string) error
	Unlock(context.Context, string) error
}
