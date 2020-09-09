package main

import (
	"fmt"
)

func main() {
	str := "Hello World"
	rChar := reverse(str)
	fmt.Println(rChar)
}

// reverse function
func reverse(str string) (result string) { 
    for _, i := range str { 
        result = string(i) + result 
    } 
    return
}