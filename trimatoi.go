package piscine

func TrimAtoi(s string) int {
	integer := 0
	flag := false
	negativeFlag := false
	arr := []rune(s)
	if len(arr) == 0 {
		return 0
	}
	for i := 0; i < len(arr); i++ {
		ascii := int(arr[i])
		if ascii == '-' && !flag {
			negativeFlag = true
		} else if !(ascii > '9' || ascii < '0') {
			number := ascii - '0'
			integer = integer*10 + number
			flag = true
		}
	}
	if negativeFlag {
		return integer * -1
	}
	return integer
}
