import { Product } from '@/components/ProductGrid';

export interface SearchResponse {
  products: Product[];
  latency_ms: number;
  is_fallback: boolean;
  error_notice?: string;
}

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

/**
 * Sends text or image query payload to the Go backend search handler.
 * HOT PATH: Primary semantic/visual vector search query execution.
 * SAD PATH: Captures fallback metadata when vector search drops back to keyword search.
 */
export async function searchProducts(
  queryText: string,
  imageBase64?: string,
  limit = 12
): Promise<SearchResponse> {
  try {
    const response = await fetch(`${API_BASE_URL}/api/search`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        query_text: queryText,
        image_base64: imageBase64,
        limit,
      }),
    });

    if (!response.ok) {
      throw new Error(`API response error: ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    // SAD PATH: Graceful UI recovery if the Go backend service is unreachable
    console.error('Search API execution failed:', error);
    return {
      products: [],
      latency_ms: 0,
      is_fallback: true,
      error_notice: 'Unable to connect to search gateway server.',
    };
  }
}