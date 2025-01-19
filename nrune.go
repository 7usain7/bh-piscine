package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	if n > len(arr) || n < 1 {
		return 0
	}
	return arr[n-1]
}
