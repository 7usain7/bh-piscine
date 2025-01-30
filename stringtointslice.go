package piscine

func StringToIntSlice(str string) []int {
	if len(str) == 0 {
		return nil
	}
	runes := []rune(str)
	arr := make([]int, len(runes))
	for i, val := range runes {
		arr[i] = int(val)
	}
	return arr
}
