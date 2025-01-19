package piscine

func AlphaCount(s string) int {
	arr := []rune(s)
	count := 0
	for i := 0; i < len(arr); i++ {
		if (arr[i] >= 'A' && arr[i] <= 'Z') || arr[i] >= 'a' && arr[i] <= 'z' {
			count++
		}
	}
	return count
}
