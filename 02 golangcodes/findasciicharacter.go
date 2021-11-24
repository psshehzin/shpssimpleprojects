package main

import "fmt"

func main() {
	var c string
	fmt.Scanln(&c)

	ac := []rune(c)[0]

	fmt.Println("ascii value of", c, "is", ac)
	fmt.Println(string(ac))

}
