package main

import "fmt"

func main() {
	var a, b, sum int8

	sum = 1
	fmt.Scanln(&a)
	for b = 1; b <= a; b++ {
		sum = sum * b
	}
	fmt.Println("factorial of ", a, " is ", sum)
}
