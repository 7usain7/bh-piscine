package piscine

func StrRev(s string) string {
	rev := ""
	arr := []rune(s)
	for i := len(arr) - 1; i >= 0; i-- {
		rev += string(arr[i])
	}
	return rev
}
