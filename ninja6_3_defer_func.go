package main

import (
	"fmt"
)

func main()  {
	foo2()
}

func foo2()  {
	defer defer_func(10)
	fmt.Println("In foo()")
}

func defer_func(x int)  {
	fmt.Println("defer func int: ", x)
}
