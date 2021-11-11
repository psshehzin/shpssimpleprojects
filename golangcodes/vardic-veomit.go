package main

import "fmt"

func omneg(num ...int) {
	for _, i := range num {
		if i >= 0 {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	omneg(1, 2, -3, 4, 5)
}
