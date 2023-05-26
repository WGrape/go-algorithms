package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var (
		arr       = []int{8, 5, 1, 2, 10, 5, 3, 9, 4, 1, 7, 6}
		sortedArr = quickSort(arr)
	)
	if !reflect.DeepEqual(sortedArr, []int{1, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10}) {
		fmt.Println(sortedArr)
		t.Fail()
		return
	}
	fmt.Println(sortedArr)
}
