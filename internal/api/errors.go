package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jacksonfishburn/go-cab/internal/caberr"
)

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusInternalServerError
	msg := "internal server error"

	switch {

	case errors.Is(err, caberr.ErrNotFound):
		status = http.StatusNotFound

	case errors.Is(err, caberr.ErrAlreadyExists):
		status = http.StatusConflict
	}

	var ce *caberr.CabErr
	if errors.As(err, &ce) && status < 500 && ce.Message != "" {
		msg = ce.Message
	}

	if status >= 500 {
		log.Printf("%s %s: %v", r.Method, r.URL.Path, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})

}
