package main

import "fmt"

func main() {
	c := "b"
	i := 1

	ac := c[0]

	for i <= 26 {
		fmt.Println(string(ac))
		ac = ac + 1
		i = i + 1

	}
}
