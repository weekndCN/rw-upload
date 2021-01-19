package handler

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/weekndCN/rw-upload/api"
	"github.com/weekndCN/rw-upload/render"
)

// HandleTree tree list view
func HandleTree() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baseDir := api.Basedir()
		if baseDir == "" {
			render.BadRequest(w, render.ErrNotFound)
			return
		}
		staticPath := path.Join(baseDir, staticDir)
		// if static not exists
		err := os.MkdirAll(staticPath, os.ModePerm)
		if err != nil {
			log.Println(err.Error())
			render.InternalError(w, err)
			return
		}

		// file tree linked list
		rootFile := api.FileIterate(staticDir)
		render.JSON(w, rootFile, 200)
	}
}
