package app

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/reearth/reearth-cms/server/internal/adapter"
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/interactor"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/appx"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
)

var (
	debugUserHeaderKey        = "X-Reearth-Debug-User"
	debugIntegrationHeaderKey = "X-Reearth-Debug-Integration"
)

func authMiddleware(cfg *ServerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			ctx := req.Context()

			ctx, err = attachUserOperator(ctx, req, cfg)
			if err != nil {
				return err
			}

			ctx, err = attachIntegrationOperator(ctx, req, cfg)
			if err != nil {
				return err
			}

			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}

func attachUserOperator(ctx context.Context, req *http.Request, cfg *ServerConfig) (context.Context, error) {
	var u *user.User
	if ai := adapter.GetAuthInfo(ctx); ai != nil {
		var err error
		userUsecase := interactor.NewUser(cfg.Repos, cfg.Gateways, cfg.Config.SignupSecret, cfg.Config.Host_Web)
		u, err = userUsecase.FindOrCreate(ctx, interfaces.UserFindOrCreateParam{
			Sub:   ai.Sub,
			ISS:   ai.Iss,
			Token: ai.Token,
		})
		if err != nil {
			return nil, err
		}
	}

	if cfg.Debug {
		if val := req.Header.Get(debugUserHeaderKey); val != "" {
			uId, err := id.UserIDFrom(val)
			if err != nil {
				return nil, err
			}
			u, err = cfg.Repos.User.FindByID(ctx, uId)
			if err != nil {
				return nil, err
			}
		}
	}

	// generate operator
	if u != nil {
		defaultLang := req.Header.Get("Accept-Language")
		op, err := generateUserOperator(ctx, cfg, u, defaultLang)
		if err != nil {
			return nil, err
		}

		ctx = adapter.AttachUser(ctx, u)
		ctx = adapter.AttachOperator(ctx, op)
	}

	return ctx, nil
}

func attachIntegrationOperator(ctx context.Context, req *http.Request, cfg *ServerConfig) (context.Context, error) {
	var i *integration.Integration
	if token := getIntegrationToken(req); token != "" {
		var err error
		i, err = cfg.Repos.Integration.FindByToken(ctx, token)
		if err != nil {
			if errors.Is(err, rerror.ErrNotFound) {
				return nil, echo.ErrUnauthorized
			}
			return nil, err
		}
	}

	if cfg.Debug {
		if val := req.Header.Get(debugIntegrationHeaderKey); val != "" {
			iId, err := id.IntegrationIDFrom(val)
			if err != nil {
				return nil, err
			}
			i, err = cfg.Repos.Integration.FindByID(ctx, iId)
			if err != nil {
				return nil, err
			}
		}
	}

	if i != nil {
		defaultLang := req.Header.Get("Accept-Language")
		op, err := generateIntegrationOperator(ctx, cfg, i, defaultLang)
		if err != nil {
			return nil, err
		}

		ctx = adapter.AttachOperator(ctx, op)
	}

	return ctx, nil
}

func PublicAPIAuthMiddleware(cfg *ServerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// TODO: support limit publication scope

			// req := c.Request()
			// ctx := req.Context()
			// if token := req.Header.Get("Reearth-Token"); token != "" {
			// 	ws, err := cfg.Repos.Project.FindByPublicAPIToken(ctx, token)
			// 	if err != nil {
			// 		if errors.Is(err, rerror.ErrNotFound) {
			// 			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			// 		}
			// 		return err
			// 	}

			// 	if !ws.IsPublic() {
			// 		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			// 	}

			// 	c.SetRequest(req.WithContext(adapter.AttachOperator(ctx, &usecase.Operator{
			// 		PublicAPIProject: ws.ID(),
			// 	})))
			// }

			return next(c)
		}
	}
}

func M2MAuthMiddleware(email string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			if ai, ok := ctx.Value(adapter.ContextAuthInfo).(appx.AuthInfo); ok {
				if ai.EmailVerified == nil || !*ai.EmailVerified || ai.Email != email {
					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
				}
				op, err := generateMachineOperator(ctx)
				if err != nil {
					return err
				}
				ctx = adapter.AttachOperator(ctx, op)
				c.SetRequest(c.Request().WithContext(ctx))
			}
			return next(c)
		}
	}
}

