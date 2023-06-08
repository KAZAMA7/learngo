package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Months #2
//
//  1. Initialize multiple constants using iota.
//  2. Please follow the instructions inside the code.
//
// EXPECTED OUTPUT
//  1 2 3
// ---------------------------------------------------------

func main() {
	const _ = iota
	// Now, use iota and initialize the following constants
	// "automatically" to 1, 2, and 3 respectively.
	const (
		Jan = 1 + iota
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)
}
