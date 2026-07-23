package api

import (
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}