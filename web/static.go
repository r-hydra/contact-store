package web

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

func StaticRouter(r *gin.Engine) {
	contentFS, _ := fs.Sub(staticFiles, "static")
	r.NoRoute(static.Serve("/", &StaticFS{
		fs:  http.FS(contentFS),
	}))
}

type StaticFS struct {
	fs  http.FileSystem
}

func (f *StaticFS) Exists(_ string, filepath string) bool {
	if filepath == "/" {
		filepath = "/index.html"
	}

	if _, err := f.fs.Open(filepath); err != nil {
		return false
	}
	return true
}

func (f *StaticFS) Open(name string) (http.File, error) {
	return f.fs.Open(name)
}
