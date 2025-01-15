package piscine

func BasicAtoi2(s string) int {
	integer := 0
	arr := []rune(s)
	for i := 0; i < len(s); i++ {
		ascii := int(arr[i])
		if !(ascii > '9' || ascii < '0'){
			number := ascii - '0'
			integer = integer*10 + number
		} else{
			integer = 0
			break
		}
		
	}
	return integer
}