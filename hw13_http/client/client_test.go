package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupTestServer() *httptest.Server {
	handler := http.NewServeMux()

	handler.HandleFunc("/getorder", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		order := Order{ID: 1, Name: "Laptop", Price: 70000}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	})

	handler.HandleFunc("/sendorder", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		receivedOrder := Order{}
		err := json.NewDecoder(r.Body).Decode(&receivedOrder)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Заказ успешно создан"))
	})

	return httptest.NewServer(handler)
}

func TestGetOrder(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	ctx := context.Background()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", server.URL+"/getorder", nil)
	if err != nil {
		t.Fatal("Ошибка при создании GET запроса:", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Ошибка при выполнении GET запроса:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Неверный статус: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	receivedOrder := Order{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Ошибка чтения тела ответа:", err)
	}

	err = json.Unmarshal(body, &receivedOrder)
	if err != nil {
		t.Fatal("Ошибка десериализации:", err)
	}

	expectedOrder := Order{ID: 1, Name: "Laptop", Price: 70000}
	if receivedOrder != expectedOrder {
		t.Errorf("got: %v, want: %v", receivedOrder, expectedOrder)
	}
}

func TestSendOrder(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	ctx := context.Background()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	newOrder := Order{ID: 2, Name: "Smartphone", Price: 45000}
	orderJSON, err := json.Marshal(newOrder)
	if err != nil {
		t.Fatal("Ошибка сериализации:", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", server.URL+"/sendorder", bytes.NewBuffer(orderJSON))
	if err != nil {
		t.Fatal("Ошибка создания POST запроса:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Ошибка выполнения POST запроса:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Неверный статус: got %v, want %v", resp.StatusCode, http.StatusCreated)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Ошибка чтения тела:", err)
	}

	expectedResponse := "Заказ успешно создан"
	if string(body) != expectedResponse {
		t.Errorf("got: %v, want: %v", string(body), expectedResponse)
	}
}
