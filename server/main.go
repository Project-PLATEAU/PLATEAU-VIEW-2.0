package main

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/eukarya-inc/reearth-plateauview/server/cms/cmswebhook"
	"github.com/eukarya-inc/reearth-plateauview/server/putil"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"golang.org/x/net/http2"
)

func main() {
	log.Infof("reearth-plateauview\n")

	conf := lo.Must(NewConfig())
	log.Infof("config: %s", conf.Print())

	if conf.GCParcent > 0 {
		debug.SetGCPercent(conf.GCParcent)
	}

	logger := log.NewEcho()
	e := echo.New()
	e.HideBanner = true
	e.Logger = logger
	e.HTTPErrorHandler = errorHandler(e.DefaultHTTPErrorHandler)
	e.Validator = &customValidator{validator: validator.New()}
	e.Use(
		middleware.Recover(),
		logger.AccessLogger(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: conf.Origin,
		}),
	)

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	}, putil.NoCacheMiddleware)

	e.GET("/proxy/*", proxyHandlerFunc, ACAOHeaderOverwriteMiddleware)

	services := lo.Must(Services(conf))
	serviceNames := lo.Map(services, func(s *Service, _ int) string { return s.Name })
	webhookHandlers := []cmswebhook.Handler{}
	for _, s := range services {
		if s.Echo != nil {
			g := e.Group("")
			if !s.DisableNoCache {
				g.Use(putil.NoCacheMiddleware)
			}
			lo.Must0(s.Echo(g))
		}
		if s.Webhook != nil {
			webhookHandlers = append(webhookHandlers, s.Webhook)
		}
	}

	cmswebhook.Echo(
		e.Group("/webhook"),
		[]byte(conf.CMS_Webhook_Secret),
		webhookHandlers...,
	)

	log.Infof("enabled services: %v", serviceNames)
	addr := fmt.Sprintf("[::]:%d", conf.Port)
	log.Fatalln(e.StartH2CServer(addr, &http2.Server{}))
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

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func funcName(i interface{}) string {
	return strings.TrimPrefix(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), "main.")
}
