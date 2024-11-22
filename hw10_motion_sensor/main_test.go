package main

import (
	"testing"
	"time"
)

func TestReadSensor(t *testing.T) {
	dataChan := make(chan int64)

	go readSensorData(dataChan)

	count := 0
	timeout := time.After(10 * time.Second)

	for {
		select {
		case _, ok := <-dataChan:
			if !ok {
				if count == 0 {
					t.Error("Данные не были считаны")
				}
				return
			}
			count++
		case <-timeout:
			if count == 0 {
				t.Error("Данные не были считаны. Timeout")
			}
			return
		}
	}
}

func TestProcessData(t *testing.T) {
	dataChan := make(chan int64, 10)
	processedDataChan := make(chan float64, 1)

	for i := 0; i < 10; i++ {
		dataChan <- int64(i)
	}
	close(dataChan)

	go processData(dataChan, processedDataChan)

	avg := <-processedDataChan
	expectedAvg := 4.5
	if avg != expectedAvg {
		t.Errorf("Ожидаемое среднее арифметическое: %.2f, полученное: %.2f", expectedAvg, avg)
	}
}
