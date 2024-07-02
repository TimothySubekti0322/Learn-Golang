/* Dalam Golang saat kita membuat tipe data, otomatis
diisikan nilai default */
/* Nil hanya bisa untuk tipe data interface, function,
map, slice, pointer dan channel */

package main

import "fmt"

// Ini Gak bisa karena returnnya map
// func NewMap(name string) string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return name
// 	}
// }

func NewMap(name string) map[string] string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// data := NewMap("Timothy")
	data := NewMap("")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}
}