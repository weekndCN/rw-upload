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
	r.HandleFunc("/{name}", handler.HandleDownload()).Methods(http.MethodGet)
	return r
}

func main() {
	/**
	r := mux.NewRouter()
	r.HandleFunc("/upload", handler.HandleUpload()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/{name}", handler.HandleDownload()).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":9090", r))
	*/
	log.Fatal(http.ListenAndServe(":9090", Handler()))
}
