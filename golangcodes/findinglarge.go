package main

import "fmt"

func main() {
	var a, b, c int16
	fmt.Scanln(&a, &b, &c)
	if a >= b && a >= c {
		fmt.Println(a, " islargest")
	} else if b >= a && b >= c {
		fmt.Println(b, " islargest")
	} else {
		fmt.Println(c, " islargest")
	}
}
