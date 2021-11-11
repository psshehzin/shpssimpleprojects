package main

import "fmt"

func main() {
	var a, b, sum int8
	fmt.Scanln(&b, &a)
	fmt.Println("reverse fib sequence starting from", b, "and,", a, "is")
	fmt.Println(b)
	fmt.Println(a)
	sum = b

	for sum != 1 {
		sum = b - a
		b = a
		a = sum
		fmt.Println(sum)
	}

}
