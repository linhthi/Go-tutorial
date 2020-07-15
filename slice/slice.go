package main
import (
	"fmt"
)

// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

func main()  {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var a[]int = primes[1:4]
	fmt.Println(a)
}