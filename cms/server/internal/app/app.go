package app

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reearth/reearth-cms/server/internal/adapter"
	"github.com/reearth/reearth-cms/server/internal/adapter/integration"
	"github.com/reearth/reearth-cms/server/internal/adapter/publicapi"
	"github.com/reearth/reearth-cms/server/internal/usecase/interactor"
	"github.com/reearth/reearthx/appx"
	rlog "github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func initEcho(ctx context.Context, cfg *ServerConfig) *echo.Echo {
	if cfg.Config == nil {
		log.Fatalln("ServerConfig.Config is nil")
	}

	e := echo.New()
	e.Debug = cfg.Debug
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = errorHandler(e.DefaultHTTPErrorHandler)

	// basic middleware
	logger := rlog.NewEcho()
	e.Logger = logger
	e.Use(
		logger.AccessLogger(),
		middleware.Recover(),
		otelecho.Middleware("reearth-cms"),
	)
	origins := allowedOrigins(cfg)
	if len(origins) > 0 {
		e.Use(
			middleware.CORSWithConfig(middleware.CORSConfig{
				AllowOrigins: origins,
			}),
		)
	}

	// GraphQL Playground without auth
	if cfg.Debug || cfg.Config.Dev {
		e.GET("/graphql", echo.WrapHandler(
			playground.Handler("reearth-cms", "/api/graphql"),
		))
		log.Printf("gql: GraphQL Playground is available")
	}

	internalJWTMiddleware := echo.WrapMiddleware(lo.Must(
		appx.AuthMiddleware(cfg.Config.JWTProviders(), adapter.ContextAuthInfo, true),
	))
	m2mJWTMiddleware := echo.WrapMiddleware(lo.Must(
		appx.AuthMiddleware(cfg.Config.AuthM2M.JWTProvider(), adapter.ContextAuthInfo, false),
	))
	usecaseMiddleware := UsecaseMiddleware(cfg.Repos, cfg.Gateways, interactor.ContainerConfig{
		SignupSecret:    cfg.Config.SignupSecret,
		AuthSrvUIDomain: cfg.Config.Host_Web,
	})

	// apis
	api := e.Group("/api", private)
	api.GET("/ping", Ping())
	api.POST(
		"/graphql", GraphqlAPI(cfg.Config.GraphQL, cfg.Config.Dev),
		internalJWTMiddleware,
		authMiddleware(cfg),
		usecaseMiddleware,
	)
	api.POST(
		"/notify", NotifyHandler(),
		m2mJWTMiddleware,
		M2MAuthMiddleware(cfg.Config.AuthM2M.Email),
		usecaseMiddleware,
	)
	api.POST("/signup", Signup(), usecaseMiddleware)

	publicapi.Echo(api.Group("/p", PublicAPIAuthMiddleware(cfg), usecaseMiddleware))
	integration.RegisterHandlers(api.Group(
		"",
		authMiddleware(cfg),
		AuthRequiredMiddleware(),
		usecaseMiddleware,
		private,
	), integration.NewStrictHandler(integration.NewServer(), nil))

	serveFiles(e, cfg.Gateways.File)
	Web(e, cfg.Config.Web, cfg.Config.AuthForWeb(), cfg.Config.Web_Disabled, nil)
	return e
}

func allowedOrigins(cfg *ServerConfig) []string {
	if cfg == nil {
		return nil
	}
	origins := append([]string{}, cfg.Config.Origins...)
	if cfg.Debug {
		origins = append(origins, "http://localhost:3000", "http://127.0.0.1:3000", "http://localhost:8080")
	}
	return origins
}

func errorMessage(err error, log func(string, ...interface{})) (int, string) {
	code := http.StatusBadRequest
	msg := err.Error()

	if err2, ok := err.(*echo.HTTPError); ok {
		code = err2.Code
		if msg2, ok := err2.Message.(string); ok {
			msg = msg2
		} else if msg2, ok := err2.Message.(error); ok {
			msg = msg2.Error()
		} else {
			msg = "error"
		}
		if err2.Internal != nil {
			log("echo internal err: %+v", err2)
		}
	} else if errors.Is(err, rerror.ErrNotFound) {
		code = http.StatusNotFound
		msg = "not found"
	} else {
		if ierr := rerror.UnwrapErrInternal(err); ierr != nil {
			code = http.StatusInternalServerError
			msg = "internal server error"
		}
	}

	return code, msg
}

func errorHandler(next func(error, echo.Context)) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		code, msg := errorMessage(err, func(f string, args ...interface{}) {
			c.Echo().Logger.Errorf(f, args...)
		})
		if err := c.JSON(code, map[string]string{
			"error": msg,
		}); err != nil {
			next(err, c)
		}
	}
}

func private(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderCacheControl, "private, no-store, no-cache, must-revalidate")
		return next(c)
	}
}
