package main

import "fmt"

func main() {
	var size int
	fmt.Scanln(&size)

	for count := 0; count < size; count++ {
		for i := 0; i < size; i++ {
			fmt.Print("# ")
		}
		fmt.Print("\n")
	}
}
