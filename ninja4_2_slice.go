package main

import (
	"fmt"
)

func main() {
	slc := make([]int, 10, 100)
	for i := 0; i < 10; i++ {
		slc[i] = i + 1
	}

	for i, v := range slc {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slc)
}
