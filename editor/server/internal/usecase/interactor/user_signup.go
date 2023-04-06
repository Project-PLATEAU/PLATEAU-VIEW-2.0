package interactor

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	htmlTmpl "html/template"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/reearth/reearth/server/internal/usecase/gateway"
	"github.com/reearth/reearth/server/internal/usecase/interfaces"
	"github.com/reearth/reearth/server/pkg/id"
	"github.com/reearth/reearth/server/pkg/user"
	"github.com/reearth/reearth/server/pkg/user/userops"
	"github.com/reearth/reearth/server/pkg/workspace"
	"github.com/reearth/reearthx/rerror"
)

func (i *User) Signup(ctx context.Context, inp interfaces.SignupParam) (*user.User, *workspace.Workspace, error) {
	if inp.Name == "" {
		return nil, nil, interfaces.ErrSignupInvalidName
	}
	if err := i.verifySignupSecret(inp.Secret); err != nil {
		return nil, nil, err
	}

	tx, err := i.transaction.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}

	ctx = tx.Context()
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	// Check if user and workspace already exists
	existedUser, existedWorkspace, err := i.userAlreadyExists(ctx, inp.User.UserID, inp.Sub, &inp.Name, inp.User.WorkspaceID)
	if err != nil {
		return nil, nil, err
	}

	if existedUser != nil {
		if existedUser.Verification() == nil || !existedUser.Verification().IsVerified() {
			// if user exists but not verified -> create a new verification
			if err := i.createVerification(ctx, existedUser); err != nil {
				return nil, nil, err
			}
			return existedUser, existedWorkspace, nil
		}
		return nil, nil, interfaces.ErrUserAlreadyExists
	}

	// Initialize user and workspace
	var auth *user.Auth
	if inp.Sub != nil {
		auth = user.AuthFrom(*inp.Sub).Ref()
	}
	u, ws, err := userops.Init(userops.InitParams{
		Email:       inp.Email,
		Name:        inp.Name,
		Sub:         auth,
		Password:    inp.Password,
		Lang:        inp.User.Lang,
		Theme:       inp.User.Theme,
		UserID:      inp.User.UserID,
		WorkspaceID: inp.User.WorkspaceID,
	})
	if err != nil {
		return nil, nil, err
	}

	if err := i.userRepo.Save(ctx, u); err != nil {
		return nil, nil, err
	}
	if err := i.workspaceRepo.Save(ctx, ws); err != nil {
		return nil, nil, err
	}

	if err := i.createVerification(ctx, u); err != nil {
		return nil, nil, err
	}

	tx.Commit()
	return u, ws, nil
}

func (i *User) SignupOIDC(ctx context.Context, inp interfaces.SignupOIDCParam) (u *user.User, _ *workspace.Workspace, err error) {
	if err := i.verifySignupSecret(inp.Secret); err != nil {
		return nil, nil, err
	}

	sub := inp.Sub
	name := inp.Name
	email := inp.Email
	if sub == "" || email == "" {
		ui, err := getUserInfoFromISS(ctx, inp.Issuer, inp.AccessToken)
		if err != nil {
			return nil, nil, err
		}
		sub = ui.Sub
		email = ui.Email
		if name == "" {
			name = ui.Nickname
		}
		if name == "" {
			name = ui.Name
		}
		if name == "" {
			name = ui.Email
		}
	}

	tx, err := i.transaction.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}

	ctx = tx.Context()
	defer func() {
		if err2 := tx.End(ctx); err == nil && err2 != nil {
			err = err2
		}
	}()

	// Check if user and workspace already exists
	if existedUser, existedWS, err := i.userAlreadyExists(ctx, inp.User.UserID, &sub, &name, inp.User.WorkspaceID); err != nil {
		return nil, nil, err
	} else if existedUser != nil || existedWS != nil {
		return nil, nil, interfaces.ErrUserAlreadyExists
	}

	// Initialize user and ws
	u, ws, err := userops.Init(userops.InitParams{
		Email:       email,
		Name:        name,
		Sub:         user.AuthFrom(sub).Ref(),
		Lang:        inp.User.Lang,
		Theme:       inp.User.Theme,
		UserID:      inp.User.UserID,
		WorkspaceID: inp.User.WorkspaceID,
	})
	if err != nil {
		return nil, nil, err
	}

	if err := i.userRepo.Save(ctx, u); err != nil {
		return nil, nil, err
	}
	if err := i.workspaceRepo.Save(ctx, ws); err != nil {
		return nil, nil, err
	}

	tx.Commit()
	return u, ws, nil
}

func (i *User) verifySignupSecret(secret *string) error {
	if i.signupSecret != "" && (secret == nil || *secret != i.signupSecret) {
		return interfaces.ErrSignupInvalidSecret
	}
	return nil
}

