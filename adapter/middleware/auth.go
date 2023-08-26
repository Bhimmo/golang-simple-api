package my_middleware

import (
	"net/http"
	"strings"

	"github.com/Bhimmo/golang-simple-api/pkg/auth"
)

func ValidToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		tokenNew := strings.ReplaceAll(token, "Bearer", "")
		tokenNew = strings.ReplaceAll(tokenNew, " ", "")

		if tokenNew == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		valido := auth.ValidToken(tokenNew)
		if !valido {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
