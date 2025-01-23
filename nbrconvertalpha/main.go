package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args
	isUpper := false
	if arr[1] == "--upper" {
		isUpper = true
	}
	for i := 1; i < len(arr); i++ {
		if i == 1 && isUpper {
			continue
		}
		integer := 0
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] < '0' || arr[i][j] > '9' {
				integer = 0
				break
			}
			number := int(rune(arr[i][j]) - '0')
			integer = integer*10 + number
		}
		if (integer > 0 && integer <= 26) && !isUpper {
			z01.PrintRune(rune('a' + integer - 1))
		} else if (integer > 0 && integer <= 26) && isUpper {
			z01.PrintRune(rune('A' + integer - 1))
		} else {
			z01.PrintRune('-')
		}
	}
	z01.PrintRune('\n')
}
