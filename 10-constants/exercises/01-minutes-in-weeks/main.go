package main

import "fmt"

func main() {
	const (
		minsPerDay = 24 * 60
		weekDays   = 7
	)
	fmt.Println("The number of minutes in two weeks are : ", 2*weekDays*minsPerDay)
}
