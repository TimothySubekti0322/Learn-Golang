package errorPackage

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	} else if id != "1" {
		return NotFoundError
	}

	return nil
}

func MainError() {
	err := GetById("1")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Error:", ValidationError.Error())
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Error:", NotFoundError.Error())
		} else {
			fmt.Println("Error:", err.Error())
		}
	}

	fmt.Println("No Error Found")
}