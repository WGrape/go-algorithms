package sort

import (
	"reflect"
	"testing"

)

func TestPopSort(t *testing.T)  {
	arr := []int{12, 10, 25, 7, 1, 9}
	res := BubbleSort(arr,2)
	descArr := []int{25,12, 10, 9,7,1}
	if !reflect.DeepEqual(res, descArr) {
		t.Fatal("降序失败")
	}
	res = BubbleSort(arr,1)
	ascArr := []int{1,7, 9, 10,12,25}
	if !reflect.DeepEqual(res,ascArr) {
		t.Fatal("升序失败")
	}
}