package main

import "fmt"

func main() {
	var a, b, c int8
	fmt.Scanln(&a, &b)
	if a > b {
		c = b
	} else {
		c = a
	}
	for c != 0 {
		if a%c == 0 && b%c == 0 {
			fmt.Println(c, "is the GCD")
			break
		}
		c = c - 1
	}
}
