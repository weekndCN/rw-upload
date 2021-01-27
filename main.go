package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/weekndCN/rw-upload/handler"
)

// Handler .
func Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/tree", handler.HandleTree()).Methods(http.MethodGet)
	r.HandleFunc("/upload", handler.HandleUpload()).Methods(http.MethodPost, http.MethodOptions)
	r.PathPrefix("/download").Handler(handler.HandleDownload())
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./uploads")))
	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":9090", Handler()))
}
