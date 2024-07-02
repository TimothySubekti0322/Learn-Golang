package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) hasMarried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko.Name) // Eko

	eko.hasMarried()

	fmt.Println(eko.Name) // Mr. Eko
}