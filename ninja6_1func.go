package main

import (
	"fmt"
)

func main() {
	a := foo()
	fmt.Println("foo int: ", a)

	b, str := bar()
	fmt.Println("bar int: ", b, ", bar string: ", str)
}

func foo() int {
	return 10;
}

func bar() (int, string) {
	return 20, "Anshul"
}