package api

import (
	"net/http"
)

func Grab(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}