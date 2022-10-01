package lotation

func Lotation_right(arr []int, s, t int) {
	temp := arr[s]
	for i := s; i < t; i++ {
		temp2 := arr[i+1]
		arr[i+1] = temp
		temp = temp2
	}
	arr[s] = temp
}

func Lotation_left(arr []int, s, t int) {
	temp := arr[s]
	for i := s; i > t; i-- {
		temp2 := arr[i-1]
		arr[i-1] = temp
		temp = temp2
	}
	arr[s] = temp
}

func Lotation_k(arr []int, s, t, k int) {

	temp := arr[t-(k-1) : t+1]
	copySlice := make([]int, (t-s+1)-k)
	temp2 := arr[s : t-k+1]
	copy(copySlice, temp2)

	temp1Index := 0
	temp2Index := 0
	for i := s; i < s+k; i++ {
		arr[i] = temp[temp1Index]
		temp1Index++
	}

	for i := s + k; i < t+1; i++ {
		arr[i] = copySlice[temp2Index]
		temp2Index++
	}
}
