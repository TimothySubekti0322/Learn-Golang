package main

import "fmt"

func recursiveFunction(value int) int {
	if value == 1 {
		return 1 
	} else {
		return value * recursiveFunction(value - 1);
	}
}

func main() {
	fmt.Println(recursiveFunction(3))
}