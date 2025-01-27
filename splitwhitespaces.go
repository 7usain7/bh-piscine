package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []string{}
	word := ""
	for i, val := range s {
		if !(val == ' ' || val == '\t' || val == 'n') {
			word += string(val)
			if i == len(s)-1 {
				arr = append(arr, word)
			}
		} else {
			arr = append(arr, word)
			word = ""
		}
	}
	return arr
}
