package main

import "fmt"

func main() {
	var a, b int16
	fmt.Scanln(&a, &b)
	fmt.Println("quotient", a/b, "remainder", a%b)

}
