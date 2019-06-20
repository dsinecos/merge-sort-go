package main

import (
	"fmt"

	"github.com/dsinecos/merge-sort-go/seqsort"
)

func main() {
	arrToSort := []int{3, 2, 4, 34, 5, 6, 2, 21}
	fmt.Println(seqsort.Sort(arrToSort))
}
