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

	m := map[string]person {
		p1.last : p1,
		p2.last : p2,
	}

	for _, v := range m {
		fmt.Println(v.first, v.last)
		for i, j := range v.flavour {
			fmt.Println(i, j)
		}
		fmt.Println("\n")
	}
}