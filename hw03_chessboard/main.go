package main

import (
	"errors"
	"fmt"
)

func main() {
	var size int
	fmt.Scanln(&size)

	if size < 0 {
		err := errors.New("вы ввели отрицательное значение")
		fmt.Println(err)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
