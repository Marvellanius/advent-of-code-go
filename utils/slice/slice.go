package slice

func PrependIntSlice(slice []int, value int) []int {
	return append([]int{value}, slice...)
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Min(array []int) int {
	min, _ := MinMax(array)
	return min
}

func Max(array []int) int {
	_, max := MinMax(array)
	return max
}
