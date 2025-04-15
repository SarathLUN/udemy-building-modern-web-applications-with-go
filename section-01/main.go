package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("result: ", Add(1, 3))
}

func Add(a, b int32) int32 {
	return a + b
}
