'use client';

import React, { useState, ChangeEvent, FormEvent } from 'react';

interface SearchBarProps {
  onSearch: (query: string, imageBase64?: string) => void;
  isLoading: boolean;
}

/**
 * SearchBar manages input state for semantic text queries and image drag-and-drop uploads.
 */
export const SearchBar: React.FC<SearchBarProps> = ({ onSearch, isLoading }) => {
  const [query, setQuery] = useState('');
  const [imagePreview, setImagePreview] = useState<string | undefined>();

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (query.trim() || imagePreview) {
      onSearch(query, imagePreview);
    }
  };

  const handleImageUpload = (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        const base64String = reader.result as string;
        setImagePreview(base64String);
      };
      reader.readAsDataURL(file);
    }
  };

  const clearImage = () => {
    setImagePreview(undefined);
  };

  return (
    <form onSubmit={handleSubmit} className="w-full max-w-3xl mx-auto mb-8">
      <div className="flex flex-col gap-3">
        <div className="relative flex items-center bg-slate-900 border border-slate-700 hover:border-indigo-500 rounded-2xl p-2 transition-all shadow-xl focus-within:ring-2 focus-within:ring-indigo-500/50">
          <input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="Search catalog semantically (e.g., 'wireless gear for remote work')..."
            className="w-full bg-transparent px-4 py-2 text-slate-100 placeholder-slate-400 focus:outline-none text-base"
          />

          <label className="cursor-pointer px-3 py-2 text-slate-400 hover:text-indigo-400 transition-colors">
            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 9a2 2 0 012-2h0.93a2 2 0 001.664-.89l0.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l0.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" />
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 13a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            <input type="file" accept="image/*" onChange={handleImageUpload} className="hidden" />
          </label>

          <button
            type="submit"
            disabled={isLoading}
            className="ml-2 px-6 py-2.5 bg-indigo-600 hover:bg-indigo-500 disabled:bg-indigo-900 text-white font-medium rounded-xl transition-colors flex items-center gap-2"
          >
            {isLoading ? (
              <span className="inline-block animate-spin rounded-full h-4 w-4 border-2 border-white border-t-transparent" />
            ) : (
              'Search'
            )}
          </button>
        </div>

        {imagePreview && (
          <div className="flex items-center gap-3 bg-slate-800/80 border border-slate-700 px-3 py-2 rounded-xl w-fit">
            <img src={imagePreview} alt="Search context preview" className="w-8 h-8 object-cover rounded-lg" />
            <span className="text-xs text-slate-300">Visual search filter active</span>
            <button type="button" onClick={clearImage} className="text-slate-400 hover:text-rose-400 text-xs font-semibold">
              Remove
            </button>
          </div>
        )}
      </div>
    </form>
  );
};