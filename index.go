package piscine

func Index(s string, toFind string) int {
	sentence := []rune(s)
	find := []rune(toFind)
	flag := false
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == find[0] {
			flag = true
			for j := 1; j < len(find); j++ {
				flag = false
				if sentence[i+j] == find[j] {
					flag = true
				}
			}
		}
		if flag {
			return i
		}
	}
	return -1
}
