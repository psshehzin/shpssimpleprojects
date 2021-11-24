package main

import "fmt"

func main() {
	var a, b int32
	fmt.Scanln(&a)
	b = 0
	for a != 0 {
		a = a / 10
		b = b + 1
	}
	fmt.Println("there are ", b, " digits ")
}
