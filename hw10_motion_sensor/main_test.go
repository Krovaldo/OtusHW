package main

import (
	"sync"
	"testing"
	"time"
)

func TestReadSensor(t *testing.T) {
	dataChan := make(chan int64)
	var wg sync.WaitGroup

	defer close(dataChan)
	defer wg.Done()

	wg.Add(1)
	go readSensorData(dataChan, &wg)

	count := 0
	timeout := time.After(time.Minute)

	for {
		select {
		case <-dataChan:
			count++
		case <-timeout:
			wg.Wait()
			if count == 0 {
				t.Error("Данные не были считаны")
			}
			return
		}
	}
}

func TestProcessData(t *testing.T) {
	dataChan := make(chan int64, 10)
	processedDataChan := make(chan float64, 1)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		dataChan <- int64(i)
	}
	close(dataChan)

	wg.Add(1)
	go processData(dataChan, processedDataChan, &wg)

	wg.Wait()

	avg := <-processedDataChan
	expectedAvg := 4.5
	if avg != expectedAvg {
		t.Errorf("Ожидаемое среднее арифметическое: %.2f, полученное: %.2f", expectedAvg, avg)
	}
}
