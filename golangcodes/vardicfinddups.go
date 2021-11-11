package main

import "fmt"

func fidup(num ...int) {
	m := 0
	arr := [5]int{}
	flag := 0
	for _, i := range num {
		k := 0
		for _, j := range num {
			if i == j {
				k = k + 1
			}

		}
		if k > 1 {
			for _, l := range arr {
				if i == l {
					flag = 1
					break
				}
			}
			if flag == 1 {
				flag = 0
				continue
			}
			fmt.Println(i, "has been passed", k, "times")
			arr[m] = i
			m = m + 1
		}
	}
}
func main() {
	fidup(1, 1, 2, 4, 2, 2, 2)
}
