package piscine

func StringToIntSlice(str string) []int {
	arr := make([]int, len(str))
	for i, val := range str {
		arr[i] = int(val)
	}
	return arr
}
