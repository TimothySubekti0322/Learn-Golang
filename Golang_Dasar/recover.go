// Recover adalah function yang bisa kita gunakan untuk menangkap data panic

package main

import "fmt"

func endOfApp() {
	fmt.Println("End app")

	// Cara yang benar untuk recover. INGAT DEFER SELALU DI EXECUTE
	message := recover()
	fmt.Println("Terjadi Error : ", message)
}

func checkAppError() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error : ", message)
	}
}

func runTheApp(error bool) {
	defer checkAppError() //Jika Error , tetep aja dijalankan -> Harus Taruh defer di awal
	if error {
		panic("ERROR APLIKASI")
	}

	// Recover dengan cara yang Salah. SETELAH PANIC, Program berhenti, message tidak terambil
	// message := recover()
	// fmt.Println("terjadi panic", message)

	fmt.Println("App running properly")
}

func main() {
	 runTheApp(true)
	 fmt.Println("Timothy Subekti")
}

