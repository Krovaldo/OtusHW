package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		num      int
		expected int
	}{
		{
			name:     "Число не найдено",
			arr:      []int{1, 2, 3, 4, 5},
			num:      100,
			expected: -1,
		},
		{
			name:     "Пустой срез",
			arr:      []int{},
			num:      1,
			expected: -1,
		},
		{
			name:     "Отсортированный массив",
			arr:      []int{1, 5, 9, 11, 12, 15},
			num:      12,
			expected: 4,
		},
		{
			name:     "Middle = num",
			arr:      []int{1, 1, 1, 1, 1},
			num:      1,
			expected: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			assert.Equal(t, tC.expected, BinarySearch(tC.arr, tC.num))
		})
	}
}
