package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func isURLValid(u string) error {
	parsedURL, err := url.Parse(u)
	if err != nil || parsedURL.Scheme != "https" || parsedURL.Host == "" {
		return fmt.Errorf("неверный URL: %s", u)
	}

	allowedHost := "localhost"
	if parsedURL.Host == allowedHost {
		return nil
	}
	return fmt.Errorf("неверный HOST: %s", parsedURL.Host)
}

type Order struct {
	ID    int
	Name  string
	Price int
}

func main() {
	url := flag.String("url", "http://localhost:8080", "URL сервера.")
	getEndpoint := flag.String("get", "/getorder", "Эндпоинт для GET.")
	postEndpoint := flag.String("post", "/sendorder", "Эндпоинт для POST.")
	flag.Parse()

	getURL := *url + *getEndpoint
	postURL := *url + *postEndpoint

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	ctx := context.Background()

	if isURLValid(getURL) != nil {
		fmt.Println("Некорректный URL:", getURL)
		return
	}
	// GET запрос
	// #nosec G107
	req, err := http.NewRequestWithContext(ctx, "GET", getURL, nil)
	if err != nil {
		fmt.Printf("Ошибка Get запроса: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка Get запроса: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка ответа сервера: ", resp.Status)
		return
	}

	var newOrder Order

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения: ", err)
		return
	}

	err = json.Unmarshal(body, &newOrder)
	if err != nil {
		fmt.Println("Ошибка десериализации: ", err)
		return
	}
	fmt.Printf("Ответ от сервера (GET запрос): %+v\n", newOrder)

	// POST запрос
	newOrder2 := Order{ID: 2, Name: "Smartphone", Price: 45000}

	order2, err := json.Marshal(newOrder2)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}

	if isURLValid(postURL) != nil {
		fmt.Println("Некорректный URL:", postURL)
		return
	}
	// #nosec G107
	req, err = http.NewRequestWithContext(ctx, "POST", postURL, bytes.NewBuffer(order2))
	if err != nil {
		fmt.Println("Ошибка отправки POST запроса:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Ошибка отправки POST запроса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Ошибка ответа сервера:", resp.Status)
		return
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка получения запроса от сервера:", err)
		return
	}
	fmt.Println("Ответ от сервера (POST запрос):", string(body))
}
