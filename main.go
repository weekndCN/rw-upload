package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	config "github.com/weekndCN/rw-upload/config"
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

	// load file
	c := config.Config{}

	s := config.Service{
		Config: &c,
		Path:   "config.yaml",
	}
	s.Reload()
	// watch file
	// go configService.Watch(time.Second * 10)
	go s.Watch(time.Second * 10)

	fmt.Println(c.GetAll())

	r := mux.NewRouter()
	r.HandleFunc("/upload", handler.HandleUpload()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/{name}", handler.HandleDownload()).Methods(http.MethodGet)
	// log.Fatal(http.ListenAndServe(":9090", Handler()))
	log.Fatal(http.ListenAndServe(":9090", r))
}
