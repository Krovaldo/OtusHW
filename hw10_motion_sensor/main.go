package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func readSensorData(dataChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(dataChan)

	src := rand.NewSource(time.Now().UnixNano())
	getRand := rand.New(src)

	startTime := time.Now()

	for time.Since(startTime) < time.Minute {
		number := getRand.Intn(1000)
		dataChan <- number
		time.Sleep(70 * time.Millisecond)
	}
}

func processData(dataChan <-chan int, calculatedData chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(calculatedData)

	count := 0
	var sum float64

	for number := range dataChan {
		count++
		sum += float64(number)

		if count == 10 {
			avg := sum / 10
			calculatedData <- float64(avg)
			count = 0
			sum = 0
		}
	}
}

func printCalculatedDate(calculatedData <-chan float64) {
	for avg := range calculatedData {
		fmt.Printf("Среднее арифметическое: %v\n", avg)
	}
}

func main() {
	// генератор рандомных чисел
	// src := rand.NewSource(time.Now().UnixNano())
	// getRand := rand.New(src)
	// fmt.Println(getRand.Intn(10)) // рандомайзер от 0 до 10

	dataChan := make(chan int)
	calculatedData := make(chan float64)

	var wg sync.WaitGroup

	wg.Add(1)
	go readSensorData(dataChan, &wg)

	wg.Add(1)
	go processData(dataChan, calculatedData, &wg)

	printCalculatedDate(calculatedData) // главная горутина main
}
