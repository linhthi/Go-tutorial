package main
import (
	"fmt"
)

func main()  {
	var a[2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11 }
	fmt.Println(primes)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}


}