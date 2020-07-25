package main

import "fmt"

type person2 struct {
	first string
	last  string
	age   int
}

func (p person2) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p1 := person2{
		first: "Anshul",
		last:  "Kishore",
		age:   24,
	}

	p1.speak()
}
