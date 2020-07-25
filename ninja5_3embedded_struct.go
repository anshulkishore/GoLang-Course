package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	s1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	s2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Silver",
		},
		luxury: true,
	}

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1.vehicle.doors)
	fmt.Println(s2.doors)
}
