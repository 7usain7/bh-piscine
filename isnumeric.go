package piscine

func IsNumeric(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := '0'; j <= '9'; j++ {
			if arr[i] == j {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
