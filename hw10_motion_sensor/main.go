package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func randomNumber(maxNum int) int64 {
	getRand, err := rand.Int(rand.Reader, big.NewInt(int64(maxNum)))
	if err != nil {
		fmt.Printf("Ошибка при генерации случайного числа: %v", err)
		return 0
	}

	return getRand.Int64()
}

func readSensorData(dataChan chan<- int64) {
	defer close(dataChan)

	timer := time.NewTimer(time.Minute)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return
		case dataChan <- randomNumber(1000):
			time.Sleep(70 * time.Millisecond)
		}
	}
}

func processData(dataChan <-chan int64, calculatedData chan<- float64) {
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

	go readSensorData(dataChan)

	go processData(dataChan, calculatedData)

	printCalculatedDate(calculatedData) // главная горутина main
}
