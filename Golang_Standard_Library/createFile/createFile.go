package createfile

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewfile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(message)

	fmt.Println("Create file success")

	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var message string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(message)

	fmt.Println("add to File Success")

	return nil
}

func CreateFileMain() {
	// createNewfile("sample.txt", "This is sample txt")
	result, _ := readFile("sample.txt")
	fmt.Println(result)

	// addToFile("sample.txt", "\nThis is ADD MESSAGE")
}