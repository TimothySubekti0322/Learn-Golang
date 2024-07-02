package main

import "fmt"

func main() {
	type NoKTP string
	var ktpEko NoKTP = "123"

	fmt.Println(ktpEko)

	name := "Timothy"
	var e uint8 = name[0]
	var alphabet NoKTP = NoKTP(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(alphabet)
}