package sort

import (
	"fmt"
	"testing"

)

func TestPopSort(t *testing.T)  {
	fmt.Printf("%v", BubbleSort([]int{12, 10, 25, 7, 1, 9},2))
}