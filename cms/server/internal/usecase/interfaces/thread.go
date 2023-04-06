package interfaces

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var (
	ErrCommentAlreadyExist = rerror.NewE(i18n.T("Comment already exist in this thread"))
	ErrCommentDoesNotExist = rerror.NewE(i18n.T("Comment does not exist in this thread"))
)

type Thread interface {
	FindByID(context.Context, id.ThreadID, *usecase.Operator) (*thread.Thread, error)
	FindByIDs(context.Context, []id.ThreadID, *usecase.Operator) (thread.List, error)
	CreateThread(context.Context, id.WorkspaceID, *usecase.Operator) (*thread.Thread, error)
	AddComment(context.Context, id.ThreadID, string, *usecase.Operator) (*thread.Thread, *thread.Comment, error)
	UpdateComment(context.Context, id.ThreadID, id.CommentID, string, *usecase.Operator) (*thread.Thread, *thread.Comment, error)
	DeleteComment(context.Context, id.ThreadID, id.CommentID, *usecase.Operator) (*thread.Thread, error)
}
