package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		counter := 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				counter++
			}
		}
		if counter%2 == 1 {
			return a[i]
		}
	}
	return -1
}
