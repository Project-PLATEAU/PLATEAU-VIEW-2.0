package repo

import (
	"context"

	"github.com/reearth/reearth/server/pkg/id"
	"github.com/reearth/reearth/server/pkg/user"
)

type User interface {
	FindByIDs(context.Context, id.UserIDList) ([]*user.User, error)
	FindByID(context.Context, id.UserID) (*user.User, error)
	FindByAuth0Sub(context.Context, string) (*user.User, error)
	FindByEmail(context.Context, string) (*user.User, error)
	FindByName(context.Context, string) (*user.User, error)
	FindByNameOrEmail(context.Context, string) (*user.User, error)
	FindByVerification(context.Context, string) (*user.User, error)
	FindByPasswordResetRequest(context.Context, string) (*user.User, error)
	Save(context.Context, *user.User) error
	Remove(context.Context, id.UserID) error
}
