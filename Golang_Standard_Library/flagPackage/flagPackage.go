package flagPackage

import (
	"flag"
	"fmt"
)

func MainFlag() {
	var username *string = flag.String("username", "admin", "Put your username")
	var password *string = flag.String("password", "admin", "Put your password")
	var host *string = flag.String("host", "localhost", "Put your host")
	var port *int = flag.Int("port", 8080, "Put your port")

	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)
}

