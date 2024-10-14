package comparator

import (
	"github.com/Krovaldo/OtusHW/hw06_testing/hw04_struct_comparator/book"
)

type CompareField int

const (
	CompareByYear CompareField = iota
	CompareBySize
	CompareByRate
)

type Comparator struct {
	compareField CompareField
}

func NewComparator(cF CompareField) *Comparator {
	return &Comparator{compareField: cF}
}

func (c *Comparator) Comprasion(book1, book2 *book.Book) bool {
	switch c.compareField {
	case CompareByYear:
		return book1.Year() > book2.Year()
	case CompareBySize:
		return book1.Size() > book2.Size()
	case CompareByRate:
		return book1.Rate() > book2.Rate()
	default:
		return false
	}
}
