package render

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorCode .
func ErrorCode(w http.ResponseWriter, err error, status int) {
	JSON(w, &Error{Message: err.Error()}, status)
}

// InternalError .
func InternalError(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 500)
}

// InternalErrorf .
func InternalErrorf(w http.ResponseWriter, format string, a ...interface{}) {
	ErrorCode(w, fmt.Errorf(format, a...), 500)
}

// NotImplemented .
func NotImplemented(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 501)
}

// NotFound .
func NotFound(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 404)
}

// NotFoundf .
func NotFoundf(w http.ResponseWriter, format string, a ...interface{}) {
	ErrorCode(w, fmt.Errorf(format, a...), 404)
}

// BadRequest .
func BadRequest(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 400)
}

// BadRequestf .
func BadRequestf(w http.ResponseWriter, format string, a ...interface{}) {
	ErrorCode(w, fmt.Errorf(format, a...), 400)
}

// JSON .
func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.Encode(v)
}
