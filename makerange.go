package piscine

func MakeRange(min, max int) []int {
	if min > max || (max == 0 && min == 0) {
		return []int(nil)
	}
	arr := make([]int, max-min)
	pointer := 0
	for i := min; i < max; i++ {
		arr[pointer] = i
		pointer++
	}
	return arr
}
