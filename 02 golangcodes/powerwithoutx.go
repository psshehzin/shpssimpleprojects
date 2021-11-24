package main

import "fmt"

func main() {
	var a, b, c, i int32
	fmt.Scanln(&a, &b)
	c = a
	m := a
	i = b
	for i != 1 {
		j := c
		for m != 1 {
			c += j
			m = m - 1
		}
		m = a
		i = i - 1
	}
	fmt.Println(a, "raised to", b, "is", c)
}
