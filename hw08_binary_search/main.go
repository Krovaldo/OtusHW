package main

import (
	"fmt"
	"sort"
)

func BinarySearch(arr []int, num int) int {
	// Сортируем массив
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (left + right) / 2
		if arr[middle] == num {
			return middle
		}
		if num < arr[middle] {
			right = middle - 1
		}
		if num > arr[middle] {
			left = middle + 1
		}
	}
	return -1 // Если число не найдено
}

func main() {
	arr := []int{1, 19, 8, 100, 9, 12, 15}
	fmt.Println(BinarySearch(arr, 8))
	fmt.Println(arr)
}
