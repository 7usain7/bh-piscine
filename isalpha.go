package piscine

func IsAlpha(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 'A'; j <= 'Z'; j++ {
			if arr[i] == j {
				flag = true
			}
		}
		for j := 'a'; j <= 'z'; j++ {
			if arr[i] == j {
				flag = true
			}
		}
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
