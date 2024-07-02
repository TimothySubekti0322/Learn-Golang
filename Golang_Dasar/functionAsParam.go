package main

import (
	"fmt"
	"strings"
)

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(input string) string {
	return strings.ReplaceAll(input, "Anjing ", "")
}

func main() {
	sayHelloWithFilter("Anjing John", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing John", filter)
}