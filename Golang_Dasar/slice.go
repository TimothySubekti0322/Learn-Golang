package main

import "fmt"

func main() {
	names := [...]string{"Timothy", "Irham", "Rizki", "Wandi"}

	// var ops []string = names[0:3]
	ops := names[0:3]

	fmt.Println(ops)

	names[2] = "Rizal"

	fmt.Println(ops)

	fmt.Println("Slice is pointer to Array")

	fmt.Println("len(ops) =", len(ops))
	fmt.Println("cap(ops) =", cap(ops))
	
	ops[0] = "Timothy"

	fmt.Println("Setelah diubah oleh slice, array akan berubah",names)

	fmt.Println("Append Slice")

	ops = append(ops, "Doe")
	ops = append(ops, "John")

	fmt.Println(ops)
	fmt.Println(names)

	fmt.Println("cap(ops) =", cap(ops))

	fmt.Println("Ada Array Baru yang dibuat oleh Slice ops, jadi sudah tidak berhubungan dengan array names")

	newSlice := make([]string, 2, 5) // 2 adalah panjang slice, 5 adalah kapasitas slice
	
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice[0] = "Timothy"
	newSlice[1] = "Doe"
	// newSlice[2] = "John" // Error karena panjang slice hanya 2 -> Harus di Appned

	newSlice = append(newSlice, "John")

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))

	newSlice = append(newSlice, "Jane")
	newSlice = append(newSlice, "Doe")
	newSlice = append(newSlice, "John")

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("----Copy Slice----")

	fromSlice := names[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println("toSlice =", toSlice)

	// HATI HATI BIKIN ARRAY DENGAN SLICE
	// array := [...]string{1,2,3}
	// slice := []string{1,2,3} 
}