package main

import "fmt"

func main() {
	// declare a string variable
	var whatToSay string
	// assign a value to the variable
	whatToSay = "Good Day World"
	// print the value of the variable
	println(whatToSay)

	// declare a int variable
	var i int
	// if we try to assign string to variable i, we will get an error
	// i = "Hello" // error: cannot use "Hello" (untyped string constant) as int value in assignment
	// assign a correct data type value to the variable i
	i = 7
	// print the value of the variable
	fmt.Println("i is set to", i)

}
