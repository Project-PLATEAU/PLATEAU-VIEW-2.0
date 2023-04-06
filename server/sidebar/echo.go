package sidebar

import (
	"net/http"
	"strings"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	CMSBaseURL string
	CMSToken   string
	AdminToken string
}

func Echo(g *echo.Group, c Config) error {
	cms, err := cms.New(c.CMSBaseURL, c.CMSToken)
	if err != nil {
		return err
	}

	initEcho(g, c, cms)
	return nil
}

func initEcho(g *echo.Group, c Config, cms cms.Interface) {
	h := NewHandler(cms)

	g.Use(middleware.CORS(), middleware.BodyLimit("5M"))

	g.GET("/:pid", h.fetchRoot())
	g.GET("/:pid/data", h.getAllDataHandler())
	g.GET("/:pid/data/:iid", h.getDataHandler())
	g.POST("/:pid/data", h.createDataHandler(), authMiddleware(c.AdminToken))
	g.PATCH("/:pid/data/:iid", h.updateDataHandler(), authMiddleware(c.AdminToken))
	g.DELETE("/:pid/data/:iid", h.deleteDataHandler(), authMiddleware(c.AdminToken))
	g.GET("/:pid/templates", h.fetchTemplatesHandler())
	g.GET("/:pid/templates/:tid", h.fetchTemplateHandler())
	g.POST("/:pid/templates", h.createTemplateHandler(), authMiddleware(c.AdminToken))
	g.PATCH("/:pid/templates/:tid", h.updateTemplateHandler(), authMiddleware(c.AdminToken))
	g.DELETE("/:pid/templates/:tid", h.deleteTemplateHandler(), authMiddleware(c.AdminToken))
}

func authMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			header := req.Header.Get("Authorization")
			token := strings.TrimPrefix(header, "Bearer ")
			if secret == "" || token != secret {
				return c.JSON(http.StatusUnauthorized, nil)
			}
			return next(c)
		}
	}
}
