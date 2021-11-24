package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("testfile")
	defer f.Close()
	const buffsize = 5
	b1 := make([]byte, buffsize)
	str := "shsha"
	ou := str[1]
	fmt.Printf("%T,%T", str, ou)

	for {
		a, err := f.Read(b1)
		if err == io.EOF {
			break
		}

		fmt.Println("Content: ", string(b1[:a]))
	}

}
