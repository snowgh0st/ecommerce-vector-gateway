package config

import (
	"os"
)

// Config holds runtime environment variables and client keys.
type Config struct {
	Port               string
	SupabaseURL        string
	SupabaseKey        string
	GCPProjectID       string
	GCPLocation        string
	GCPCredentialsJSON string
}

// LoadConfig fetches configuration values from environment variables or sensible local defaults.
func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL == "" {
		supabaseURL = "https://your-supabase-project.supabase.co"
	}

	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")
	if supabaseKey == "" {
		supabaseKey = "your-supabase-anon-key"
	}

	gcpProjectID := os.Getenv("GCP_PROJECT_ID")
	if gcpProjectID == "" {
		gcpProjectID = "demo-gcp-project"
	}

	gcpLocation := os.Getenv("GCP_LOCATION")
	if gcpLocation == "" {
		gcpLocation = "us-central1"
	}

	return &Config{
		Port:               port,
		SupabaseURL:        supabaseURL,
		SupabaseKey:        supabaseKey,
		GCPProjectID:       gcpProjectID,
		GCPLocation:        gcpLocation,
		GCPCredentialsJSON: os.Getenv("GCP_CREDENTIALS_JSON"),
	}
}
