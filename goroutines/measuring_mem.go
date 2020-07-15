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

}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}