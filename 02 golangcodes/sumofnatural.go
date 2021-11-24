package main

import "fmt"

func main() {
	var a, b, sum int8

	sum = 0
	fmt.Scanln(&a)
	for b = 1; b <= a; b++ {
		sum = sum + b
	}
	fmt.Println("sum of ", a, " numebrs is ", sum)
}
