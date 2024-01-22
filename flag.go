package main

import (
	"flag"
	"fmt"
)

func main() {
	var usernames *string = flag.String("usernames", "root", "database user name")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "root", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("username:", *usernames)
	fmt.Println("password:", *password)
	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
}
