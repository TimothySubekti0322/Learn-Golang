// Panic adalah function yang bisa kita gunakan untuk menghentikan program tetapi defer tetap dijalankan

package main

import "fmt"

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp() //Jika Error , tetep aja dijalankan -> Harus Taruh endApp di awal
	if error {
		panic("ERROR APLIKASI")
	}
	fmt.Println("App running properly")
}

func main() {
	 runApp(true)
}
