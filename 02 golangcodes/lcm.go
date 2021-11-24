package main

import "fmt"

func main() {
	var a, b, c, i int8
	fmt.Scanln(&a, &b)
	if a < b {
		c = b
	} else {
		c = a
	}
	i = 1
	for c != 0 {
		if (c*i)%a == 0 && (c*i)%b == 0 {
			fmt.Println((c * i), "is the LCM")
			break
		}

		i = i + 1
	}
}
