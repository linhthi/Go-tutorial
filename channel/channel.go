package main
import (
	"fmt"
)

// Channels are the pipes that connect concurrent goroutines
// We can send values into channels from one goroutine and receive those values into another goroutine.
// Create a channel: chan val-type, ex: chan string
// Send a value into a channel using <-

func main()  {
	messages := make(chan string)

	go func ()  {messages <- "ping"} ()

	msg := <- messages
	fmt.Println(msg)
}