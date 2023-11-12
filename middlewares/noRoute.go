package middlewares

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"strings"
)

func embedFolder(fsEmbed fs.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(_ string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

// NoRouteMiddleware is the middleware that processes http 404 errors.
// If the request is an API call, then a JSON 404 error is returned.
// In any other case, if a file exists on the embedded UI filesystem, it is
// served, otherwise the index.html file is served so that the UI can
// render an error page.
func NoRouteMiddleware(urlPrefix string, fsEmbed fs.FS, targetPath string) gin.HandlerFunc {
	filesystem := embedFolder(fsEmbed, targetPath)
	fileserver := http.FileServer(filesystem)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}

	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			// serve a json 404 error if it's an api call
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			// serve the / path
			if !filesystem.Exists(urlPrefix, c.Request.URL.Path) {
				c.Request.URL.Path = "/"
			}
			// serve ui and let it handle the error otherwise
			fileserver.ServeHTTP(c.Writer, c.Request)
		}
		c.Abort()
	}
}
