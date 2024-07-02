package main

import "fmt"

type Address struct {
	City string
	Province string
	Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // copy value address 1 ke address 2

	address2.City = "Jakarta"

	fmt.Println(address1) // address 1 tidak berubah, karena address 2 copy value address 1
	fmt.Println(address2)

	// -------------------- Pointer --------------------

	fmt.Println("----------- Pointer ------------")

	address3 := &address1 // address 3 adalah pointer ke ADDRESS 1
	// var address3 *Address = &address1

	address3.City = "Bandung"

	fmt.Println(address1) // address 1 berubah, karena address 3 adalah pointer ke address 1
	fmt.Println(*address3) // *address3 adalah value dari pointer address 3

	address4 := &address1 // address 4 adalah pointer ke ADDRESS 1

	address4.City = "Surabaya"
	
	address4 = &Address{"Malang", "Jawa Timur", "Indonesia"} // address 4 diubah menjadi pointer ke ADDRESS baru

	fmt.Println(address1) // address 1 tidak berubah, karena address 4 sudah diubah menjadi pointer ke ADDRESS baru

	fmt.Println("Gunakan Asterisk")

	address5 := &address1 // address 5 adalah pointer ke ADDRESS 1

	*address5 = Address{"Semarang", "Jawa Tengah", "Indonesia"} // *address5 = value dari ADDRESS si address5 kita ubah jadi Adress{"Semarang", "Jawa Tengah", "Indonesia"}

	// Kode diatas membuat value yang diacu oleh address5 menjadi nilai baru

	fmt.Println("Address 1 setelah diubah melalui asterik operator :", address1)

	// -------------------- New Pointer --------------------

	fmt.Println("----------- New Pointer ------------")

	alamat1 := new(Address)
	alamat2 := alamat1
	alamat3 := alamat1

	fmt.Println("alamat3 = " , *alamat3)

	alamat2.Country = "Indonesia"

	fmt.Println("alamat3 = ", *alamat3)
}