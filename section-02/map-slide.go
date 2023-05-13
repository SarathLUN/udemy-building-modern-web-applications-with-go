package main

import (
	"log"
	"sort"
)

type People struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "Hello"
	myMap["other-dog"] = "World!"
	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)
	myMap2["first"] = 1
	myMap2["second"] = 2
	log.Println(myMap2["first"], myMap2["second"])

	myMap3 := make(map[string]People)
	me := People{
		FirstName: "Tony",
		LastName:  "Stark",
	}
	myMap3["me"] = me
	log.Println(myMap3["me"].FirstName)

	var mySlice []string
	mySlice = append(mySlice, "Tony")
	mySlice = append(mySlice, "Stark")
	log.Println(mySlice)

	mySlice2 := []int{4, 2, 6, 1, 3, 5, 7, 10, 8}
	mySlice2 = append(mySlice2, 9)
	log.Println(mySlice2)
	sort.Ints(mySlice2)
	log.Println(mySlice2)
	// print first 3 value
	log.Println(mySlice2[:3])
	log.Println(mySlice2[3:5])

}
