package concsort

import (
	"math"
	"runtime"
)

// MaxGoroutine Captures the maximum number of goroutines in a single run
var MaxGoroutine int

// Sort To concurrently merge sort a Slice of int
func Sort(array []int) <-chan []int {
	return split(array)
}

func split(array []int) <-chan []int {
	outChan := make(chan []int)

	go func() {
		defer close(outChan)

		if runtime.NumGoroutine() > MaxGoroutine {
			MaxGoroutine = runtime.NumGoroutine()
		}

		// fmt.Println("Splitting array ", array)
		// fmt.Println("Number of goroutines running ", runtime.NumGoroutine())

		lenArr := len(array)
		if lenArr == 1 {
			outChan <- array
			return
		}

		sliceAt := math.Floor(float64(lenArr / 2))

		arr1 := make([]int, int(sliceAt))
		copy(arr1, array[:int(sliceAt)])

		arr2 := make([]int, lenArr-int(sliceAt))
		copy(arr2, array[int(sliceAt):])

		mergedValue := <-merge(split(arr1), split(arr2))
		outChan <- mergedValue

		// outChan <- <-merge(split(chanArr1), split(chanArr2))
	}()

	return outChan
}

func merge(chanArr1 <-chan []int, chanArr2 <-chan []int) <-chan []int {
	outChan := make(chan []int)

	arr1 := <-chanArr1
	arr2 := <-chanArr2

	// fmt.Println("Merging arrays ", arr1, arr2)
	// fmt.Println("Number of goroutines running ", runtime.NumGoroutine())

	go func() {
		defer close(outChan)

		if runtime.NumGoroutine() > MaxGoroutine {
			MaxGoroutine = runtime.NumGoroutine()
		}

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
		outChan <- resultArr
	}()

	return outChan
}
