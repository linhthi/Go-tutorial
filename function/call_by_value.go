package main
import(
	"fmt"
)

// Call by value

func swap(x int, y int) int {
	var temp int

	temp = x
	x = y
	y = temp

	return temp
}

func main()  {
	var a int = 10
	var b int = 20

	fmt.Println("Before swap: a, b: %d %d", a, b)

	swap(a, b)
	fmt.Println("After swap: a, b: %d %d", a, b)

}