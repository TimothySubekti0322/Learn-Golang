package bufioPackage

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func MainBufioPackage() {
	// Read
	input := strings.NewReader("this is long string\nwith new line\n")

	reader := bufio.NewReader(input)

	for {
		line, _ , err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	// Write
	write := bufio.NewWriter(os.Stdout)
	_, _ = write.WriteString("Hello World\n")
	_, _ = write.WriteString("Selamat Belajar\n")
	write.Flush()
}