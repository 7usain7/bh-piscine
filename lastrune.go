package piscine

func FirstRune(s string) rune {
	arr := []rune(s)
	return arr[len(arr)-1]
}