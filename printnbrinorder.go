package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		str := ""
		for i := n; i > 0; i /= 10 {
			digit := i % 10
			str += string(rune(digit) + '0')
		}
		arr := []rune(str)
		for i := 0; i < len(arr); i++ {
			index := i
			flag := false
			for j := i + 1; j < len(arr); j++ {
				if arr[index] > arr[j] {
					index = j
					flag = true
				}
			}
			if flag {
				temp := arr[i]
				arr[i] = arr[index]
				arr[index] = temp
			}
		}
		for i := 0; i < len(arr); i++ {
			z01.PrintRune(arr[i])
		}
	}
}
