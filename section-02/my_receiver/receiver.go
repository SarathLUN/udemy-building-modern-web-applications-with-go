package my_receiver

import "log"

type myStruct struct {
	FirstName string
}

func (m myStruct) printFirstName() string {
	return m.FirstName
}

// Start : this package we will learn about receiver in Go.
func Start() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to:", myVar.FirstName)
	log.Println("myVar with receiver:", myVar.printFirstName())
	log.Println("myVar2 is set to:", myVar2.FirstName)
	log.Println("myVar2 with receiver:", myVar2.printFirstName())

}
