package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		sum := 1
		for i := 0; i < power; i++ {
			sum *= nb
		}
		return sum
	}
}
