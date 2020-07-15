package main

import (
	"fmt"
	"os"
	"errors"
)

func openfile(name string)  {
	f, er := os.Open(name)

	if er != nil {
		return "", errors.New("Custom error message: File name is wrong")
	} else {
		return f.Name(), nil
	}
}

func main()  {
	filename, error := openfile("invalid.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("file opened", filename)
	}
}