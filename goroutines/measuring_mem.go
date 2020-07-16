package main

import (
    "runtime"
    "fmt"
)

func test()  {
	fmt.Println("Hello")
}

func main() {
    // Below is an example of using our PrintMemUsage() function
    // Print our starting memory usage (should be around 0mb)
    PrintMemUsage()
    go test()
    PrintMemUsage()
    test()
    PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v KB", m.Alloc)
        fmt.Printf("\tTotalAlloc = %v KB\n", m.TotalAlloc)
        
}

