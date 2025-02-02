package handlers

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/Krovaldo/OtusHW/hw15_go_sql/internal/repository"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderHandler struct {
	db *pgxpool.Pool
}

func NewOrderHandler(db *pgxpool.Pool) *OrderHandler {
	return &OrderHandler{db: db}
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)
	id := vars["id"]

	idNum, err := strconv.Atoi(id)

	if idNum < math.MinInt32 || idNum > math.MaxInt32 {
		log.Printf("ID out of range for int32: %d", idNum)
		return
	}

	if err != nil {
		log.Printf("Failed to convert ID to int: %s", err)
		http.Error(w, "Failed to convert ID:", http.StatusInternalServerError)
		return
	}

	queries := repository.New(h.db)

	err = queries.DeleteOrder(ctx, int32(idNum)) //nolint
	if err != nil {
		log.Printf("Failed to delete order %v: %s", idNum, err)
		http.Error(w, "Failed to delete order", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order deleted successfully"))
}

func (h *OrderHandler) InsertProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	request := repository.InsertOrderParams{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	queries := repository.New(h.db)

	err = queries.InsertOrder(ctx, request)
	if err != nil {
		log.Printf("Failed to insert order: %s", err)
		http.Error(w, "Failed to insert order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order created successfully"))
}

func (h *OrderHandler) GetOrdersByUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)
	userIDStr := vars["userID"]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println("Invalid userID:", err)
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	userIDInt32 := int32(userID) //nolint

	queries := repository.New(h.db)

	orders, err := queries.GetOrdersByUser(ctx, userIDInt32)
	if err != nil {
		log.Println("Error getting orders:", err)
		http.Error(w, "Failed to get orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
