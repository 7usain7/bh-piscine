package main

import (
	"fmt"
	"os"
	"strconv"
)

func validOp(n string) bool {
	return n == "+" || n == "-" || n == "/" || n == "*" || n == "%"
}

func main() {
	args := os.Args[1:]

	if len(args) != 3 || !validOp(args[1]) {
		return
	}
	a, err := strconv.Atoi(args[0])
	b, err1 := strconv.Atoi(args[2])
	if err != nil || err1 != nil {
		return
	}
	if (a > 0 && b > 0 && a+b < 0) || (a < 0 && b < 0 && a+b > 0) {
		return
	}
	if args[1] == "+" {
		fmt.Println(a + b)
	} else if args[1] == "-" {
		fmt.Println(a - b)
	} else if args[1] == "*" {
		fmt.Println(a * b)
	} else if args[1] == "/" {
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("No division by 0")
		}
	} else if args[1] == "%" {
		if b != 0 {
			fmt.Println(a % b)
		} else {
			fmt.Println("No modulo by 0")
		}
	}
}
