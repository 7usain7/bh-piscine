package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := 0
	descending := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			descending++
		} else if f(a[i], a[i+1]) < 0 {
			ascending++
		}
	}
	return ascending == 0 || descending == 0
}
