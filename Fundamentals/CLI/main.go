package main

import (
    "bufio"
    "os"
    "strings"
	"fmt"
)

func main() {
	fmt.Println("What would you like me to scream?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	
	
	fmt.Println(s + "!")
}