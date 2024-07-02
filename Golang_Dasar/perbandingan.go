package main

import "fmt"

func main() {
	var name = "John Wick"
	var name2 = "John Wick"

	fmt.Println("name = ", name)
	fmt.Println("name2 = ", name2)

	fmt.Println("name = name2 ?", name == name2)
	fmt.Println("name[0] > name2[0] ?", name[0] > name2[0])
}