package main

import (
	"fmt"
)

const (
	a = iota //a = 0
	b = iota //b = 1
	c = iota //c = 2
)

const (
	d = iota * 2 //d = 0
	e = iota * 3 //e = 3
	f = iota * 4 //f = 8
)

func main() {
	fmt.Println("const a : ", a)
	fmt.Println("const b : ", b)
	fmt.Println("const c : ", c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println("const d : ", d)
	fmt.Println("const e : ", e)
	fmt.Println("const f : ", f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
