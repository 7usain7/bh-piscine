package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	path := os.Args[0]
	arr := []rune(path)
	index := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == '\\' {
			index = i
			break
		}
	}
	for i := index + 1; i < len(arr)-4; i++ {
		z01.PrintRune(arr[i])
	}
}
