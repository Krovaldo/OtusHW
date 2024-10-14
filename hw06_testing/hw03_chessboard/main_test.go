package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	testCases := []struct {
		name string
		size int
		err  error
	}{
		{
			name: "пять",
			size: 5,
			err:  nil,
		},
		{
			name: "ноль",
			size: 0,
			err:  nil,
		},
		{
			name: "минус один",
			size: -1,
			err:  fmt.Errorf("вы ввели отрицательное значение"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			assert.Equal(t, tC.err, SizeOfGrid(tC.size))
		})
	}
}

func TestCreateGrid(t *testing.T) {
	testCases := []struct {
		name     string
		size     int
		expected string
	}{
		{
			name:     "размер 0",
			size:     0,
			expected: "",
		},
		{
			name:     "размер 1",
			size:     1,
			expected: "#\n",
		},
		{
			name:     "размер 2",
			size:     2,
			expected: "# \n #\n",
		},
		{
			name:     "размер 3",
			size:     3,
			expected: "# #\n # \n# #\n",
		},
		{
			name:     "размер 4",
			size:     4,
			expected: "# # \n # #\n# # \n # #\n",
		},
		{
			name:     "размер 5",
			size:     5,
			expected: "# # #\n # # \n# # #\n # # \n# # #\n",
		},
		{
			name:     "размер 6",
			size:     6,
			expected: "# # # \n # # #\n# # # \n # # #\n# # # \n # # #\n",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			assert.Equal(t, tC.expected, CreateGrid(tC.size))
		})
	}
}
