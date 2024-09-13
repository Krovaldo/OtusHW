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

func (b *Book) NewBook(id int, title, author string, year, size int, rate float32) {
	b.id = id
	b.title = title
	b.author = author
	b.year = year
	b.size = size
	b.rate = rate
}

func (b Book) Book() interface{} {
	return b
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func (b Book) ID() (id int) {
	return b.id
}

func (b Book) Title() (title string) {
	return b.title
}

func (b Book) Author() (author string) {
	return b.author
}

func (b Book) Year() (year int) {
	return b.year
}

func (b Book) Size() (size int) {
	return b.size
}

func (b Book) Rate() (rate float32) {
	return b.rate
}

func main() {
	new1 := Book{}
	new1.NewBook(50, "HarryPotter", "J. K. Rowling", 2010, 512, 9.2)
	fmt.Println(new1.Author())

	new2 := Book{}
	new2.NewBook(51, "The Lord of the Rings", "J. R. R. Tolkien", 2008, 1423, 8.9)
	fmt.Println(new2.Book())

	comp2book := NewComparator(CompareByRate)
	fmt.Println("Comprasion by year:", comp2book.Comprasion(&new1, &new2))
}
