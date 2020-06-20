package main

import(
	"fmt"
)

type person struct {
	first string
	last string
	flavour []string
}

func main() {

	p1 := person{
		first : "Anshul",
		last : "Kishore",
		flavour: []string{
			"choco",
			"vanilla",
		},
	}

	p2 := person{
		first : "Komal",
		last : "Sethia",
		flavour: []string{
			"choco",
			"black current",
			"hazlenut",
		},
	}

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.flavour {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.flavour {
		fmt.Println(i, v)
	}

}