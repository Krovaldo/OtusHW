package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID     int64
	Title  string
	Author string
	Year   int32
	Size   int64
	Rate   float32
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID     int64   `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   int32   `json:"year"`
		Size   int64   `json:"size"`
		Rate   float32 `json:"rate"`
	}{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	jsonBook := &struct {
		ID     int64   `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   int32   `json:"year"`
		Size   int64   `json:"size"`
		Rate   float32 `json:"rate"`
	}{}

	if err := json.Unmarshal(data, &jsonBook); err != nil {
		return err
	}

	b.ID = jsonBook.ID
	b.Title = jsonBook.Title
	b.Author = jsonBook.Author
	b.Year = jsonBook.Year
	b.Size = jsonBook.Size
	b.Rate = jsonBook.Rate

	return nil
}

func main() {
	bk := Book{12, "ewgsfds", "hello", 5234, 4322, 4.3}
	jsonBK, err1 := bk.MarshalJSON()

	if err1 != nil {
		fmt.Println(fmt.Errorf("ошибка сереализации %w", err1))
	}
	fmt.Println(string(jsonBK))

	jsonData := `{
		"id": 21,
		"title": "asdfqf",
		"author": "Hello",
		"year": 2020,
		"size": 123,
		"rate": 4.2
	}`

	var book Book

	err2 := json.Unmarshal([]byte(jsonData), &book)
	if err2 != nil {
		fmt.Println(fmt.Errorf("ошибка десериализации %w", err2))
	}
	fmt.Println(book)
}
