/* Struct adalah sebuah template data yang digunakan untuk
menggabungkan nol atau lebih tipe data lainnya dalam satu
kesatuan */

package main

import "fmt"

type User struct {
	// AWAL DENGAN HURUF BESAR
	Email, Password string
	RememberMe bool
}

func (user User) sayHello(name string) {
	fmt.Println("Hello", name, "your email is :", user.Email)
}

func main() {
	var user User
	user.Email = "Timothy@gmail.com"
	user.Password = "123456"
	user.RememberMe = true

	fmt.Println(user)
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	fmt.Println(user.RememberMe)

	marvel := User{
		Email: "marvel@gmail.com",
		Password: "123456",
		RememberMe: true,
	}

	fmt.Println(marvel)

	joko := User{"joko@gmail.com", "123456", true}

	fmt.Println(joko)

	// Struct Method
	joko.sayHello("joko")
}
