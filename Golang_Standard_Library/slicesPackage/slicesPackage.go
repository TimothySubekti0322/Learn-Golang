package slicesPackage

import (
	"fmt"
	"slices"
)

func MainSlicesPackage() {
	names := []string{"john", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Ringo"))
	fmt.Println((slices.Index(names, "George")))
	fmt.Println(slices.Index(names, "Timothy"))
}