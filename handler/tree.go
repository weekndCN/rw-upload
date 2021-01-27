package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/weekndCN/rw-upload/api"
	"github.com/weekndCN/rw-upload/render"
)

// HandleTree tree list view
func HandleTree() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baseDir, staticPath, _ := api.Basedir(staticDir, "")
		if baseDir == "" {
			render.BadRequest(w, render.ErrNotFound)
			return
		}
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
