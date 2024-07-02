package main

import "fmt"

func getGoodBye(name string) string {
	return "goodbye " + name
}

func main() {
	newFunction := getGoodBye
	fmt.Println(newFunction("Timothy"))
}