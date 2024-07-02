package main

import (
	"belajar-golang-dasar/database"
	"belajar-golang-dasar/helper"
	 _ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	result := helper.SayHello("Timothy")
	result2 := helper.SayHi("Timothy")
	
	fmt.Println(result)
	fmt.Println(result2)

	fmt.Println("---------------- helper.main -----------------")

	helper.Main()
	fmt.Println(helper.Application)

	fmt.Println("---------------- database -----------------")

	fmt.Println(database.GetDatabase())
}