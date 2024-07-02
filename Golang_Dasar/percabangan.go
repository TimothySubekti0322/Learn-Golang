package main

import "fmt"

func main() {

	// If Else
	var suhu int = 17

	if suhu < 25 {
		fmt.Println("Dingin")
	} else if suhu == 25 {
		fmt.Println("Sejuk")
	} else {
		fmt.Println("Panas")
	}

	// Short Statement
	var name string = "Timothy"

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	// Switch Case
	var day int = 3

	switch day {
	case 0 :
		fmt.Println("Minggu")
	case 1 :
		fmt.Println("Senin")
	case 2 :
		fmt.Println("Selasa")
	case 3 :
		fmt.Println("Rabu")
	case 4 :
		fmt.Println("Kamis")
	case 5 :
		fmt.Println("Jumat")
	case 6 :
		fmt.Println("Sabtu")
	default :
		fmt.Println("Hari tidak dikenal")
	}


	// Switch Case with Short Statement
	switch newDay := 3; newDay {
	case 0 :
		fmt.Println("Minggu")
	case 1 :
		fmt.Println("Senin")
	case 2 :
		fmt.Println("Selasa")
	case 3 :
		fmt.Println("Rabu")
	case 4 :
		fmt.Println("Kamis")
	case 5 :
		fmt.Println("Jumat")
	case 6 :
		fmt.Println("Sabtu")
	default :
		fmt.Println("Hari tidak dikenal")
	}

	// Switch Case with Multiple Case
	var fullName string = "Timothy Subekti"

	switch {
		case len(fullName) > 10 :
			fmt.Println("Nama terlalu panjang")
		case len(fullName) > 5 :
			fmt.Println("Nama sudah benar")
		default :
			fmt.Println("Nama terlalu pendek")
	}

}