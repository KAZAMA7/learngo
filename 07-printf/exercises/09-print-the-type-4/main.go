package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the Type #4
//  Print the type and value of true using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of true is bool
// ---------------------------------------------------------

func main() {
	value := true
	fmt.Printf("Type of %v is %[1]T", value)
}
