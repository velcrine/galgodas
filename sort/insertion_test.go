package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{2, 3, 4, 21, 5, 6, 2, 2, 5}
	sortedArray := []int{2, 2, 2, 3, 4, 5, 5, 6, 21}

	InsertionSort(&arr)
	if !reflect.DeepEqual(arr, sortedArray) {
		t.Errorf("arr currently is %#v, but should be %#v after being sorted",
			arr, sortedArray)
	}
}
