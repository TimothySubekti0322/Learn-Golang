package main

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "eko" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	err := SaveData("timothy", nil)
	if err != nil {
		switch err.(type) {
		case *validationError:
			println("validation error", err.Error())
		case *notFoundError:
			println("not found error", err.Error())
		default:
			println("unknown error", err.Error())
		}
	} else {
		println("success")
	}
}