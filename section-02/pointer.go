package main

import "fmt"


func main() {
  myString := "Green"
  fmt.Println("Before calling function:",myString)
  // change variable whatWasSaid by using Pointer
  changeUsingPointer(&myString)
  fmt.Println("After calling function:",myString)
}

// create function that change value of a variable
func changeUsingPointer(s *string)  {
  newValue := "Red" 
  *s = newValue
}
