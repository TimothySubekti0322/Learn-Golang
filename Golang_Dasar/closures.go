package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++

		// counter2 := 100

		// increment2 := func() {
		// 	fmt.Println("Increment2")
		// 	counter2++
		// 	counter++
		// }

		// increment2()

		// fmt.Println(counter2)
	}

	increment()
	increment()
	fmt.Println(counter)
}