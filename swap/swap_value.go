package swap

func Swap_value(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
