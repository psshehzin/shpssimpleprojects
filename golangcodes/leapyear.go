package main

import "fmt"

func main() {
	var a int16
	fmt.Scanln(&a)
	if a%4 == 0 && a%100 == 0 {
		if a%400 == 0 {
			fmt.Println("Leap year it is")
		} else {
			fmt.Println("not a leapyear")
		}

	} else if a%4 == 0 {

		fmt.Println("Leap year it is")
	} else {
		fmt.Println("not a leapyear")
	}

}
