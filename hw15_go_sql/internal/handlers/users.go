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

type UserHandler struct {
	db *pgxpool.Pool
}

func NewUserHandler(db *pgxpool.Pool) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	username := vars["username"]

	queries := repository.New(h.db)

	ctx := context.Background()
	user, err := queries.GetUserByUsername(ctx, username)
	if err != nil {
		log.Println("Error getting user:", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) InsertUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	request := repository.InsertUserParams{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	queries := repository.New(h.db)

	ctx := context.Background()
	err = queries.InsertUser(ctx, request)
	if err != nil {
		log.Println("Error inserting user:", err)
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func (h *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
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

	err = queries.DeleteUserByID(ctx, int32(idNum)) //nolint
	if err != nil {
		log.Printf("Failed to delete user with ID %v: %s", idNum, err)
		http.Error(w, "Failed to delete user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
