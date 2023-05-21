package main

import "log"

// Start : this package we will learn about if and switch-case statement
func main() {
	var isTrue bool
	isTrue = true
	if isTrue {
		log.Println("isTrue is set to:", isTrue)
	} else {
		log.Println("isTrue is set to:", isTrue)
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("cat is set to:", cat)
	} else {
		log.Println("cat is set to:", cat)
	}

	myNum := 100
	isTrue = false
	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99, and isTrue is set to true.")
	} else if myNum < 100 && isTrue {
		log.Println("1")
	} else if myNum == 101 || isTrue {
		log.Println("2")
	} else if myNum > 1000 && isTrue {
		log.Println("3")
	}

	myVar := "cat"
	switch myVar {
	case "cat":
		log.Println("myVar is set to cat.")
	case "dog":
		log.Println("myVar is set to dog.")
	default:
		log.Println("myVar is set to something else.")
	}
}
