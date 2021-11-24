package main

import "fmt"

func main() {
	var a, b int16
	fmt.Scanln(&a, &b)
	fmt.Println("a was", a, "b was", b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("now a is", a, "b is", b)

}
