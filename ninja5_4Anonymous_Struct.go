package main

import(
	"fmt"
)

func main() {

	p1 := struct{
		first string
		last string
	}{
		first : "Anshul",
		last : "Kishore",
	}

	fmt.Println(p1.first, p1.last)
}