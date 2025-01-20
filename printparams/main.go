package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args
	for i := 1; i < len(arr); i++ {
		word := []rune(arr[i])
		for j := 0; j < len(word); j++ {
			z01.PrintRune(word[j])
		}
		z01.PrintRune('\n')
	}
}
