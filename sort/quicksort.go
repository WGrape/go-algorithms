package sort

func QuickSort(arr []int) []int {
	newArr := make([]int, len(arr))
	for i, v := range arr {
		newArr[i] = v
	}
	recursiveSort(newArr, 0, len(newArr)-1)
	return newArr
}

func recursiveSort(arr []int, start, end int) {
	if end <= start {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			if splitIndex != i {
				temp := arr[splitIndex]
				arr[splitIndex] = arr[i]
				arr[i] = temp
			}
			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	recursiveSort(arr, start, splitIndex-1)
	recursiveSort(arr, splitIndex+1, end)
}
