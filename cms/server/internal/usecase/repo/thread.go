package repo

import (
	"context"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var (
	ErrCommentNotFound error = rerror.NewE(i18n.T("comment not found"))
)

type Thread interface {
	Save(context.Context, *thread.Thread) error
	Filtered(filter WorkspaceFilter) Thread
	FindByID(ctx context.Context, id id.ThreadID) (*thread.Thread, error)
	FindByIDs(context.Context, id.ThreadIDList) ([]*thread.Thread, error)
}
