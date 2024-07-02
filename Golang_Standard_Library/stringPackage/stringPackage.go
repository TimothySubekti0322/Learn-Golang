package stringPackage

import (
	"fmt"
	"strings"
)

func MainString() {
	
	// Trim - menghilangkan karakter tertentu di awal dan akhir string
	fmt.Print("Trim sHelloWorlds : ")
	fmt.Println(strings.Trim("sHelloWorlds", "s"))

	// Split
	fmt.Print("Split Hello World : ")
	fmt.Println(strings.Split("Hello World", " "))

	// To Lower
	fmt.Print("ToLower HELLO WORLD : ")
	fmt.Println(strings.ToLower("HELLO WORLD"))

	// To Upper
	fmt.Print("ToUpper hello world : ")
	fmt.Println(strings.ToUpper("hello world"))

	// Contains
	fmt.Print(" 'Hello World' contains 'World' ? : ")
	fmt.Println(strings.Contains("Hello World", "World"))

	// Replace All
	fmt.Print("Replace All World become golang : ")
	fmt.Println(strings.ReplaceAll("Hello World", "World", "Golang"))
}