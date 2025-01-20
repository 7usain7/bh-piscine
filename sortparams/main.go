package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args

	for i := 1; i < len(arr)-1; i++ {
		smallest := i
		flag := false
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallest] {
				smallest = j
				flag = true
			}
		}
		if flag {
			temp := arr[i]
			arr[i] = arr[smallest]
			arr[smallest] = temp
		}
	}
	for i := 1; i < len(arr); i++ {
		z01.PrintRune(rune(arr[i][0]))
		z01.PrintRune('\n')
	}
}
