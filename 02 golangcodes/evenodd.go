package main

import "fmt"

func main() {
	var a int16
	fmt.Scanln(&a)
	if a%2 == 0 {
		fmt.Println(" i am even")

	} else {

		fmt.Println(" odd it is ")
	}

}
