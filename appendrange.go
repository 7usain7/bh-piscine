package piscine

func AppendRange(min, max int) []int {
	arr := []int{}
	if min > max || (max == 0 && min == 0) {
		return []int(nil)
	}
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
