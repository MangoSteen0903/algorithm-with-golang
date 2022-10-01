package swap

func Swap_arr(arr *[]int, a, b int) {
	var actualArr = *arr
	actualArr[a], actualArr[b] = actualArr[b], actualArr[a]
}
