package piscine

func IsLower(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 'a'; j <= 'z'; j++ {
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
