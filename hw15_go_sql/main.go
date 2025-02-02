package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Krovaldo/OtusHW/hw15_go_sql/internal/handlers"
	"github.com/Krovaldo/OtusHW/hw15_go_sql/pkg/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env") // Загружаем переменные окружения
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	pool, err := db.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	fmt.Println("Database successfully connected!")

	userHandler := handlers.NewUserHandler(pool)
	orderHandler := handlers.NewOrderHandler(pool)
	productHandler := handlers.NewProductHandler(pool)

	r := mux.NewRouter()
	r.HandleFunc("/user/create", userHandler.InsertUser).Methods("POST")
	r.HandleFunc("/user/{username}", userHandler.GetUserByUsername).Methods("GET")
	r.HandleFunc("/users_and_products", userHandler.GetAllUsersAndProducts).Methods("GET")
	r.HandleFunc("/user/delete/{id}", userHandler.DeleteUserByID).Methods("POST")

	r.HandleFunc("/order/create", orderHandler.InsertProduct).Methods("POST")
	r.HandleFunc("/order/delete/{id}", orderHandler.DeleteOrder).Methods("POST")
	r.HandleFunc("/order/{userID}", orderHandler.GetOrdersByUser).Methods("GET")

	r.HandleFunc("/product", productHandler.GetAllProducts).Methods("GET")
	r.HandleFunc("/product/create", productHandler.InsertProduct).Methods("POST")
	r.HandleFunc("/product/delete", productHandler.DeleteProductHandler).Methods("POST")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Printf("Server failed: %v", err)
		return
	}
}
