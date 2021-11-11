package main

import "fmt"

func main() {
	var a, b, c, i int32
	fmt.Scanln(&a, &b)
	c = a
	i = b
	for i != 1 {
		c = c * a
		i = i - 1
	}
	fmt.Println(a, "raised to", b, "is", c)
}
