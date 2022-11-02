package main

import (
	"fmt"
)

func main() {
	var a float64 = 0.1;
	var b float64 = 0.2;
	fmt.Printf("%f", a + b );
	fmt.Printf("%f", .1 + .2);
	fmt.Println( a + b );
	fmt.Println( .1 + .2 );
}