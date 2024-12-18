package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	counter++
	fmt.Printf("Worker %v done! Counter: %v\n", id, counter)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	val := 10
	wg.Add(val)
	for i := 0; i < val; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
}
