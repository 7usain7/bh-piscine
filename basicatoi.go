package piscine

func BasicAtoi(s string) int {
	integer := 0
	for i := 0; i < len(s); i++ {
		number := int(s[i] - '0')
		integer = integer*10 + number
	}
	return integer
}
