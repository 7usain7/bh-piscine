package main

import (
	"fmt"
	"os"
)

func isAscii(str string) bool {
	for _, e := range str {
		if e > 127 {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 3 {
		fmt.Println("Invalid number of arguements.")
		return
	}
	var sentence string
	var substr string
	var color string
	if len(args) == 3 {
		sentence = args[2]
		substr = args[1]
		color = args[0]
	} else if len(args) == 2 {
		sentence = args[1]
		color = args[0]
	} else {
		sentence = args[0]
	}
	if !isAscii(sentence) {
		fmt.Println("Only Ascii is accepted.")
		return
	}
	if sentence != "" {
		err := input(sentence, substr, color)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
