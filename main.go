package main

import (
	"fmt"
	"math/rand"
	"runtime"

	"github.com/dsinecos/merge-sort-go/concsort"
	"github.com/dsinecos/merge-sort-go/seqsort"
)

func main() {
	fmt.Println("Number of Logical CPUs ", runtime.NumCPU())
	fmt.Println("GOMAXPROCS ", runtime.GOMAXPROCS(0))

	// arrToSort := []int{3, 2, 4, 34, 5, 6, 2, 21}
	arrToSort := []int{3, 2, 1}
	// arrToSort := createArray(10000000)
	// fmt.Println("Array to sort ", arrToSort)
	// seqsort.Sort(arrToSort)
	fmt.Println("Sequential sort", seqsort.Sort(arrToSort))

	select {
	case <-concsort.Sort(arrToSort):
		// fmt.Println("Concurrent sort", resultArray)
		fmt.Println("Maximum goroutines launched ", concsort.MaxGoroutine)
	}
}

func createArray(size int) []int {
	var array []int
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(10000000))
	}
	return array
}
