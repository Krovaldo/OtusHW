package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrder(t *testing.T) {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", "/getorder", nil)
	if err != nil {
		t.Fatal(err)
	}

	respRec := httptest.NewRecorder()

	handler := http.HandlerFunc(getorder)
	handler.ServeHTTP(respRec, req)

	status := respRec.Code
	if status != http.StatusOK {
		t.Errorf("Неверный статус: got %v, want %v", status, http.StatusOK)
	}

	expectedOrder := order1
	receivedOrder := Order{}
	err = json.Unmarshal(respRec.Body.Bytes(), &receivedOrder)
	if err != nil {
		t.Fatal("Ошибка десериализации:", err)
	}

	if receivedOrder != expectedOrder {
		t.Errorf("Обработчик вернул неверный заказ: got %+v, want %+v", receivedOrder, expectedOrder)
	}
}

func TestSendOrder(t *testing.T) {
	ctx := context.Background()

	testOrder := Order{ID: 2, Name: "Smartphone", Price: 30000}
	orderJSON, err := json.Marshal(testOrder)
	if err != nil {
		t.Fatal("Ошибка сериализации заказа:", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "/sendorder", bytes.NewBuffer(orderJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	respRec := httptest.NewRecorder()

	handler := http.HandlerFunc(sendorder)
	handler.ServeHTTP(respRec, req)

	status := respRec.Code
	if status != http.StatusCreated {
		t.Errorf("Неверный статус: got %v, want %v", status, http.StatusCreated)
	}

	expectedResponse := "Обработан POST запрос."
	if respRec.Body.String() != expectedResponse {
		t.Errorf("Неверный ответ: got %v, want %v", respRec.Body.String(), expectedResponse)
	}

	if order2 != testOrder {
		t.Errorf("Обработчик не обновил заказ: got %v, want %v", order2, testOrder)
	}
}
