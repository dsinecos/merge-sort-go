package main

import (
	"fmt"

	"github.com/dsinecos/merge-sort-go/concsort"
	"github.com/dsinecos/merge-sort-go/seqsort"
)

func main() {
	arrToSort := []int{3, 2, 4, 34, 5, 6, 2, 21}
	fmt.Println("Array to sort ", arrToSort)
	fmt.Println("Sequential sort", seqsort.Sort(arrToSort))

	select {
	case resultArray := <-concsort.Sort(arrToSort):
		fmt.Println("Concurrent sort", resultArray)
	}
}
