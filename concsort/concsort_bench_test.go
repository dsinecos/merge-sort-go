package concsort

import (
	"math/rand"
	"testing"
)

func BenchmarkConcSort(b *testing.B) {

	b.ResetTimer()

	arrayToSort := createArray(1000000)

	for i := 0; i < b.N; i++ {
		select {
		case <-Sort(arrayToSort):
		}
	}
}

func createArray(size int) []int {
	var array []int
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(10000000))
	}
	return array
}