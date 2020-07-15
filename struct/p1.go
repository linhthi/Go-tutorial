package main

import (
	"fmt"
	"time"
)

type emp struct {
	name string
	address string
	age int
}

// Function
func display(e emp) {
	fmt.Println(e.name)
}

// Method
func (e emp) display1()  {
	fmt.Println(e.name)
}

func calc(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}

func presentation(ch chan int)  {
	time.Sleep(5 * time.Second)
	fmt.Println("Inside display")
	ch <- 1234
}

func add_to_channel(ch chan int)  {
	fmt.Println("Send data")
	for i := 0; i < 10; i++ {
		ch <- i  //pushing data to channel
	}
	close(ch) // closing the channal	
}

func fecth_from_channel(ch chan int)  {
	fmt.Println("Read data")
	for {
		// fetch data
		x, flag := <- ch

		if flag == true {
			fmt.Println(x)
		} else {
			fmt.Println("Empty channel")
			break
		}
	}
}

func main() {
	x, y := 5, 10
	sum, diff := calc(x,y)
	fmt.Println("Sum", sum)
	fmt.Println("Diff", diff)

	var emp1data emp
	emp1data.name = "linh"
	emp1data.address ="Hanoi"
	emp1data.age = 21

	emp2data := emp{"Linh", "Paris", 21}

	display(emp1data)
	display(emp2data)
	emp1data.display1()


	// ch := make(chan int)
	// // Goroutine
	// go presentation(ch)
	// z := <- ch
	// fmt.Println("Inside main")
	// fmt.Println("Printing x in main() after taking from channel", z)

	ch := make(chan int)

	go add_to_channel(ch)
	go fecth_from_channel(ch)

	// deplay to prevent the existing of main() before gorountines finish
	time.Sleep(5 * time.Second)
	fmt.Println("Inside main()")

	
	

}
