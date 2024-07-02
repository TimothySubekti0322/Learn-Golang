package mathPackage

import (
	"fmt"
	"math"
)

func MainMathPackage() {
	// Round
	fmt.Println("Round(1.7) =", math.Round(1.7))

	// Floor
	fmt.Println("Floor(1.7) =", math.Floor(1.7))

	// Ceil
	fmt.Println("Ceil(1.7) =", math.Ceil(1.7))

	// Max
	fmt.Println("Max(10, 20) =", math.Max(10, 20))

	// Min
	fmt.Println("Min(10, 20) =", math.Min(10, 20))

	// Pow
	fmt.Println("Pow(2, 3) =", math.Pow(2, 3))

	// Sqrt
	fmt.Println("Sqrt(16) =", math.Sqrt(16))

	// Abs
	fmt.Println("Abs(-10) =", math.Abs(-10))
}