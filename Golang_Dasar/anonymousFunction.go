package main

import (
	"fmt"
	"strings"
)

type Blacklist func(string) bool

func registeredUser(name string, blacklist Blacklist) {
	if (blacklist(name)) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	registeredUser("admin anjing", func(name string) bool {
		if(strings.Contains(name, "anjing")) {
			return true
		} else {
			return false
		}
	})
}