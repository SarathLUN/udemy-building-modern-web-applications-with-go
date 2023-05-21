package main

import "fmt"

// Start : this package we will learn about pointer
func main() {
	myString := "Green"
	fmt.Println("Before calling function:", myString)
	// change variable whatWasSaid by using Pointer
	changeUsingPointer(&myString)
	fmt.Println("After calling function:", myString)
}

// create function that change value of a variable
func changeUsingPointer(s *string) {
	fmt.Println("s is set to:", s)
	newValue := "Red"
	*s = newValue
}
