package services

import (
	"context"
	"fmt"
	"time"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"
)

// VertexService manages interactions with Google Vertex AI for multimodal embeddings.
type VertexService struct {
	client    *aiplatform.PredictionClient
	projectID string
	location  string
}

// NewVertexService initializes the Vertex AI Prediction Client.
func NewVertexService(ctx context.Context, projectID, location, credentialsJSON string) (*VertexService, error) {
	client, err := aiplatform.NewPredictionClient(ctx, option.WithCredentialsJSON([]byte(credentialsJSON)))
	if err != nil {
		return nil, fmt.Errorf("failed to create vertex ai client: %w", err)
	}

	return &VertexService{
		client:    client,
		projectID: projectID,
		location:  location,
	}, nil
}

// GenerateTextEmbedding generates a 768-dimensional dense vector representation for text search.
// HOT PATH: Executed synchronously on every search request; optimized for low latency (<20ms target).
func (v *VertexService) GenerateTextEmbedding(ctx context.Context, text string) ([]float32, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // Hard SLA cutoff for Sad Path protection
	defer cancel()

	endpoint := fmt.Sprintf("projects/%s/locations/%s/publishers/google/models/multimodalembedding", v.projectID, v.location)

	// Construct request payload matching Vertex AI Multimodal Embedding specs
	req := &aiplatformpb.PredictRequest{
		Endpoint: endpoint,
		// Payload instance construction omitted for brevity; returns dense embedding slice
	}

	_ = req // Suppress unused var check for initial layout

	// Mock vector output for structural demonstration
	mockEmbedding := make([]float32, 768)
	for i := range mockEmbedding {
		mockEmbedding[i] = 0.01
	}

	return mockEmbedding, nil
}