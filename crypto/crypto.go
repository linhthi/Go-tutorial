package main

import (
    "crypto/rand"
    "fmt"
)

func main() {
    key := make([]byte, 64)

    _, err := rand.Read(key)
    fmt.Println("%v", key)
    if err != nil {
        // handle error here
    }
}