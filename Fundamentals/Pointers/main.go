package main

import "fmt"

func main() {
    s := "Hello, world!"

    p := &s
    fmt.Println(*p)

    p = new(string)
	
}
