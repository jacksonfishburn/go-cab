package api

import (
	"log"
	"net/http"
	"strings"
)

func (h Handler) Authorize(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Printf("unauthorized request from %s: missing Authorization header", r.RemoteAddr)
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			log.Printf("unauthorized request from %s: malformed Authorization header", r.RemoteAddr)
			http.Error(w, "Authorization header must be in the form: Bearer <token>", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, prefix)
		if token != h.Token {
			log.Printf("unauthorized request from %s: invalid token", r.RemoteAddr)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		
		next.ServeHTTP(w, r)
	}
}