package main

import (
	"fmt"
)

func main() {
	unbuffer := make(chan string)
	go func() {
		fmt.Println(<-unbuffer)
	}()
	unbuffer <- "Hi there"
	// fmt.Println(<-unbuffer)
	// go func(){
	// 	fmt.Println(<-unbuffer)
	// }()
}