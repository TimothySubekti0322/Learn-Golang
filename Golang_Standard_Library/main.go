package main

import (
	"belajar-golang-standard-library/bufioPackage"
	"belajar-golang-standard-library/containerListPackage"
	"belajar-golang-standard-library/containerRingPackage"
	createfile "belajar-golang-standard-library/createFile"
	encodingPackage "belajar-golang-standard-library/encodingPackage"
	"belajar-golang-standard-library/errorPackage"
	"belajar-golang-standard-library/flagPackage"
	"belajar-golang-standard-library/fmtPackage"
	"belajar-golang-standard-library/ioPackage"
	"belajar-golang-standard-library/mathPackage"
	"belajar-golang-standard-library/osPackage"
	"belajar-golang-standard-library/pathPackage"
	reflectpackage "belajar-golang-standard-library/reflectPackage"
	regexpPackage "belajar-golang-standard-library/regexpPackage"
	"belajar-golang-standard-library/slicesPackage"
	"belajar-golang-standard-library/sortPackage"
	"belajar-golang-standard-library/strconvPackage"
	"belajar-golang-standard-library/stringPackage"
	"belajar-golang-standard-library/timePackage"
	"fmt"
)

func main() {
	fmtPackage.MainPrintf()

	fmt.Println("----------------- This is error Package -----------------")
	errorPackage.MainError()

	fmt.Println("----------------- This is os Package -----------------")
	osPackage.MainOS()

	fmt.Println("----------------- This is flag Package -----------------")
	flagPackage.MainFlag()

	fmt.Println("----------------- This is string Package -----------------")
	stringPackage.MainString()

	fmt.Println("----------------- This is strconv Package -----------------")
	strconvPackage.MainStrconvPackage()

	fmt.Println("----------------- This is math Package -----------------")
	mathPackage.MainMathPackage()

	fmt.Println("----------------- This is container List Package -----------------")
	containerListPackage.MainContainerListPackage()

	fmt.Println("----------------- This is container Ring Package -----------------")
	containerRingPackage.MainContainerRingPackage()

	fmt.Println("----------------- This is Sort Package -----------------")
	sortPackage.MainSortPackage()

	fmt.Println("----------------- This is Time Package -----------------")
	timePackage.MainTimePackage()

	fmt.Println("----------------- This is Reflect Package -----------------")
	reflectpackage.MainReflectPackage()

	fmt.Println("----------------- This is Regexp Package -----------------")
	regexpPackage.MainRegexpPackage()

	fmt.Println("----------------- This is Encoding Package -----------------")
	encodingPackage.MainEncodingPackage()

	fmt.Println("----------------- This is Slices Package -----------------")
	slicesPackage.MainSlicesPackage()

	fmt.Println("----------------- This is path Package -----------------")
	pathPackage.MainPathPackage()

	fmt.Println("----------------- This is io Package -----------------")
	ioPackage.MainIOPackage()

	fmt.Println("----------------- This is bufio Package -----------------")
	bufioPackage.MainBufioPackage()

	fmt.Println("----------------- This is Create File -----------------")
	createfile.CreateFileMain()

}