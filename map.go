package piscine

func Map(f func(int) bool, a []int) []bool {
	arr := make([]bool, len(a))
	for i, val := range a {
		arr[i] = f(val)
	}
	return arr
}
