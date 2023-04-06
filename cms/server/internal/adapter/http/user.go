package http

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/user"

	"golang.org/x/text/language"
)

type UserController struct {
	usecase interfaces.User
}

func NewUserController(usecase interfaces.User) *UserController {
	return &UserController{
		usecase: usecase,
	}
}

type PasswordResetInput struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Password string `json:"password"`
}

type SignupInput struct {
	Sub         *string         `json:"sub"`
	Secret      *string         `json:"secret"`
	UserID      *id.UserID      `json:"userId"`
	WorkspaceID *id.WorkspaceID `json:"workspaceId"`
	Name        string          `json:"username"`
	Email       string          `json:"email"`
	Password    string          `json:"password"`
	Theme       *user.Theme     `json:"theme"`
	Lang        *language.Tag   `json:"lang"`
}

type CreateVerificationInput struct {
	Email string `json:"email"`
}

type VerifyUserOutput struct {
	UserID   string `json:"userId"`
	Verified bool   `json:"verified"`
}

type CreateUserInput struct {
	Sub         string          `json:"sub"`
	Secret      string          `json:"secret"`
	UserID      *id.UserID      `json:"userId"`
	WorkspaceID *id.WorkspaceID `json:"workspaceId"`
}

type SignupOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *UserController) Signup(ctx context.Context, input SignupInput) (SignupOutput, error) {

	if input.Sub != nil && *input.Sub != "" && input.Email != "" && input.Name != "" {
		u, err := c.usecase.SignupOIDC(ctx, interfaces.SignupOIDC{
			Name:   input.Name,
			Email:  input.Email,
			Sub:    *input.Sub,
			Secret: input.Secret,
		})

		if err != nil {
			return SignupOutput{}, err
		}

		return SignupOutput{
			ID:    u.ID().String(),
			Name:  u.Name(),
			Email: u.Email(),
		}, nil
	}

	u, err := c.usecase.Signup(ctx, interfaces.SignupParam{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		Secret:      input.Secret,
		UserID:      input.UserID,
		WorkspaceID: input.WorkspaceID,
		Lang:        input.Lang,
		Theme:       input.Theme,
	})

	if err != nil {
		return SignupOutput{}, err
	}

	return SignupOutput{
		ID:    u.ID().String(),
		Name:  u.Name(),
		Email: u.Email(),
	}, nil
}
