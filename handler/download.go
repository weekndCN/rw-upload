package handler

import (
	"fmt"
	"net/http"
)

// HandleDownload handler download func
func HandleDownload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("err")
	}
}
