package encodingPackage

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func MainEncodingPackage() {
	value := "john wick"
	var encoded = base64.StdEncoding.EncodeToString([]byte(value))

	fmt.Println("encoded:", encoded)

	decoded , err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding")
	} else {
		fmt.Println("decoded:", string(decoded))
	}

	// CSV Reader
	csvString := "Timothy Subekti, 23\n" +
		"John Doe, 25\n" +
		"Marvel Subekti, 23"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	// CSV Writer
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"Name", "Age"})
	writer.Write([]string{"Timothy Subekti", "23"})
	writer.Write([]string{"John Doe", "25"})
	writer.Write([]string{"John Wick", "40"})

	writer.Flush()

}