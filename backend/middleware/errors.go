package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/snowgh0st/ecommerce-vector-gateway/models"
)

// RecoveryMiddleware catches runtime panics in HTTP handlers to prevent microservice crashes.
// SAD PATH: Logs stack trace details internally while returning a clean 500 error to clients.
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC RECOVERY] Internal server panic caught: %v", err)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(w).Encode(models.SearchResponse{
					IsFallback:  true,
					ErrorNotice: "An unexpected internal gateway error occurred.",
				})
			}
		}()

		next(w, r)
	}
}
