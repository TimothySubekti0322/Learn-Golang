package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Timothy"
	names[1] = "Doe"
	names[2] = "John"

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	names[2] = "Jane"
	fmt.Println(names[2])

	fmt.Println("Deklarasi Array Langsung")

	var fruits = [3]string{"Apple", "Banana", "Mango"}

	fmt.Println(fruits)

	var numbers = [...]int{2, 3, 4, 5, 6, 7}

	fmt.Println(numbers)
	fmt.Println(len(numbers))
}
