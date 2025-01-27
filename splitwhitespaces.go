package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []string{}
	word := ""
	wordd := false
	for i, val := range s {
		if !(val == ' ' || val == '\t' || val == '\n') {
			word += string(val)
			wordd = true
			if i == len(s)-1 {
				arr = append(arr, word)
			}
		} else if (val == ' ' || val == '\t' || val == '\n') && wordd {
			arr = append(arr, word)
			word = ""
			wordd = false
		}
	}
	return arr
}
