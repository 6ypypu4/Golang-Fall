package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"practice5-sqlx/internal/repository"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// GetProductsHandler handles GET /products requests with filtering, sorting, and pagination
func GetProductsHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse query parameters
		filters := repository.ProductFilters{}

		if category := r.URL.Query().Get("category"); category != "" {
			filters.Category = category
		}

		if minPriceStr := r.URL.Query().Get("min_price"); minPriceStr != "" {
			if minPrice, err := strconv.Atoi(minPriceStr); err == nil {
				filters.MinPrice = &minPrice
			}
		}

		if maxPriceStr := r.URL.Query().Get("max_price"); maxPriceStr != "" {
			if maxPrice, err := strconv.Atoi(maxPriceStr); err == nil {
				filters.MaxPrice = &maxPrice
			}
		}

		if sort := r.URL.Query().Get("sort"); sort != "" {
			filters.Sort = sort
		}

		if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
			if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
				filters.Limit = limit
			}
		}

		if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
			if offset, err := strconv.Atoi(offsetStr); err == nil && offset > 0 {
				filters.Offset = offset
			}
		}

		// Query products
		products, err := repository.GetProductsWithFilters(db, filters)
		if err != nil {
			log.Printf("Error querying products: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Set response headers
		w.Header().Set("Content-Type", "application/json")

		// Send JSON response
		if err := json.NewEncoder(w).Encode(products); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
