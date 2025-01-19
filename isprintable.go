package piscine

func IsPrintable(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 0; j <= 32; j++ {
			if arr[i] == rune(j) {
				flag = true
			}
		}
		if flag {
			return false
		}
	}
	return true
}
