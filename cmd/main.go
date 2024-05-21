package main

import (
	"net-cat/server"
	"os"
)

var PORT = "4040"

func main() {
	if len(os.Args[1:]) == 1 {
		PORT = os.Args[1]
	}

	server.StartServer(PORT)
}
