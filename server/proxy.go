package main

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func proxyHandlerFunc(c echo.Context) error {
	// Extract the target URL from the request path
	targetPath := c.Param("*")

	// This shouldn't be done by us but It'll do for now: @pyshx
	if strings.HasPrefix(targetPath, "http:/") && len(targetPath) > 6 && targetPath[6] != '/' {
		targetPath = "http://" + strings.TrimPrefix(targetPath, "http:/")
	} else if strings.HasPrefix(targetPath, "https:/") && len(targetPath) > 7 && targetPath[7] != '/' {
		targetPath = "https://" + strings.TrimPrefix(targetPath, "https:/")
	}

	targetURL, err := url.Parse(targetPath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid target URL",
		})
	}

	// Define the ProxyConfig object with custom Rewrite rules and ModifyResponse function
	proxyConfig := middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: targetURL,
			},
		}),
	}

	// Create a new middleware.ProxyWithConfig() middleware with the target URL
	proxyMiddleware := middleware.ProxyWithConfig(proxyConfig)(func(c echo.Context) error {
		return nil
	})

	// Invoke the proxy middleware to handle the request and return the response
	if err := proxyMiddleware(c); err != nil {
		return err
	}

	return nil
}

func setResponseACAOHeaderFromRequest(c echo.Context) {
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin,
		c.Request().Header.Get(echo.HeaderOrigin))
}

func ACAOHeaderOverwriteMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Before(func() {
			setResponseACAOHeaderFromRequest(c)
		})
		return next(c)
	}
}
