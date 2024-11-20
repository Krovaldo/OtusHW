package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

func randomNumber(max int) int64 {
	getRand, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		fmt.Printf("Ошибка при генерации случайного числа: %v", err)
		return 0
	}

	return getRand.Int64()
}

func readSensorData(dataChan chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(dataChan)

	startTime := time.Now()
	for time.Since(startTime) < time.Minute {
		dataChan <- randomNumber(1000)
		time.Sleep(70 * time.Millisecond)
	}
}

func processData(dataChan <-chan int64, calculatedData chan<- float64, wg *sync.WaitGroup) {
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
	// генератор рандомных чисел math/rand (линтеру не нравится :( )
	// src := rand.NewSource(time.Now().UnixNano())
	// getRand := rand.New(src)
	// fmt.Println(getRand.Intn(10)) // рандомайзер от 0 до 10

	dataChan := make(chan int64)
	calculatedData := make(chan float64)

	var wg sync.WaitGroup

	wg.Add(1)
	go readSensorData(dataChan, &wg)

	wg.Add(1)
	go processData(dataChan, calculatedData, &wg)

	printCalculatedDate(calculatedData) // главная горутина main
}
