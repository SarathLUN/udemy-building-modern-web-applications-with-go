package main

import "log"

func main() {
	//for i := 0; i <= 10; i++ {
	//	log.Println(i)
	//}
	//
	//animals := []string{"dog", "cat", "fish", "horse"}
	//for _, animal := range animals {
	//	log.Println(animal)
	//}
	//
	//moreAnimals := make(map[string]string)
	//moreAnimals["dog"] = "Fido"
	//moreAnimals["cat"] = "Fluffy"
	//for animalType, animalName := range moreAnimals {
	//	log.Println(animalType, animalName)
	//}

	// loop over string
	//firstLine := "Once upon the time"
	//for i, l := range firstLine {
	//	log.Println(i, ":", l)
	//}
	//log.Println("memory address before re-assigned", &firstLine)
	//firstLine = "Hello World!"
	//log.Println("memory address after re-assigned", &firstLine)

	// range over object
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var users []User
	users = append(users, User{"Jonh", "Smith", "jonh@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 23})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 40})
	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}
}
