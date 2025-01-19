package piscine

func Compare(a, b string) int {
	arrA := []rune(a)
	arrB := []rune(b)
	length := len(arrA)
	if length > len(arrB) {
		length = len(arrB)
	}
	for i := 0; i < length; i++ {
		if arrA[i] > arrB[i] {
			return 1
		} else if arrA[i] < arrB[i] {
			return -1
		}
	}
	if len(arrA) > len(arrB) {
		return 1
	} else if len(arrA) < len(arrB) {
		return -1
	} else {
		return 0
	}
}
