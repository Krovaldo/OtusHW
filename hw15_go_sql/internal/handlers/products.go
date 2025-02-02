package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Krovaldo/OtusHW/hw15_go_sql/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductHandler struct {
	db *pgxpool.Pool
}

func NewProductHandler(db *pgxpool.Pool) *ProductHandler {
	return &ProductHandler{db: db}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()

	queries := repository.New(h.db)

	products, err := queries.GetAllProducts(ctx)
	if err != nil {
		log.Println("Error getting products:", err)
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) InsertProduct(w http.ResponseWriter, r *http.Request) {
	var request repository.InsertProductParams

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	queries := repository.New(h.db)

	ctx := context.Background()

	err = queries.InsertProduct(ctx, request)
	if err != nil {
		log.Println("Error inserting product:", err)
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Product created successfully"})
}

func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Price pgtype.Numeric `json:"price"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	queries := repository.New(h.db)

	ctx := context.Background()

	err = queries.DeleteProduct(ctx, request.Price)
	if err != nil {
		log.Println("Error deleting product:", err)
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}