func getIntegrationToken(req *http.Request) string {
	token := strings.TrimPrefix(req.Header.Get("authorization"), "Bearer ")
	if strings.HasPrefix(token, "secret_") {
		return token
	}
	return ""
}

func generateUserOperator(ctx context.Context, cfg *ServerConfig, u *user.User, defaultLang string) (*usecase.Operator, error) {
	if u == nil {
		return nil, nil
	}

	uid := u.ID()

	w, err := cfg.Repos.Workspace.FindByUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	rw := w.FilterByUserRole(uid, user.RoleReader).IDs()
	ww := w.FilterByUserRole(uid, user.RoleWriter).IDs()
	mw := w.FilterByUserRole(uid, user.RoleMaintainer).IDs()
	ow := w.FilterByUserRole(uid, user.RoleOwner).IDs()

	rp, wp, mp, op, err := operatorProjects(ctx, cfg, w, rw, ww, mw, ow)
	if err != nil {
		return nil, err
	}

	lang := u.Lang().String()
	if lang == "" || lang == "und" {
		lang = defaultLang
	}

	return &usecase.Operator{
		User:        &uid,
		Integration: nil,

		Lang: lang,

		ReadableWorkspaces:     rw,
		WritableWorkspaces:     ww,
		MaintainableWorkspaces: mw,
		OwningWorkspaces:       ow,

		ReadableProjects:     rp,
		WritableProjects:     wp,
		MaintainableProjects: mp,
		OwningProjects:       op,
	}, nil
}

func operatorProjects(ctx context.Context, cfg *ServerConfig, w user.WorkspaceList, rw, ww, mw, ow user.WorkspaceIDList) (id.ProjectIDList, id.ProjectIDList, id.ProjectIDList, id.ProjectIDList, error) {
	rp := id.ProjectIDList{}
	wp := id.ProjectIDList{}
	mp := id.ProjectIDList{}
	op := id.ProjectIDList{}

	var cur *usecasex.Cursor
	for {
		projects, pi, err := cfg.Repos.Project.FindByWorkspaces(ctx, w.IDs(), usecasex.CursorPagination{
			After: cur,
			First: lo.ToPtr(int64(100)),
		}.Wrap())
		if err != nil {
			return nil, nil, nil, nil, err
		}

		for _, p := range projects {
			if ow.Has(p.Workspace()) {
				op = append(op, p.ID())
			} else if mw.Has(p.Workspace()) {
				mp = append(mp, p.ID())
			} else if ww.Has(p.Workspace()) {
				wp = append(wp, p.ID())
			} else if rw.Has(p.Workspace()) {
				rp = append(rp, p.ID())
			}
		}

		if !pi.HasNextPage {
			break
		}
		cur = pi.EndCursor
	}
	return rp, wp, op, mp, nil
}

func generateIntegrationOperator(ctx context.Context, cfg *ServerConfig, i *integration.Integration, lang string) (*usecase.Operator, error) {
	if i == nil {
		return nil, nil
	}

	iId := i.ID()
	w, err := cfg.Repos.Workspace.FindByIntegration(ctx, iId)
	if err != nil {
		return nil, err
	}

	rw := w.FilterByIntegrationRole(iId, user.RoleReader).IDs()
	ww := w.FilterByIntegrationRole(iId, user.RoleWriter).IDs()
	mw := w.FilterByIntegrationRole(iId, user.RoleMaintainer).IDs()
	ow := w.FilterByIntegrationRole(iId, user.RoleOwner).IDs()

	rp, wp, mp, op, err := operatorProjects(ctx, cfg, w, rw, ww, mw, ow)
	if err != nil {
		return nil, err
	}

	return &usecase.Operator{
		User:                   nil,
		Integration:            &iId,
		Lang:                   lang,
		ReadableWorkspaces:     rw,
		WritableWorkspaces:     ww,
		MaintainableWorkspaces: mw,
		OwningWorkspaces:       ow,

		ReadableProjects:     rp,
		WritableProjects:     wp,
		MaintainableProjects: mp,
		OwningProjects:       op,
	}, nil
}

func generateMachineOperator(ctx context.Context) (*usecase.Operator, error) {
	return &usecase.Operator{
		User:        nil,
		Integration: nil,
		Machine:     true,
	}, nil
}

func AuthRequiredMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			if adapter.Operator(ctx) == nil {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
