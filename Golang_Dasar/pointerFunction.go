package main

import "fmt"

type Address struct {
	City string
	Province string
	Country string
}

func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func ChangeAddressToMalaysia(address *Address) {
	address.Country = "Malaysia"
}

func main() {
	address := Address{"Subang", "Jawa Barat", "Amerika"}
	ChangeAddressToIndonesia(address)

	fmt.Println(address) // Country tidak berubah, karena address di passing by value

	// ChangeAddressToMalaysia(&address)

	// fmt.Println(address) // Country berubah, karena address di passing by reference	

	// Atau

	address2 := &address
	ChangeAddressToMalaysia(address2)

	fmt.Println(address) // Country berubah, karena address di passing by reference
}