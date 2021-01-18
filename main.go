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
	r.HandleFunc("/upload", handler.HandleUpload()).Methods(http.MethodPost, http.MethodOptions)
	r.PathPrefix("/download").Handler(handler.HandleDownload()).S
	// r.HandleFunc("/download/{name}", handler.HandleDownload()).Methods(http.MethodGet)
	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":9090", Handler()))
}
