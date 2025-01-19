package piscine

func ToLower(s string) string {
	arr := []rune(s)
	result := ""
	for i := 0; i < len(arr); i++ {
		for j := 'A'; j <= 'Z'; j++ {
			if j == arr[i] {
				arr[i] = (j - 'A') + 'a'
				break
			}
		}
		result += string(arr[i])
	}
	return result
}
