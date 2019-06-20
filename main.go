package main

import (
	"fmt"
	"math"
)

func main() {
	arrToSort := []int{3, 2, 4, 34, 5, 6, 2, 21}
	fmt.Println(sort(arrToSort))
}

func sort(array []int) []int {
	lenArr := len(array)

	if lenArr == 1 {
		return array
	}

	sliceAt := math.Floor(float64(lenArr / 2))

	arr1 := make([]int, int(sliceAt))
	copy(arr1, array[:int(sliceAt)])

	arr2 := make([]int, lenArr-int(sliceAt))
	copy(arr2, array[int(sliceAt):])

	return merge(sort(arr1), sort(arr2))
}

func merge(arr1 []int, arr2 []int) []int {
	size := len(arr1) + len(arr2)
	var workArr []int

	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] <= arr2[0] {
			workArr = append(workArr, arr1[0])
			arr1 = arr1[1:]
		} else if arr1[0] > arr2[0] {
			workArr = append(workArr, arr2[0])
			arr2 = arr2[1:]
		}
	}

	if len(arr1) > 0 {
		workArr = append(workArr, arr1...)
	}

	if len(arr2) > 0 {
		workArr = append(workArr, arr2...)
	}

	resultArr := make([]int, size)
	copy(resultArr, workArr)

	return resultArr
}
