package main

import "fmt"

func main() {
	var n, sum, i, a, b int8
	a = 1
	b = 1
	fmt.Scanln(&n)
	fmt.Println("first ", n, "numbers in fib sequence is")
	if n == 1 {
		fmt.Println(a)
	} else {
		fmt.Println(a)
		fmt.Println(a)
	}

	for i = 3; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
		fmt.Println(sum)
	}

}
