package api

import (
	"net/http"
)

func Del(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}