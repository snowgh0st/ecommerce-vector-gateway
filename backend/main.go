package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/snowgh0st/ecommerce-vector-gateway/config"
	"github.com/snowgh0st/ecommerce-vector-gateway/handlers"
	"github.com/snowgh0st/ecommerce-vector-gateway/services"
)

// corsMiddleware allows frontend apps running on Next.js (port 3000) to communicate with the Go backend (port 8080).
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS requests from Next.js
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	cfg := config.LoadConfig()
	ctx := context.Background()

	log.Println("Initializing E-Commerce Vector Search Gateway services...")

	// Initialize Google Vertex AI service
	vertexSvc, err := services.NewVertexService(ctx, cfg.GCPProjectID, cfg.GCPLocation, cfg.GCPCredentialsJSON)
	if err != nil {
		log.Printf("Warning: Failed to initialize Vertex AI client (%v). Falling back to mock embeddings.", err)
	}

	// Initialize Supabase database client
	supabaseSvc := services.NewSupabaseService(cfg.SupabaseURL, cfg.SupabaseKey)

	// Initialize HTTP handlers
	searchHandler := handlers.NewSearchHandler(vertexSvc, supabaseSvc)

	// Route registrations
	http.HandleFunc("/health", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
	}))

	http.HandleFunc("/api/search", corsMiddleware(searchHandler.HandleSearch))

	// Server startup
	serverAddr := ":" + cfg.Port
	log.Printf("🚀 Vector Gateway microservice running on http://localhost%s", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Server stopped unexpectedly: %v", err)
	}
}
