package piscine

func IsUpper(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 'A'; j <= 'Z'; j++ {
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
