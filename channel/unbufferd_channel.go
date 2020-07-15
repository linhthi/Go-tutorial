package main

import(
	"fmt"
	"time"
)

func test(ch chan int)  {
	fmt.Println("Func goroutine begins sending data")
	ch <- 1
	fmt.Println("Func goroutine ends sending data")
}

func main()  {
	ch := make(chan int)

	go test(ch)

	fmt.Println("Main goroutines sleep 2 second")
	time.Sleep(2 * time.Second)
	
	fmt.Println("Main goroutines begins receiving data")
	d := <- ch
	fmt.Println("Main gourtines received data", d)
	time.Sleep(time.Second)
	
}