package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	number := nb
	for {
		flag := true
		for i := 2; i < nb; i++ {
			if number%i == 0 && number != i {
				flag = false
				break
			}
		}
		if flag {
			break
		}
		number++
	}
	return number
}
