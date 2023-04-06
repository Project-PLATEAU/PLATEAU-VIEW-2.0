package cmswebhook

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reearth/reearthx/log"
)

func EchoMiddleware(secret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return errors.New("unprocessable entity")
			}

			sig := c.Request().Header.Get(SignatureHeader)
			log.Debugf("webhook: received: sig=%s", sig)
			if !validateSignature(sig, body, secret) {
				return c.JSON(http.StatusUnauthorized, "unauthorized")
			}

			log.Debugf("webhook: body: %s", body)

			p := &Payload{}
			if err := json.Unmarshal(body, p); err != nil {
				return c.JSON(http.StatusBadRequest, "invalid payload")
			}

			p.Body = body
			p.Sig = sig
			req := c.Request()
			ctx := AttachPayload(req.Context(), p)
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}

func Echo(g *echo.Group, secret []byte, handlers ...Handler) {
	g.POST("", func(c echo.Context) error {
		ctx := c.Request().Context()
		w := GetPayload(ctx)
		if w == nil {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}

		if err := c.JSON(http.StatusOK, "ok"); err != nil {
			return err
		}

		for _, h := range handlers {
			if err := h(c.Request(), w); err != nil {
				return err
			}
		}

		return nil
	}, EchoMiddleware(secret))
}
