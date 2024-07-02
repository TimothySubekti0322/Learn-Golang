package strconvPackage

import (
	"fmt"
	"strconv"
)

func MainStrconvPackage() {
	result, err := strconv.ParseBool("SALAH") // Ini akan Error , hanya bisa menerima "true" atau "false"
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// ParseInt
	resultInt, err := strconv.ParseInt("100000", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resultInt)
	}

	// Atoi
	resultAtoi, err := strconv.Atoi("700")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultAtoi)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt)
	
}