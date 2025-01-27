package piscine

func ConcatParams(args []string) string {
	result := ""
	for _, val := range args {
		result += val + "\n"
	}
	return result
}
