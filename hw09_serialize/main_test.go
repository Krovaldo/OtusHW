package main

import (
	"testing"

	"github.com/fixme_my_friend/hw09_serialize/pb"
	"github.com/stretchr/testify/assert"
)

func TestBookMarshalJSON(t *testing.T) {
	testBook := &Book{
		ID:     1,
		Title:  "TestBook",
		Author: "TestAuthor",
		Year:   2000,
		Size:   100,
		Rate:   5.0,
	}

	expectedJSON := `{"id":1,"title":"TestBook","author":"TestAuthor","year":2000,"size":100,"rate":5}`

	jsonRes, err := testBook.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(jsonRes))
}

func TestBookUnmarshalJSON(t *testing.T) {
	jsonData := `{"id":1,"title":"TestBook","author":"TestAuthor","year":2000,"size":100,"rate":5}`

	var book Book

	err := book.UnmarshalJSON([]byte(jsonData))
	assert.NoError(t, err)

	expectedBook := Book{
		ID:     1,
		Title:  "TestBook",
		Author: "TestAuthor",
		Year:   2000,
		Size:   100,
		Rate:   5.0,
	}

	assert.Equal(t, expectedBook, book)
}

func TestBookUnmarshalJSONError(t *testing.T) {
	jsonData := `{"id": "invalid", "title": "Test Book", "author": "Test Author", "year": 2023, "size": 100, "rate": 4.5}`

	var book Book
	err := book.UnmarshalJSON([]byte(jsonData))
	assert.Error(t, err)
}

func TestSerializationDeserializationBooks(t *testing.T) {
	books := []*pb.Book{
		{Id: 1, Title: "Book1", Author: "Author1", Year: 2000, Size: 100, Rate: 4.1},
		{Id: 2, Title: "Book2", Author: "Author2", Year: 2024, Size: 200, Rate: 4.2},
	}

	serialized, err := pb.SerializeBooks(books)
	assert.NoError(t, err)

	deserialized, err := pb.DeserializeBooks(serialized)
	assert.NoError(t, err)

	for i, book := range books {
		assert.Equal(t, book.Id, deserialized[i].Id)
		assert.Equal(t, book.Title, deserialized[i].Title)
		assert.Equal(t, book.Author, deserialized[i].Author)
		assert.Equal(t, book.Year, deserialized[i].Year)
		assert.Equal(t, book.Size, deserialized[i].Size)
		assert.Equal(t, book.Rate, deserialized[i].Rate)
	}
}
