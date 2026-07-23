package api

import (
	"net/http"
)

func Peek(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}