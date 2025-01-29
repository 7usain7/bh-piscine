package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := map[string]int{}
	word := ""
	wordSlice := []string{}
	for _, val := range str {
		if val != ' ' {
			word += string(val)
		} else if word != "" {
			wordSlice = append(wordSlice, word)
			word = ""
		}
	}
	if word != "" {
		wordSlice = append(wordSlice, word)
	}
	for _, val := range wordSlice {
		_, found := result[val]
		if found {
			result[val] += 1
		} else {
			result[val] = 1
		}
	}
	return result
}
