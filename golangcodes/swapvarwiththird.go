package main

import "fmt"

func main() {
	var a, b, c int16
	fmt.Scanln(&a, &b)
	fmt.Println("a was", a, "b was", b)
	c = a
	a = b
	b = c
	fmt.Println("now a is", a, "b is", b)

}
