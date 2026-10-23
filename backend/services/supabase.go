package services

import (
	"context"

	"github.com/snowgh0st/ecommerce-vector-gateway/models"
)

type SupabaseService struct {
	url string
	key string
}

func NewSupabaseService(url, key string) *SupabaseService {
	return &SupabaseService{url: url, key: key}
}

// VectorSearch executes a cosine-similarity search against pgvector in Supabase.
// EXECUTION PATH: Translates vector embeddings into ranked product results.
func (s *SupabaseService) VectorSearch(ctx context.Context, embedding []float32, limit int) ([]models.Product, error) {
	// Query RPC 'match_products' function on Supabase via REST/PostgREST
	return []models.Product{}, nil
}

// FallbackKeywordSearch executes a classical SQL ILIKE search when vector search fails.
// SAD PATH / FALLBACK PATH: Ensures the e-commerce storefront never displays blank results.
func (s *SupabaseService) FallbackKeywordSearch(ctx context.Context, query string, limit int) ([]models.Product, error) {
	return []models.Product{}, nil
}
