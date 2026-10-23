'use client';

import React, { useState, useEffect } from 'react';
import { SearchBar } from '@/components/SearchBar';
import { ProductGrid, Product } from '@/components/ProductGrid';
import { searchProducts, SearchResponse } from '@/lib/api';

export default function StorefrontPage() {
  const [products, setProducts] = useState<Product[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [telemetry, setTelemetry] = useState<{ latency: number; fallback: boolean; notice?: string }>({
    latency: 0,
    fallback: false,
  });

  const handleSearch = async (queryText: string, imageBase64?: string) => {
    setIsLoading(true);
    const result: SearchResponse = await searchProducts(queryText, imageBase64);
    
    setProducts(result.products);
    setTelemetry({
      latency: result.latency_ms,
      fallback: result.is_fallback,
      notice: result.error_notice,
    });
    setIsLoading(false);
  };

  useEffect(() => {
    // Initial catalog load
    handleSearch('');
  }, []);

  return (
    <main className="min-h-screen bg-slate-950 text-slate-100 p-6 md:p-12">
      <header className="max-w-7xl mx-auto text-center mb-10">
        <h1 className="text-4xl md:text-5xl font-extrabold tracking-tight bg-gradient-to-r from-indigo-400 to-cyan-400 bg-clip-text text-transparent">
          Vector Storefront Gateway
        </h1>
        <p className="text-slate-400 mt-3 text-lg">
          Powered by Go microservices, Google Vertex AI multimodal embeddings & Supabase pgvector
        </p>
      </header>

      <section className="max-w-7xl mx-auto">
        <SearchBar onSearch={handleSearch} isLoading={isLoading} />

        {/* Telemetry and Latency Monitor Bar */}
        <div className="flex flex-wrap items-center justify-between bg-slate-900/60 border border-slate-800 rounded-xl px-4 py-2.5 mb-6 text-xs text-slate-400">
          <div className="flex items-center gap-4">
            <span>
              Latency: <strong className="text-indigo-400 font-mono">{telemetry.latency}ms</strong>
            </span>
            <span>
              Engine: <strong className="text-cyan-400">{telemetry.fallback ? 'SQL Keyword Fallback' : 'pgvector HNSW Search'}</strong>
            </span>
          </div>

          {telemetry.notice && (
            <div className="text-amber-400 font-medium bg-amber-950/40 px-2.5 py-1 rounded-md border border-amber-800/50">
              ⚠️ {telemetry.notice}
            </div>
          )}
        </div>

        {/* Results Grid */}
        <ProductGrid products={products} isLoading={isLoading} />
      </section>
    </main>
  );
}