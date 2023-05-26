package sort

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 定位基准值(以最后一个元素为基准)
	//var (
	//	pivot    = arr[len(arr)-1] // 选择最后一个元素作为基准值
	//	leftArr  []int
	//	rightArr []int
	//)
	//for _, num := range arr[:len(arr)-1] {
	//	if num < pivot {
	//		leftArr = append(leftArr, num)
	//	} else {
	//		rightArr = append(rightArr, num)
	//	}
	//}

	// 定位基准值(以第一个元素为基准)
	var (
		pivot    = arr[0] // 选择第一个元素作为基准值
		leftArr  []int
		rightArr []int
	)
	for i := 1; i <= len(arr)-1; i++ {
		var num = arr[i]
		if num < pivot {
			leftArr = append(leftArr, num)
		} else {
			rightArr = append(rightArr, num)
		}
	}

	// 分治
	leftArr = quickSort(leftArr)
	rightArr = quickSort(rightArr)

	// 合并
	return append(append(leftArr, pivot), rightArr...)
}
