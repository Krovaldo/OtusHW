package main

import (
	"fmt"
	"strings"
)

func SizeOfGrid(size int) error {
	if size < 0 {
		return fmt.Errorf("вы ввели отрицательное значение")
	}
	return nil
}

func CreateGrid(size int) string {
	var f strings.Builder

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				f.WriteString("#")
			} else {
				f.WriteString(" ")
			}
		}
		f.WriteString("\n")
	}
	return f.String()
}

func main() {
	var size int
	fmt.Scanln(&size)

	if err := SizeOfGrid(size); err != nil {
		fmt.Println(err)
	}

	fmt.Print(CreateGrid(size))
}
