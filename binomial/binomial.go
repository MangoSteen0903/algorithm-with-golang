package binomial

func Binomial(n, m int) int {
	var result [200][200]int
	if result[n][m] > 0 {
		return result[n][m]
	}
	if m == 0 || n == m {
		result[n][m] = 1
		return result[n][m]
	}
	result[n][m] = Binomial(n-1, m-1) + Binomial(n-1, m)
	return result[n][m]
}
