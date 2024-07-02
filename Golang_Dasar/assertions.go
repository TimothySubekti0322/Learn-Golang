// Merubah tipe data menjadi tipe data yang diinginkan

/* Fitur ini sering sekali digunakan ketika kita bertemu
dengan data interface kosong */

package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string: 
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	case bool:
		fmt.Println("Value", value, "is bool")
	default:
		fmt.Println("Unknown type")
	}

	// INI AKAN ERROR KARENA TIDAK SESUAI DENGAN TIPE DATA YANG DIINGINKAN
	// result = random()
	// resultInt := result.(int)
	// fmt.Println(resultInt)
}