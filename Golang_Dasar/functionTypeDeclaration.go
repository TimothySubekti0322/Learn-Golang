package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(input string) string {
	return strings.ReplaceAll(input, "Anjing", "***")
}

func main() {
	sayHelloWithFilter("Anjing John", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing John", filter)
}