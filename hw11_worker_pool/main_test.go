package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	testCases := []struct {
		name       string
		numWorkers int
		expected   int
	}{
		{
			name:       "5 workers",
			numWorkers: 5,
			expected:   5,
		},
		{
			name:       "1 workers",
			numWorkers: 1,
			expected:   1,
		},
		{
			name:       "10 workers",
			numWorkers: 10,
			expected:   10,
		},
		{
			name:       "0 workers",
			numWorkers: 0,
			expected:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			counter = 0
			var wg sync.WaitGroup

			wg.Add(tC.numWorkers)
			for i := 0; i < tC.numWorkers; i++ {
				go worker(i, &wg)
			}
			wg.Wait()

			assert.Equal(t, tC.expected, counter)
		})
	}
}
