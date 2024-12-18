package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Order struct {
	ID    int
	Name  string
	Price int
}

var (
	order1 = Order{ID: 1, Name: "Laptop", Price: 70000}
	order2 = Order{}
)

func main() {
	url := flag.String("url", "localhost", "URL address")
	port := flag.String("port", "8080", "Port number")
	flag.Parse()

	http.HandleFunc("/getorder", getorder)
	http.HandleFunc("/sendorder", sendorder)
	fmt.Println("Server started.")

	server := &http.Server{
		Addr:         *url + ":" + *port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Ошибка сервера:", err)
	}
}

func getorder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(order1)
	fmt.Println("Обработан GET запрос.")
}

func sendorder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&order2)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Обработан POST запрос."))
	fmt.Printf("Обработан POST запрос. Создан заказ: %+v\n", order2)
}
