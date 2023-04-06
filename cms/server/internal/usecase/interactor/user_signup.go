package interactor

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	htmlTmpl "html/template"
	"net/http"
	"net/url"
	"path"
	"strings"
	textTmpl "text/template"

	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
)

type mailContent struct {
	UserName    string
	Message     string
	Suffix      string
	ActionLabel string
	ActionURL   htmlTmpl.URL
}

type OpenIDConfiguration struct {
	UserinfoEndpoint string `json:"userinfo_endpoint"`
}

type UserInfo struct {
	Sub      string `json:"sub"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Error    string `json:"error"`
}

var (
	//go:embed emails/auth_html.tmpl
	autHTMLTMPLStr string
	//go:embed emails/auth_text.tmpl
	authTextTMPLStr string

	authTextTMPL *textTmpl.Template
	authHTMLTMPL *htmlTmpl.Template
)

func init() {
	var err error
	authTextTMPL, err = textTmpl.New("passwordReset").Parse(authTextTMPLStr)
	if err != nil {
		log.Panicf("password reset email template parse error: %s\n", err)
	}
	authHTMLTMPL, err = htmlTmpl.New("passwordReset").Parse(autHTMLTMPLStr)
	if err != nil {
		log.Panicf("password reset email template parse error: %s\n", err)
	}
}

func (i *User) Signup(ctx context.Context, param interfaces.SignupParam) (u *user.User, err error) {
	if err := i.verifySignupSecret(param.Secret); err != nil {
		return nil, err
	}

	u, workspace, err := user.Init(user.InitParams{
		Email:       param.Email,
		Name:        param.Name,
		Password:    lo.ToPtr(param.Password),
		Lang:        param.Lang,
		Theme:       param.Theme,
		UserID:      param.UserID,
		WorkspaceID: param.WorkspaceID,
	})
	if err != nil {
		return nil, err
	}

	vr := user.NewVerification()
	u.SetVerification(vr)

	if err := i.repos.User.Create(ctx, u); err != nil {
		if errors.Is(err, repo.ErrDuplicatedUser) {
			return nil, interfaces.ErrUserAlreadyExists
		}
		return nil, err
	}
	if err := i.repos.Workspace.Save(ctx, workspace); err != nil {
		return nil, err
	}

	if err := i.sendVerificationMail(ctx, u, vr); err != nil {
		return nil, err
	}

	return u, nil
}

func (i *User) SignupOIDC(ctx context.Context, param interfaces.SignupOIDC) (*user.User, error) {
	if err := i.verifySignupSecret(param.Secret); err != nil {
		return nil, err
	}
	if param.Sub == "" || param.Name == "" || param.Email == "" {
		return nil, rerror.NewE(i18n.T("invalid parameters"))
	}

	eu, err := i.repos.User.FindByEmail(ctx, param.Email)
	if err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return nil, err
	}
	if eu != nil {
		return nil, repo.ErrDuplicatedUser
	}

	eu, err = i.repos.User.FindBySub(ctx, param.Sub)
	if err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return nil, err
	}
	if eu != nil {
		return nil, repo.ErrDuplicatedUser
	}

	u, workspace, err := user.Init(user.InitParams{
		Email: param.Email,
		Name:  param.Name,
		Sub:   user.AuthFromAuth0Sub(param.Sub).Ref(),
	})
	if err != nil {
		return nil, err
	}

	if err := i.repos.User.Create(ctx, u); err != nil {
		return nil, err
	}

	if err := i.repos.Workspace.Save(ctx, workspace); err != nil {
		return nil, err
	}

	return u, nil
}

func (i *User) FindOrCreate(ctx context.Context, param interfaces.UserFindOrCreateParam) (u *user.User, err error) {
	return Run1(ctx, nil, i.repos, Usecase().Transaction(), func(ctx context.Context) (*user.User, error) {
		if param.Sub == "" {
			return nil, rerror.ErrNotFound
		}

		// Check if user already exists
		existedUser, err := i.repos.User.FindBySub(ctx, param.Sub)
		if err != nil && !errors.Is(err, rerror.ErrNotFound) {
			return nil, err
		} else if existedUser != nil {
			return existedUser, nil
		}

		ui, err := getUserInfoFromISS(ctx, param.ISS, param.Token)
		if err != nil {
			return nil, err
		}

		u, workspace, err := user.Init(user.InitParams{
			Email: ui.Email,
			Name:  ui.Name,
			Sub:   user.AuthFromAuth0Sub(param.Sub).Ref(),
		})
		if err != nil {
			return nil, err
		}

		u2, err := i.repos.User.FindBySubOrCreate(ctx, u, param.Sub)
		if err != nil {
			return nil, err
		}

		if err := i.repos.Workspace.Save(ctx, workspace); err != nil {
			return nil, err
		}

		return u2, nil
	})
}

func (i *User) sendVerificationMail(ctx context.Context, u *user.User, vr *user.Verification) error {
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

	if err := i.gateways.Mailer.SendMail(
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

func getUserInfoFromISS(ctx context.Context, iss, accessToken string) (UserInfo, error) {
	if accessToken == "" {
		return UserInfo{}, rerror.NewE(i18n.T("invalid access token"))
	}
	if iss == "" {
		return UserInfo{}, rerror.NewE(i18n.T("invalid issuer"))
	}

	var u string
	c, err := getOpenIDConfiguration(ctx, iss)
	if err != nil {
		u2 := issToURL(iss, "/userinfo")
		if u2 == nil {
			return UserInfo{}, rerror.NewE(i18n.T("invalid iss"))
		}
		u = u2.String()
	} else {
		u = c.UserinfoEndpoint
	}
	return getUserInfo(ctx, u, accessToken)
}

func getOpenIDConfiguration(ctx context.Context, iss string) (c OpenIDConfiguration, err error) {
	WKUrl := issToURL(iss, "/.well-known/openid-configuration")
	if WKUrl == nil {
		err = rerror.NewE(i18n.T("invalid iss"))
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	req, err2 := http.NewRequestWithContext(ctx, http.MethodGet, WKUrl.String(), nil)
	if err2 != nil {
		err = err2
		return
	}

	res, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		err = err2
		return
	}

	if res.StatusCode != http.StatusOK {
		err = rerror.NewE(i18n.T("could not get user info"))
		return
	}

	if err2 := json.NewDecoder(res.Body).Decode(&c); err2 != nil {
		err = fmt.Errorf("could not get user info: %w", err2)
		return
	}

	return
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

	if res.StatusCode != http.StatusOK {
		err = rerror.NewE(i18n.T("could not get user info"))
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

func (i *User) verifySignupSecret(secret *string) error {
	if i.signupSecret != "" && (secret == nil || *secret != i.signupSecret) {
		return interfaces.ErrSignupInvalidSecret
	}
	return nil
}
