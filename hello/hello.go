package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)
	for i := 1; i <= t; i++ {
		var n, k int
		scanf("%d %d\n", &n, &k)
		printf("%d\n", n/k*k+min(n%k, k/2))
	}
}
