package main

import (
	"fmt"
)

func main() {
    var a int = 1
    defer p(a)
    a = 2
    p(a)
    defer p(a)
}

func p(a int){
    fmt.Println(a)
}