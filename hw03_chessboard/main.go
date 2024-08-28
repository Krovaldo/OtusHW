package main

import "fmt"

func main() {
	var size int
	fmt.Scanln(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i%2 == 0 {
				fmt.Print("# ")
			} else {
				fmt.Print(" #")
			}

		}
		fmt.Print("\n")
	}
}
