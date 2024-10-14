package comparator

import (
	"testing"

	"github.com/Krovaldo/OtusHW/hw06_testing/hw04_struct_comparator/book"
	"github.com/stretchr/testify/assert"
)

func TestComparator(t *testing.T) {
	book1 := &book.Book{}
	book1.NewBook(50, "HarryPotter", "J. K. Rowling", 2010, 512, 9.2)

	book2 := &book.Book{}
	book2.NewBook(51, "The Lord of the Rings", "J. R. R. Tolkien", 2008, 1423, 8.9)

	compByYear := NewComparator(CompareByYear)
	assert.True(t, compByYear.Comprasion(book1, book2))
	assert.False(t, compByYear.Comprasion(book2, book1))

	compBySize := NewComparator(CompareBySize)
	assert.True(t, compBySize.Comprasion(book2, book1))
	assert.False(t, compBySize.Comprasion(book1, book2))

	compByRate := NewComparator(CompareByRate)
	assert.True(t, compByRate.Comprasion(book1, book2))
	assert.False(t, compByRate.Comprasion(book2, book1))
}
