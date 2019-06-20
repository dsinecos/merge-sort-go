package concsort

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkConcSort(b *testing.B) {

	b.ResetTimer()

	arrayToSort := createArray(1000000)

	var numGoroutine []int

	for i := 0; i < b.N; i++ {
		select {
		case <-Sort(arrayToSort):
			numGoroutine = append(numGoroutine, MaxGoroutine)
			// fmt.Println("\nMax Goroutines launched", MaxGoroutine)
		}
	}

	fmt.Println("\nAvg and max goroutines spawned")
	avg, max := stats(numGoroutine)
	fmt.Println("\t", avg, max)
}

func stats(array []int) (int, int) {

	var sum int
	var maxValue int

	for _, value := range array {
		sum += value

		if value > maxValue {
			maxValue = value
		}
	}

	avgValue := sum / len(array)

	return avgValue, maxValue
}

func createArray(size int) []int {
	var array []int
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(10000000))
	}
	return array
}
