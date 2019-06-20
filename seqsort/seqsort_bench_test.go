package seqsort

import (
	"math/rand"
	"testing"
)

func BenchmarkSeqSort(b *testing.B) {

	b.ResetTimer()

	arrayToSort := createArray(1000000)

	for i := 0; i < b.N; i++ {
		Sort(arrayToSort)
	}
}

func createArray(size int) []int {
	var array []int
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(10000000))
	}
	return array
}