package handler

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/weekndCN/rw-upload/api"
	"github.com/weekndCN/rw-upload/render"
)

const (
	tmpDir    = "temp"
	staticDir = "uploads"
)

// HandleDownload handler download func
func HandleDownload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get path file
		name := strings.TrimPrefix(r.URL.Path, "/download/")
		fmt.Println("name:", name)

		baseDir := api.Basedir()
		if baseDir == "" {
			render.BadRequest(w, render.ErrNotFound)
			return
		}

		// static files directory
		staticPath := path.Join(baseDir, staticDir)
		// temp files directory
		tempPath := path.Join(baseDir, tmpDir)
		// request file full path
		fullPath := path.Join(staticPath, name)

		if !api.Exists(fullPath) {
			render.BadRequest(w, render.ErrNotFound)
			return
		}

		// if request file is directory. then server file as zip file. else download file directly
		if api.IsDir(fullPath) {
			// if name is directory then zip folder
			if err := api.Zip(staticPath, tempPath, name); err != nil {
				render.InternalError(w, err)
				return
			}
			http.ServeFile(w, r, path.Join(tempPath, name+".zip"))
			err := api.DelZip(name)
			if err != nil {
				render.InternalError(w, err)
				return
			}
		} else {
			http.ServeFile(w, r, fullPath)
		}
	}
}
