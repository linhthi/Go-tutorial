package main
import (
	"fmt"
)

// Buffered channels accept a limited number of values without a corresponding receiver for those values.

func main()  {
	messages := make(chan string, 2)

	messages <- "bufferd"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}