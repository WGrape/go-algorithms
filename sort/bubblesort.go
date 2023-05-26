package sort

const(
	ASC = 1
	DESC = 2
)

func BubbleSort(arr []int, mod int) []int {
	l := len(arr)
	if l <=0 {
		return arr
	}
	for i:=0; i<l-1; i++ {

		if mod == DESC {
			for j := 0; j < l-1; j++ {
				if arr[j] < arr[j+1] {
					tmp := arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = tmp
				}
			}
		} else if mod == ASC {
			for j := 0; j < l-1; j++ {
				if arr[j] > arr[j+1] {
					tmp := arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = tmp
				}
			}
		}
	}
	return arr
}