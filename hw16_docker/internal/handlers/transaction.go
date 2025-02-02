package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Krovaldo/OtusHW/hw16_docker/internal/repository"
	"github.com/jackc/pgx/v5"
)

func (h *UserHandler) GetAllUsersAndProducts(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()

	tx, err := h.db.Begin(ctx)
	if err != nil {
		log.Println("Error starting transaction:", err)
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	var txCommited bool
	defer func(tx pgx.Tx, ctx context.Context) {
		if !txCommited {
			err = tx.Rollback(ctx)
			if err != nil {
				log.Printf("failed to rollback transaction: %s", err)
			}
		}
	}(tx, ctx)

	queriesWithTx := repository.New(h.db).WithTx(tx)

	users, err := queriesWithTx.GetAllUsers(ctx)
	if err != nil {
		log.Printf("failed to get all users: %s", err)
		http.Error(w, "failed to get all users", http.StatusInternalServerError)
		return
	}

	products, err := queriesWithTx.GetAllProducts(ctx)
	if err != nil {
		log.Printf("failed to get all products: %s", err)
		http.Error(w, "failed to get all products", http.StatusInternalServerError)
		return
	}

	if err = tx.Commit(ctx); err != nil {
		log.Println("Error committing transaction:", err)
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	txCommited = true

	response := struct {
		Users    []*repository.GetAllUsersRow `json:"users"`
		Products []*repository.Product        `json:"products"`
	}{
		Users:    users,
		Products: products,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("failed to encode response: %s", err)
		return
	}
}
