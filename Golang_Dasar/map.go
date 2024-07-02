package main

import "fmt"

func main() {
	person := map[string]string{
		"name" : "Timothy",
		"location": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["location"])

	person["title"] = "Programmer"

	fmt.Println(person["title"])
	fmt.Println("len of person: ", len(person))

	// Membuat Map Baru
	// person := make(map[string]string)

	// Isi Map
	// person["name"] = "Timothy"

	// Delete Map
	delete(person, "name")
	fmt.Println("Person after delete name : ", person)
}