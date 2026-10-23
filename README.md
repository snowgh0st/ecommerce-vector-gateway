# Ultra-Low Latency E-Commerce Vector Search & Recommendation Gateway

<div align="center">

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Next.js](https://img.shields.io/badge/Next.js-15-000000?style=for-the-badge&logo=nextdotjs&logoColor=white)](https://nextjs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-3178C6?style=for-the-badge&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-v4-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)](https://tailwindcss.com/)
[![Supabase](https://img.shields.io/badge/Supabase-pgvector-3FCF8E?style=for-the-badge&logo=supabase&logoColor=white)](https://supabase.com/)
[![Google Cloud](https://img.shields.io/badge/Google_Vertex_AI-Multimodal-4285F4?style=for-the-badge&logo=googlecloud&logoColor=white)](https://cloud.google.com/vertex-ai)

</div>

---

## Overview

The **Ultra-Low Latency E-Commerce Vector Search & Recommendation Gateway** is a high-throughput microservice architecture built to deliver sub-20ms hybrid vector and multimodal product search across high-volume e-commerce catalog datasets.

By coupling a concurrent Go HTTP API gateway with PostgreSQL `pgvector` HNSW (Hierarchical Navigable Small World) indexing and Google Vertex AI 768-dimensional multimodal embeddings, the gateway provides true semantic search and visual similarity lookup. The service is protected by automated SLA fallback logic, ensuring uninterrupted catalog availability during upstream AI service degraded states.

---

## Key Features

- **Sub-20ms Vector Hot Path:** Performs fast cosine similarity vector retrieval using Supabase PostgreSQL `pgvector` with HNSW indexing.
- **Multimodal Hybrid Search:** Embeds both raw natural language text queries (e.g., *"wireless setup for remote work"*) and uploaded binary image representations into unified 768-dimensional dense vectors via Google Vertex AI.
- **Resilient Sad Path Fallback:** Integrates automated circuit-breaking and error recovery to seamlessly downshift from vector similarity to traditional SQL trigram keyword search when upstream AI latency exceeds SLA thresholds.
- **Microsecond Go Gateway:** Zero-allocation memory design featuring lightweight Go goroutines and CORS middleware tuned for web client interactions.
- **Reactive Next.js Storefront:** Built with Next.js App Router, TypeScript, and Tailwind CSS v4, complete with client-side drag-and-drop visual image search, live search execution timing, and telemetry indicators.

---

## Architecture Diagram


```

+-----------------------------------------------------------------------------------+
|                                 Next.js Storefront                                |
|                        (React, TypeScript, Tailwind CSS v4)                        |
+-----------------------------------------------------------------------------------+
|
| HTTP POST /api/search
v
+-----------------------------------------------------------------------------------+
|                                 Go Microservice                                   |
|                             (CORS & Panic Recovery)                               |
+-----------------------------------------------------------------------------------+
/                                   

HOT PATH     /                                     \   SAD PATH
(Sub-20ms SLA)  /                                       \  (Fallback)
v                                         v
+----------------------------------+     +----------------------------------+
|        Google Vertex AI          |     |    PostgreSQL Trigram Search     |
| (768-dim Multimodal Embeddings)  |     |   (pg_trgm Fallback Matching)    |
+----------------------------------+     +----------------------------------+
|                                         |
v                                         |
+----------------------------------+                       |
|        Supabase pgvector         |                       |
|       (HNSW Index Cosine RPC)    |<----------------------'
+----------------------------------+

```

---

## Tech Stack

| Component | Technology | Description |
| :--- | :--- | :--- |
| **API Gateway** | **Go (1.22+)** | High-concurrency backend API handling JSON serialization and vector pipelines. |
| **Frontend UI** | **Next.js 15 / React / TS** | Reactive storefront app with client-side search UI and drag-and-drop upload. |
| **Styling Engine** | **Tailwind CSS v4** | Modern CSS engine utilizing direct `@import "tailwindcss";` compilation. |
| **Vector DB** | **Supabase (PostgreSQL)** | Persistent storage featuring `pgvector` and HNSW vector similarity indexes. |
| **AI Embeddings** | **Google Vertex AI** | Generates 768-dimensional multimodal embeddings for text and images. |

---

## Local Development & Setup

### Prerequisites
- [Go 1.22+](https://go.dev/doc/install)
- [Node.js 18+](https://nodejs.org/)
- [Supabase CLI](https://supabase.com/docs/guides/cli) or an active Supabase project instance

### 1. Repository Setup
```bash
git clone [https://github.com/yourusername/ecommerce-vector-gateway.git](https://github.com/yourusername/ecommerce-vector-gateway.git)
cd ecommerce-vector-gateway

```

### 2. Database Migration (Supabase pgvector)

Run the migration script against your PostgreSQL instance to enable `pgvector` and create the `match_products` stored procedure:

```bash
# Execute migration against your Supabase SQL Editor or CLI
psql -h <your-db-host> -U postgres -d postgres -f supabase/migrations/0001_init_vector_catalog.sql

```

### 3. Backend Service Setup

```bash
cd backend

# Install dependencies and sync modules
go mod tidy

# Start Go server
go run main.go

```

*The backend API will start on `http://localhost:8080`.*

### 4. Frontend Application Setup

In a separate terminal window:

```bash
cd frontend

# Install Node dependencies
npm install

# Run Next.js development server
npm run dev

```

*Access the storefront web application at `http://localhost:3000`.*

---

## API Reference

### Health Check

Verify API gateway availability and health metrics.

* **URL:** `/health`
* **Method:** `GET`
* **Response:**
```json
{
  "status": "ok",
  "timestamp": "2026-10-22T22:15:00Z"
}

```



### Vector Search Endpoint

Executes multimodal vector search or falls back to text search based on query inputs and SLA performance.

* **URL:** `/api/search`
* **Method:** `POST`
* **Headers:** `Content-Type: application/json`
* **Request Body:**
```json
{
  "query_text": "ergonomic mechanical keyboard",
  "image_base64": "data:image/png;base64,...",
  "limit": 12
}

```


* **Response Body:**
```json
{
  "products": [
    {
      "id": "e3a89e1c-5d91-4e78-9e5e-63f2b2f8a111",
      "title": "Ergonomic Mechanical Keyboard",
      "description": "Low-profile wireless mechanical keyboard with custom tactile switches.",
      "category": "Electronics",
      "price": 149.99,
      "image_url": "[https://images.unsplash.com/photo-1587829741301-dc798b83add3](https://images.unsplash.com/photo-1587829741301-dc798b83add3)",
      "in_stock": true,
      "similarity": 0.9412
    }
  ],
  "latency_ms": 14,
  "is_fallback": false
}

```



---

## License

Distributed under the [MIT License](https://www.google.com/search?q=LICENSE).
