package ui

import (
	"embed"
	"github.com/gin-gonic/gin"
	"monogoly/env"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
)

var (
	//go:embed dist
	distFS embed.FS
	HttpFS = http.FS(distFS)
)

func BindUiRoutes(e *gin.Engine) {
	if env.DevMode && len(env.UiDevUrl) > 0 {
		e.NoRoute(func(ctx *gin.Context) {
			if strings.Contains(strings.ToLower(ctx.Request.URL.Path), "/api/") {
				ctx.JSON(404, gin.H{"message": "Route not found"})
				return
			}

			proxyRequest(ctx)
		})
	} else {
		e.StaticFileFS("/index.html", "dist/main.html", HttpFS)
		e.StaticFileFS("/", "dist/main.html", HttpFS)

		e.NoRoute(func(ctx *gin.Context) {
			if strings.Contains(strings.ToLower(ctx.Request.URL.Path), "/api/") {
				ctx.JSON(404, gin.H{"message": "Route not found"})
				return
			}

			//SPA logic
			if strings.Contains(ctx.Request.URL.Path, ".") {
				ctx.FileFromFS(path.Join("dist/", ctx.Request.URL.Path), HttpFS)
			} else {
				ctx.FileFromFS("dist/main.html", HttpFS)
			}
		})
	}
}

func proxyRequest(c *gin.Context) {
	remote, err := url.Parse(env.UiDevUrl)

	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
