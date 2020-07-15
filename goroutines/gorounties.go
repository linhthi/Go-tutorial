package main
import(
	"fmt"
	"time"
)

func longLastingProcess(c chan string)  {
	time.Sleep(2 * time.Second)
	c <- "linh"
}

func main()  {
	c := make(chan string)
	
}