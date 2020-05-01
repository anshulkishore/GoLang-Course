package main

import "fmt"

//zero values example
var i1 int
var s1 string
var b1 bool

var i2 = 10
var s2 = "Anshul Kishore"
var b2 = true

type anshul int
var x anshul
var y int

func main() {
	fmt.Println(i1)
	fmt.Println(s1)
	fmt.Println(b1)

	s := fmt.Sprintf("%v\t%v\t%v\n", i2, s2, b2)
	fmt.Println(s)

	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 100
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)

}