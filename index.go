package piscine

func Index(s string, toFind string) int {
	sentence := []rune(s)
	find := []rune(toFind)
	for i := 0; i <= len(sentence)-len(find); i++ {
		count := 0
		for j := 0; j < len(find); j++ {
			if sentence[i+j] != find[j] {
				break
			}
			count++
		}
		if count == len(find) {
			return i
		}
	}
	return -1
}
