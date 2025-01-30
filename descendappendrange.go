package piscine

func DescendAppendRange(max, min int) []int {
	arr := []int{}
	if min > max || (max == 0 && min == 0) {
		return []int(nil)
	}
	for i := max; i > min; i-- {
		arr = append(arr, i)
	}
	return arr
}
