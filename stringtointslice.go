package piscine

func StringToIntSlice(str string) []int {
	runes := []rune(str)
	arr := make([]int, len(runes))
	for i, val := range runes {
		arr[i] = int(val)
	}
	return arr
}
