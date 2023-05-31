package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
		[
			{
				"first_name":"Clerk",
				"last_name":"Kent",
				"hair_color":"black",
				"has_dog":true
},
			{
				"first_name":"Bruce",
				"last_name":"Wayne",
				"hair_color":"black",
				"has_dog":false
}
]
`

	var unMarshalled []Person

	// read json into struct
	err := json.Unmarshal([]byte(myJson), &unMarshalled)
	if err != nil {
		log.Fatalln("fail unmarshal json")
	}
	log.Printf("unmarshalled: %v", unMarshalled)

	// write struct to json
	var mySlice []Person
	m1 := Person{
		"test1",
		"head1",
		"brown",
		true,
	}
	m2 := Person{
		"test2",
		"head2",
		"white",
		false,
	}
	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)
	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Fatalln("fail to marshal:", err)
	}
	log.Println(string(newJson))
}
