package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func filterOdd(numbers ...int) []int {
	var result []int
	for _, number := range numbers {
		if number%2 != 0 {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	result := sumAll(10, 10, 10, 10, 10)
	fmt.Println(result)

	// Using Slice
	numbers := []int{10, 10, 10, 10, 10}
	result = sumAll(numbers...)
	fmt.Println(result)

	fmt.Println(filterOdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}