package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Area
//
//  Fix the following program.
//
// RESTRICTION
//  You should not use any variables.
//
// EXPECTED OUTPUT
//  area = 1250
// ---------------------------------------------------------

func main() {
	const (
		width  int16 = 25
		height int32 = int32(width) * 2
	)

	fmt.Printf("area = %d\n", int32(width)*height)
}
