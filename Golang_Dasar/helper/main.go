package helper

import "fmt"

var version = 1 // Private
var Application = "Golang"

// Must be Uppercase
func Main() {
	fmt.Println(SayHello("Timothy"))
	fmt.Println(SayHi("Timothy"))
}