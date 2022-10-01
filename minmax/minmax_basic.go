package minmax

func Max(a, b, c int) int {
	var max int
	switch {
	case a > b && a > c:
		max = a
	case b > a && b > c:
		max = b
	case c > a && c > b:
		max = c
	}
	return max
}

func Min(a, b, c int) int {
	var min int
	switch {
	case a < b && a < c:
		min = a
	case b < a && b < c:
		min = b
	case c < a && c < b:
		min = c
	}
	return min
}
