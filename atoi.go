package piscine

func Atoi(s string) int {
	integer := 0
	arr := []rune(s)
	for i := 0; i < len(s); i++ {
		ascii := int(arr[i])
		if !(ascii > '9' || ascii < '0') {
			number := ascii - '0'
			integer = integer*10 + number
		} else if arr[0] == '+' || arr[0] == '-' {
			continue
		} else {
			integer = 0
			break
		}
	}
	if arr[1] == '-' || arr[1] == '+' {
		return 0
	} else if arr[0] == '-' {
		return integer * -1
	} else {
		return integer
	}
}
