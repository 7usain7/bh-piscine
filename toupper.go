package piscine

func ToUpper(s string) string {
	arr := []rune(s)
	result := ""
	for i := 0; i < len(arr); i++ {
		for j := 'a'; j <= 'z'; j++ {
			if j == arr[i] {
				arr[i] = (j - 'a') + 'A'
				break
			}
		}
		result += string(arr[i])
	}
	return result
}
