package concsort

import (
	"reflect"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestSorting(t *testing.T) {

	arrayToSort := []int{3, 2, 4, 34, 5, 6, 2, 21}

	t.Log("Given the need to sort an array")
	{

		t.Logf("\tSorting array %v", arrayToSort)

		var sortedArray []int

		select {
		case sortedArray = <-Sort(arrayToSort):
		}

		// sortedArray := Sort(arrayToSort)
		expectedArray := []int{2, 2, 3, 4, 5, 6, 21, 34}

		if !reflect.DeepEqual(sortedArray, expectedArray) {
			t.Fatal("\t\tArray not sorted", ballotX, sortedArray)
		}

		t.Log("\t\tArray sorted", checkMark, sortedArray)
	}
}
