package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("a + b =", c)
	fmt.Println("a - b =", d)
	fmt.Println("a * b =", e)
	fmt.Println("a / b =", f)
	fmt.Println("a % b =", g)

	fmt.Println("Augmented Assignments")
	a += 10
	fmt.Println("a += 10 ->", a)

	fmt.Println("Unary Operator")
	a++
	fmt.Println("a++ ->", a)
	a--
	fmt.Println("a-- ->", a)
	a = -a
	fmt.Println("a = -a->", a)
	
	booleanKebalikan := !true
	fmt.Println("!true = ", booleanKebalikan)

}