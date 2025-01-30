package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := 0
	for i, val := range a {
		if val > a[max] {
			max = i
		}
	}
	return a[max]
}
