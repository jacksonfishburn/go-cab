package api

import (
	"log"
	"net/http"
	"strings"
)

func (h Handler) Authorize(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		
		const prefix = "Bearer "
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != h.Token {
			log.Printf("unauthorized request from %s: invalid token", r.RemoteAddr)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
