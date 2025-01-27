package piscine

func ConcatParams(args []string) string {
	result := ""
	for index, val := range args {
		result += val
		if index != len(args)-1 {
			result += "\n"
		}
	}
	return result
}
