package piscine

func Pow(int1 int, int2 int) int {
	power := int1
	for i := 0; i < int2-1; i++ {
		power = power * int1
	}
	return power
}

func BasicAtoi(s string) int {
	integer := 0
	for i := 0; i < len(s); i++ {
		number := int(s[i] - '0')
		integer = integer*10 + number
	}
	return integer
}
