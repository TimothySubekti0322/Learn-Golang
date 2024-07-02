/* An interface type is defined as a set of method
signatures. */

package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func Ups() interface{} {
	// return 1
	// return true
	return "Ups"
}

func main() {
	person := Person{
		"Timothy",
	}

	SayHello(person)

	// Interface kosong = Any
	var animal any
	animal = "Gajah"
	fmt.Println(animal)
	fmt.Println(Ups())
}