func (i *User) userAlreadyExists(ctx context.Context, userID *id.UserID, sub *string, name *string, workspaceID *id.WorkspaceID) (*user.User, *workspace.Workspace, error) {
	// Check if user already exists
	var existedUser *user.User
	var err error

	if userID != nil {
		existedUser, err = i.userRepo.FindByID(ctx, *userID)
		if err != nil && !errors.Is(err, rerror.ErrNotFound) {
			return nil, nil, err
		}
	} else if sub != nil {
		// Check if user already exists
		existedUser, err = i.userRepo.FindByAuth0Sub(ctx, *sub)
		if err != nil && !errors.Is(err, rerror.ErrNotFound) {
			return nil, nil, err
		}
	} else if name != nil {
		existedUser, err = i.userRepo.FindByName(ctx, *name)
		if err != nil && !errors.Is(err, rerror.ErrNotFound) {
			return nil, nil, err
		}
	}

	if existedUser != nil {
		ws, err := i.workspaceRepo.FindByID(ctx, existedUser.Workspace())
		if err != nil && !errors.Is(err, rerror.ErrNotFound) {
			return nil, nil, err
		}
		return existedUser, ws, nil
	}

	// Check if workspace already exists
	if workspaceID != nil {
		existed, err := i.workspaceRepo.FindByID(ctx, *workspaceID)
		if err != nil && !errors.Is(err, rerror.ErrNotFound) {
			return nil, nil, err
		}
		if existed != nil {
			return nil, existed, nil
		}
	}

	return nil, nil, nil
}

func getUserInfoFromISS(ctx context.Context, iss, accessToken string) (UserInfo, error) {
	if accessToken == "" {
		return UserInfo{}, errors.New("invalid access token")
	}
	if iss == "" {
		return UserInfo{}, errors.New("invalid issuer")
	}

	var u string
	c, err := getOpenIDConfiguration(ctx, iss)
	if err != nil {
		u2 := issToURL(iss, "/userinfo")
		if u2 == nil {
			return UserInfo{}, errors.New("invalid iss")
		}
		u = u2.String()
	} else {
		u = c.UserinfoEndpoint
	}
	return getUserInfo(ctx, u, accessToken)
}

type OpenIDConfiguration struct {
	UserinfoEndpoint string `json:"userinfo_endpoint"`
}

func getOpenIDConfiguration(ctx context.Context, iss string) (c OpenIDConfiguration, err error) {
	url := issToURL(iss, "/.well-known/openid-configuration")
	if url == nil {
		err = errors.New("invalid iss")
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	req, err2 := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err2 != nil {
		err = err2
		return
	}

	res, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		err = err2
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		err = errors.New("could not get user info")
		return
	}
	if err2 := json.NewDecoder(res.Body).Decode(&c); err2 != nil {
		err = fmt.Errorf("could not get user info: %w", err2)
		return
	}
	return
}

type UserInfo struct {
	Sub      string `json:"sub"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Error    string `json:"error"`
}

func getUserInfo(ctx context.Context, url, accessToken string) (ui UserInfo, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	req, err2 := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err2 != nil {
		err = err2
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	res, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		err = err2
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		err = errors.New("could not get user info")
		return
	}

	if err2 := json.NewDecoder(res.Body).Decode(&ui); err2 != nil {
		err = fmt.Errorf("could not get user info: %w", err2)
		return
	}

	if ui.Error != "" {
		err = fmt.Errorf("could not get user info: %s", ui.Error)
		return
	}
	if ui.Sub == "" {
		err = fmt.Errorf("could not get user info: invalid response")
		return
	}
	if ui.Email == "" {
		err = fmt.Errorf("could not get user info: email scope missing")
		return
	}

	return
}

func issToURL(iss, p string) *url.URL {
	if iss == "" {
		return nil
	}

	if !strings.HasPrefix(iss, "https://") && !strings.HasPrefix(iss, "http://") {
		iss = "https://" + iss
	}

	u, err := url.Parse(iss)
	if err == nil {
		u.Path = path.Join(u.Path, p)
		if u.Path == "/" {
			u.Path = ""
		}
		return u
	}

	return nil
}

func (i *User) createVerification(ctx context.Context, u *user.User) error {
	vr := user.NewVerification()
	u.SetVerification(vr)

	if err := i.userRepo.Save(ctx, u); err != nil {
		return err
	}

	var text, html bytes.Buffer
	link := i.authSrvUIDomain + "/?user-verification-token=" + vr.Code()
	signupMailContent := mailContent{
		Message:     "Thank you for signing up to Re:Earth. Please verify your email address by clicking the button below.",
		Suffix:      "You can use this email address to log in to Re:Earth account anytime.",
		ActionLabel: "Activate your account and log in",
		UserName:    u.Email(),
		ActionURL:   htmlTmpl.URL(link),
	}
	if err := authTextTMPL.Execute(&text, signupMailContent); err != nil {
		return err
	}
	if err := authHTMLTMPL.Execute(&html, signupMailContent); err != nil {
		return err
	}

	if err := i.mailer.SendMail(
		[]gateway.Contact{
			{
				Email: u.Email(),
				Name:  u.Name(),
			},
		},
		"email verification",
		text.String(),
		html.String(),
	); err != nil {
		return err
	}

	return nil
}
