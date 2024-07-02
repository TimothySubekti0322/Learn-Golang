package main

import "fmt"

func main() {
	var name string
	var age int
	var isMarried bool

	name = "Timothy"
	age = 23
	isMarried = false

	fmt.Println("Nama saya adalah", name)
	fmt.Println("Umur saya adalah", age)
	if isMarried {
		fmt.Println("Status menikah saya adalah", "sudah menikah")
	} else {
		fmt.Println("Status menikah saya adalah", "belum menikah")
	}

	// Deklarasi variabel dengan tipe data langsung

	fmt.Println("-------Deklarasi Variabel dengan Tipe Data Langsung-------")

	var name2 = "Subekti"

	fmt.Println("Nama belakang saya adalah", name2)

	// Deklarasi variabel tanpa menggunakan kata kunci var

	fmt.Println("-------Deklarasi Variabel Tanpa var-------")

	name3 := "Timothy Subekti"

	fmt.Println("Nama lengkap saya adalah", name3)
	fmt.Printf("Tipe Data name3 adalah %T\n", name3)

	// Deklarasi Multi Variabel

	fmt.Println("-------Deklarasi Multi Variabel-------")

	var (
		firstname = "Timothy"
		lastName = "Subekti"
	)

	fmt.Println("Nama saya adalah", firstname, lastName)

	siblingsFirstName, siblingsLastName := "Mar", "Subekti"

	fmt.Println("Nama saudara saya adalah", siblingsFirstName, siblingsLastName)

	siblingsFirstName = "Marvel"

	fmt.Println("Nama saudara saya adalah", siblingsFirstName, siblingsLastName)
}