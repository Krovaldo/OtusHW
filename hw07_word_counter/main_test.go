package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformStr(t *testing.T) {
	testCases := []struct {
		name     string
		str      string
		expected []string
	}{
		{
			name:     "пустая строка",
			str:      "",
			expected: []string{},
		},
		{
			name:     "Привет, мир!",
			str:      "Привет, мир!",
			expected: []string{"привет", "мир"},
		},
		{
			name:     "Пунктуация",
			str:      "Тест.... На, пунктуацию!?",
			expected: []string{"тест", "на", "пунктуацию"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			assert.Equal(t, tC.expected, TransformStr(tC.str))
		})
	}
}

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name     string
		str      []string
		expected map[string]int
	}{
		{
			name:     "Пустая строка",
			str:      []string{},
			expected: map[string]int{},
		},
		{
			name:     "привет мир",
			str:      []string{"привет", "мир"},
			expected: map[string]int{"привет": 1, "мир": 1},
		},
		{
			name:     "несколько одинаковых слов",
			str:      []string{"привет", "мир", "мир"},
			expected: map[string]int{"привет": 1, "мир": 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			assert.Equal(t, tC.expected, CountWords(tC.str))
		})
	}
}

func TestIntegration(t *testing.T) {
	testCases := []struct {
		name     string
		str      string
		expected map[string]int
	}{
		{
			name:     "пустая строка",
			str:      "",
			expected: map[string]int{},
		},
		{
			name:     "Привет, мир!",
			str:      "Привет, мир!",
			expected: map[string]int{"привет": 1, "мир": 1},
		},
		{
			name:     "Пунктуация",
			str:      "Тест.... На, пунктуацию!?",
			expected: map[string]int{"тест": 1, "на": 1, "пунктуацию": 1},
		},
		{
			name:     "несколько одинаковых слов",
			str:      "Привет, мир, мир!!!",
			expected: map[string]int{"привет": 1, "мир": 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			assert.Equal(t, tC.expected, CountWords(TransformStr(tC.str)))
		})
	}
}
