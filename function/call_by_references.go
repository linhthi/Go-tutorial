package main

import (
	"fmt"
)

// Call function by reference

func swap(x *int, y *int)  {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main()  {
	var a int = 1
	var b int = 2

	fmt.Println("Before swap: a, b: %d %d", a, b)

	swap(&a, &b)
	fmt.Println("After swap: a, b: %d %d", a, b)
	
}