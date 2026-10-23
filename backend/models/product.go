package models

import "time"

// Product represents the core e-commerce catalog item.
type Product struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	InStock     bool      `json:"in_stock"`
	CreatedAt   time.Time `json:"created_at"`
}

// SearchRequest represents incoming search queries (text or base64 image).
type SearchRequest struct {
	QueryText  string `json:"query_text,omitempty"`
	ImageBase64 string `json:"image_base64,omitempty"`
	Limit      int    `json:"limit"`
}

// SearchResponse delivers products along with latency telemetry for monitoring.
type SearchResponse struct {
	Products    []Product `json:"products"`
	LatencyMS   int64     `json:"latency_ms"`
	IsFallback  bool      `json:"is_fallback"`
	ErrorNotice string    `json:"error_notice,omitempty"`
}