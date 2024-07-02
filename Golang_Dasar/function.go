package main

import "fmt"

func sayHello(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func add(a int, b int) int {
	return a + b
}

// Function with multiple return value
func getFullName() (string, string) {
	return "Timothy", "Subekti"
}

// Function with named return value
func getFullName2() (firstname, lastname, nickname string, age int) {
	firstname = "Timothy"
	lastname = "Subekti"
	nickname = "Tim"
	age = 25
	return firstname, lastname, nickname, age
}

func main() {
	sayHello("Timothy", "Subekti")

	a := 10
	b := 20
	c := add(a, b)
	fmt.Println(a,"+",b,"=",c)

	firstname, lastname := getFullName()
	fmt.Println("Fullname:", firstname, lastname)

	newFirstName, _ := getFullName()
	fmt.Println("New Firstname:", newFirstName)

	firstname2, lastname2, nickname, age := getFullName2()
	fmt.Println("Fullname:", firstname2, lastname2, "Nickname:", nickname, "Age:", age)
}