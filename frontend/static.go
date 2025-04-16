package frontend

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed dist/*
var embeddedFiles embed.FS

func RegisterFrontend(e *echo.Echo) {
	fsys, err := fs.Sub(embeddedFiles, "dist")
	if err != nil {
		panic(err)
	}

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(fsys),
		HTML5:      true,
	}))
}
