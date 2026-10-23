import React from 'react';

export interface Product {
  id: string;
  title: string;
  description: string;
  category: string;
  price: number;
  image_url: string;
  in_stock: boolean;
}

interface ProductGridProps {
  products: Product[];
  isLoading: boolean;
}

// ProductGrid renders responsive e-commerce cards with Tailwind CSS styling.
export const ProductGrid: React.FC<ProductGridProps> = ({ products, isLoading }) => {
  if (isLoading) {
    return (
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 p-4">
        {[...Array(8)].map((_, i) => (
          <div key={i} className="animate-pulse bg-slate-800 h-72 rounded-xl border border-slate-700" />
        ))}
      </div>
    );
  }

  if (products.length === 0) {
    return (
      <div className="text-center py-16 text-slate-400">
        <p className="text-lg">No products found for this search.</p>
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 p-4">
      {products.map((product) => (
        <div
          key={product.id}
          className="group bg-slate-900 border border-slate-800 hover:border-indigo-500/50 rounded-xl overflow-hidden transition-all duration-200 hover:shadow-lg hover:shadow-indigo-500/10 flex flex-col"
        >
          <div className="h-48 w-full bg-slate-800 relative overflow-hidden">
            <img
              src={product.image_url || '/placeholder-product.jpg'}
              alt={product.title}
              className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
            />
            <span className="absolute top-2 right-2 px-2.5 py-1 text-xs font-semibold bg-slate-950/80 backdrop-blur-md text-slate-300 rounded-full border border-slate-700">
              {product.category}
            </span>
          </div>
          <div className="p-4 flex flex-col flex-grow justify-between">
            <div>
              <h3 className="font-semibold text-slate-100 group-hover:text-indigo-400 transition-colors line-clamp-1">
                {product.title}
              </h3>
              <p className="text-sm text-slate-400 mt-1 line-clamp-2">{product.description}</p>
            </div>
            <div className="mt-4 flex items-center justify-between">
              <span className="text-lg font-bold text-emerald-400">${product.price.toFixed(2)}</span>
              <button className="px-3 py-1.5 text-xs font-medium bg-indigo-600 hover:bg-indigo-500 text-white rounded-lg transition-colors">
                Add to Cart
              </button>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
};