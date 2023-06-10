package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the Type #3
//
//  Print the type and value of "hello" using fmt.Printf
//
// EXPECTED OUTPUT
// 	Type of hello is string
// ---------------------------------------------------------

func main() {
	hello := "hello"
	fmt.Printf("Type of %v is %[1]T\n", hello)
}
