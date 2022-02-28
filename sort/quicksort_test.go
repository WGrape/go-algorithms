package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	fmt.Printf("%v", QuickSort([]int{12, 10, 25, 7, 1, 9}))
}
