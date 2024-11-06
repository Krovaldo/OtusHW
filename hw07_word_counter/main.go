package main

import (
	"fmt"
	"strings"
	"unicode"
)

func TransformStr(str string) []string {
	var woPunct []rune

	for _, v := range str {
		if !unicode.IsPunct(v) {
			woPunct = append(woPunct, v)
		}
	}

	toLower := string(woPunct)
	toLower = strings.ToLower(toLower)
	result := strings.Fields(toLower)

	return result
}

func CountWords(str []string) map[string]int {
	wordCount := make(map[string]int)

	for _, v := range str {
		wordCount[v]++
	}
	return wordCount
}

func main() {
	str := "Привет, мир мир. мир. мир!"

	fmt.Print(CountWords(TransformStr(str)))
}
