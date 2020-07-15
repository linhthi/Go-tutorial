package main
import (
	"fmt"
	"encoding/json"
)

type respone1 struct {
	Page int
	Fruits []string
}

type respone2 struct {
	Page int 	`json:"page`
	Fruits []string	`json:"fruits`
}

func main()  {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(int(intB))

}