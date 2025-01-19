package piscine

func Capitalize(s string) string {
	arr := []rune(s)
	result := ""
	flag := false
	for i := 0; i < len(arr); i++ {
		if ((arr[i] >= 'a' && arr[i] <= 'z') || (arr[i] >= 'A' && arr[i] <= 'Z') || (arr[i] >= '0' && arr[i] <= '9')) && !flag {
			for j := 'a'; j <= 'z'; j++ {
				if j == arr[i] {
					arr[i] = (j - 'a') + 'A'
					break
				}
			}
			flag = true
		} else if (arr[i] >= 'a' && arr[i] <= 'z') || (arr[i] >= 'A' && arr[i] <= 'Z') || (arr[i] >= '0' && arr[i] <= '9') && flag {
			for j := 'A'; j <= 'Z'; j++ {
				if j == arr[i] {
					arr[i] = (j - 'A') + 'a'
					break
				}
			}
		} else {
			flag = false
		}
		result += string(arr[i])
	}
	return result
}
