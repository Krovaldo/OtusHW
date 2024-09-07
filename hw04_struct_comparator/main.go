package main

import "fmt"

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

func (c *Comparator) Comprasion(book1, book2 *Book) bool {
	switch c.compareField {
	case CompareByYear:
		return book1.year > book2.year
	case CompareBySize:
		return book1.size > book2.size
	case CompareByRate:
		return book1.rate > book2.rate
	default:
		return false
	}
}

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b *Book) SetNewBook(id int, title, author string, year, size int, rate float32) {
	b.id = id
	b.title = title
	b.author = author
	b.year = year
	b.size = size
	b.rate = rate
	fmt.Println("New book successfully added")
}

func (b *Book) GetBook() (id int, title, author string, year, size int, rate float32) {
	return b.id, b.title, b.author, b.year, b.size, b.rate
}

func main() {
	new1 := Book{}
	new1.SetNewBook(50, "HarryPotter", "J. K. Rowling", 2010, 512, 9.2)
	fmt.Println(new1.GetBook())

	new2 := Book{}
	new2.SetNewBook(51, "The Lord of the Rings", "J. R. R. Tolkien", 2008, 1423, 8.9)
	fmt.Println(new2.GetBook())

	comp2book := NewComparator(CompareByRate)
	fmt.Println("Comprasion by year:", comp2book.Comprasion(&new1, &new2))
}
