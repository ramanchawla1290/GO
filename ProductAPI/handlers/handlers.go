package handlers

import (
	"context"
	"log"
	"net/http"

	data "github.com/ramanchawla1290/GO/ProductAPI/data"
)

// Handler Struct for Products

type ProductHandler struct {
	logger *log.Logger
}

func NewProductHandler(logger *log.Logger) *ProductHandler {
	return &ProductHandler{logger}
}

// MIDDLEWARE

// Middleware for Logging : ALL HTTP METHODS

func (ph *ProductHandler) MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ph.logger.Printf("[%s] %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

// Middleware for Product Validation : PUT & POST Only

func (ph *ProductHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			ph.logger.Println("ERROR deserializing product json")
			http.Error(w, "ERROR Invalid JSON format", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "prod", prod)
		req := r.WithContext(ctx)

		err = prod.Validate() // Using GO VALIDATOR
		if err != nil {
			ph.logger.Println("Product Validation ERROR")
			http.Error(w, "Product Validation ERROR:\n"+err.Error(), http.StatusUnprocessableEntity)
			return
		}

		next.ServeHTTP(w, req)

		if prod.ID >= 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			err = prod.ToJSON(w)
			if err != nil {
				ph.logger.Println("ERROR serializing product json")
				http.Error(w, "ERROR sending json response", http.StatusInternalServerError)
				return
			}

		}
	})
}
