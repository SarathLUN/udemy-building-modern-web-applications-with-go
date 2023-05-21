package main

import "fmt"

// declare a package level variable
var myInt int

// Start : in this package we will learn about variable and function.
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

	// declare a variable by using ':=' to store return from function saySomething
	whatWasSaid := saySomething()
	// print the value of the variable
	fmt.Println("The function returned", whatWasSaid)

	// use variable whatWasSaid and myInt to store return from function saySomethingElse
	whatWasSaid, myInt = saySomethingElse()
	// print the value of the variable
	fmt.Println("The function returned", whatWasSaid, myInt)

}

// create function saySomething that returns a string "Something"
func saySomething() string {
	return "Something"
}

// create function saySomethingElse that returns a string "Something Else" and an int 7
func saySomethingElse() (string, int) {
	return "Something Else", 7
}
