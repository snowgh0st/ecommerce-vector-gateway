package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/snowgh0st/ecommerce-vector-gateway/models"
	"github.com/snowgh0st/ecommerce-vector-gateway/services"
)

// SearchHandler processes low-latency semantic and visual search requests.
type SearchHandler struct {
	vertexSvc   *services.VertexService
	supabaseSvc *services.SupabaseService
}

func NewSearchHandler(v *services.VertexService, s *services.SupabaseService) *SearchHandler {
	return &SearchHandler{
		vertexSvc:   v,
		supabaseSvc: s,
	}
}

// HandleSearch manages the primary search request flow across Hot, Execution, and Sad paths.
func (h *SearchHandler) HandleSearch(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w.Header().Set("Content-Type", "application/json")

	var req models.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// SAD PATH: Invalid payload structure
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.SearchResponse{
			ErrorNotice: "Invalid JSON request payload",
		})
		return
	}

	if req.Limit <= 0 {
		req.Limit = 12
	}

	ctx := r.Context()

	// -------------------------------------------------------------------------
	// HOT PATH: Generate vector embedding via Google Vertex AI
	// -------------------------------------------------------------------------
	embedding, err := h.vertexSvc.GenerateTextEmbedding(ctx, req.QueryText)
	if err != nil {
		// SAD PATH / FALLBACK: Vertex AI degraded or timed out -> Fall back to SQL text search
		products, fbErr := h.supabaseSvc.FallbackKeywordSearch(ctx, req.QueryText, req.Limit)
		if fbErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.SearchResponse{
				ErrorNotice: "Both vector and fallback search failed",
			})
			return
		}

		json.NewEncoder(w).Encode(models.SearchResponse{
			Products:    products,
			LatencyMS:   time.Since(startTime).Milliseconds(),
			IsFallback:  true,
			ErrorNotice: "Vector search unavailable. Serving keyword fallback results.",
		})
		return
	}

	// -------------------------------------------------------------------------
	// EXECUTION PATH: Query Supabase pgvector using calculated embedding
	// -------------------------------------------------------------------------
	products, err := h.supabaseSvc.VectorSearch(ctx, embedding, req.Limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.SearchResponse{
			ErrorNotice: "Database vector retrieval error",
		})
		return
	}

	// Successful Hot Path Response
	json.NewEncoder(w).Encode(models.SearchResponse{
		Products:   products,
		LatencyMS:  time.Since(startTime).Milliseconds(),
		IsFallback: false,
	})
}
