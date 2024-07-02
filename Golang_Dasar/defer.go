package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	fmt.Println("Run Application")
	defer logging() // Ini artinya kita akan memanggil logging ketika function selesai
}

func main() {
	runApplication()
}