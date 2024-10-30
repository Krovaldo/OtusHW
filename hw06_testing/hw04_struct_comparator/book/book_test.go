package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBook(t *testing.T) {
	book := Book{}
	book.NewBook(50, "HarryPotter", "J. K. Rowling", 2010, 512, 9.2)

	assert.Equal(t, 50, book.ID())
	assert.Equal(t, "HarryPotter", book.Title())
	assert.Equal(t, "J. K. Rowling", book.Author())
	assert.Equal(t, 2010, book.Year())
	assert.Equal(t, 512, book.Size())
	assert.Equal(t, float32(9.2), book.Rate())

	book.SetID(51)
	book.SetTitle("The Lord of the Rings")
	book.SetAuthor("J. R. R. Tolkien")
	book.SetYear(2008)
	book.SetSize(1423)
	book.SetRate(8.9)

	assert.Equal(t, 51, book.ID())
	assert.Equal(t, "The Lord of the Rings", book.Title())
	assert.Equal(t, "J. R. R. Tolkien", book.Author())
	assert.Equal(t, 2008, book.Year())
	assert.Equal(t, 1423, book.Size())
	assert.Equal(t, float32(8.9), book.Rate())
}
