package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	number := nb
	for true {
		flag := true
		for i := 2; i < nb; i++ {
			if number%i == 0 && number != i {
				flag = false
			}
		}
		if flag {
			break
		}
		number++
	}
	return number
}
