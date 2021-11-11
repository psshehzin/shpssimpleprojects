package main

import "fmt"

func main() {
	var a, b, c int32
	fmt.Scanln(&a)
	b = 0
	for a != 0 {
		c = a % 10
		a = a / 10
		b = b*10 + c
	}
	fmt.Println("number reversed is", b)
}
