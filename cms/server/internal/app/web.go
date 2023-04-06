package app

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/afero"
)

type WebConfig map[string]string

func Web(e *echo.Echo, wc WebConfig, ac *AuthConfig, disabled bool, fs afero.Fs) {
	if disabled {
		return
	}

	if fs == nil {
		fs = afero.NewOsFs()
	}
	if _, err := fs.Stat("web"); err != nil {
		return // web won't be delivered
	}

	e.Logger.Info("web: web directory will be delivered\n")

	config := map[string]string{}
	if ac != nil {
		if ac.ISS != "" {
			config["auth0Domain"] = strings.TrimSuffix(ac.ISS, "/")
		}
		if ac.ClientID != nil {
			config["auth0ClientId"] = *ac.ClientID
		}
		if len(ac.AUD) > 0 {
			config["auth0Audience"] = ac.AUD[0]
		}
	}

	for k, v := range wc {
		config[k] = v
	}

	e.GET("/reearth_config.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, config)
	})

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "web",
		Index:      "index.html",
		Browse:     false,
		HTML5:      true,
		Filesystem: afero.NewHttpFs(fs),
	}))
}
