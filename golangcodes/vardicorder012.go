package main

import "fmt"

func order012(num ...int) {
	m := 0
	n := 0
	o := 0
	for _, i := range num {
		if i == 0 {
			m += 1
		} else if i == 1 {
			n += 1
		} else if i == 2 {
			o += 1
		}
	}
	for lop := 0; lop < m; lop++ {
		fmt.Print(0, " ")
	}
	for lop := 0; lop < n; lop++ {
		fmt.Print(1, " ")
	}
	for lop := 0; lop < o; lop++ {
		fmt.Print(2, " ")
	}
}
func main() {
	order012(0, 2, 1, 1, 2, 0)
}
