package server

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	TYPE = "tcp"
)

var PORT = "4040"

func StartServer(PORT string) {
	// Listening to all incoming messages
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()
	fmt.Println("Listening on the port " + ":" + PORT)

	for {

		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		numberOfuser++
		go HandleRequest(conn)
	}
}

func Asciicum() string {
	asciiwelcome := `
Welcome to TCP-Chat!
        _nnnn_
       dGGGGMMb
      @p~qp~~qMb
      M|@||@) M|
      @,----.JM|
     JS^\__/  qKL
    dZP        qKRb
   dZP          qKKb
  fZP            SMMb
  HZM            MMMM
  FqM            MMMM
 __| ".        |\dS"qML
 |    '.       | '' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     '-'       '--'
`
	return asciiwelcome
}
