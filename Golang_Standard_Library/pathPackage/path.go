package pathPackage

import (
	"fmt"
	"path"
	"path/filepath"
)

func MainPathPackage() {
	fmt.Println(path.Dir("pathPackage/path/world.go"))
	fmt.Println(path.Base("pathPackage/path/world.go"))
	fmt.Println(path.Ext("hello/world.html"))
	fmt.Println(path.Join("hello", "world.go"))

	fmt.Println("--- File Path ---")

	fmt.Println(filepath.Dir("pathPackage/path/world.go"))
	fmt.Println(filepath.Base("pathPackage/path/world.go"))
	fmt.Println(filepath.Ext("hello/world.html"))
	fmt.Println(filepath.IsAbs("D:\\Web_Development_Learning\\First_Golang_With_Eko"))
	fmt.Println(filepath.IsLocal("hello/world.html"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}