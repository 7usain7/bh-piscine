package piscine

func Sqrt(nb int) int {
	i := 1
	for i*i <= nb {
		if i*i == nb {
			return i
		}
		i++
	}
	return 0
}
