package main

import (
	"fmt"

	"github.com/Krovaldo/OtusHW/hw06_testing/hw04_struct_comparator/book"
	"github.com/Krovaldo/OtusHW/hw06_testing/hw04_struct_comparator/comparator"
)

func main() {
	new1 := book.Book{}
	new1.NewBook(50, "HarryPotter", "J. K. Rowling", 2010, 512, 9.2)
	fmt.Println(new1.Author())

	new2 := book.Book{}
	new2.NewBook(51, "The Lord of the Rings", "J. R. R. Tolkien", 2008, 1423, 8.9)
	fmt.Println(new2.Book())

	comp2book := comparator.NewComparator(comparator.CompareByRate)
	fmt.Println("Comprasion by year:", comp2book.Comprasion(&new1, &new2))
}
