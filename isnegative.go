package main

import "github.com/01-edu/z01 v0.1.0"

func IsNegative(nb int) {

	if nb < 0 {
		z01.PrintRune('T')
	} else{
		z01.PrintRune('F')
	}
}