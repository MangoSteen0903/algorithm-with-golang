package minmax

func Max_arr(numList []int, length int) int {
	var max int
	for i := 0; i < length; i++ {
		if numList[i] > max {
			max = numList[i]
		}
	}
	return max
}
