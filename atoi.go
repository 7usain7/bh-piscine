package piscine

func Atoi(s string) int {
	integer := 0
	arr := []rune(s)
	if len(arr) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		ascii := int(arr[i])
		if !(ascii > '9' || ascii < '0') {
			number := ascii - '0'
			integer = integer*10 + number
		} else if i == 0 && (arr[i] == '+' || arr[i] == '-') {
			continue
		} else {
			integer = 0
			break
		}
	}
	if len(arr) == 1 {
		return integer
	}
	if arr[0] == '-' {
		return integer * -1
	} else {
		return integer
	}
}
