package server

import (
	"bufio"
	"net"
	"strings"
	"time"
)

var (
	history      = ""
	clients      = make(map[*Client]bool)
	clientMutex  = make(chan bool, 1)
	usernames    = []string{}
	numberOfuser = 0
)

type Client struct {
	Connection net.Conn
	Username   string
}

// Parsing incoming requests
func HandleRequest(conn net.Conn) {
	defer func() { numberOfuser-- }()
	currentTime := time.Now()
	conn.Write([]byte(Asciicum()))

	reader := bufio.NewReader(conn)

	// Ask for the username.
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	username, _ := reader.ReadString('\n')
	for _, v := range usernames {
		if v == username {
			conn.Write([]byte("The username has already been taken"))
			HandleRequest(conn)
		}
	}
	usernames = append(usernames, username)
	username = strings.Replace(username, "\n", "", -1)
	// Register new client
	client := &Client{Connection: conn, Username: username}
	clientMutex <- true
	clients[client] = true
	<-clientMutex
	if numberOfuser > 10 {

		conn.Write([]byte("The limit of users exceeded. Maximum is 10\n"))
		conn.Close()
	}

	BroadcastJoinMSG(client, "\n"+client.Username+" has joined our chat!")
	conn.Write([]byte(history))

	for {

		client.Connection.Write([]byte("[" + currentTime.Format("2006-01-02 15:04:05") + "]" + "[" + client.Username + "]" + ":"))
		message, err := reader.ReadString('\n')
		if err != nil {

			BroadcastJoinMSG(client, "\n"+client.Username+" has left our chat...")
			usernames = deletor(client.Username)
			clientMutex <- true
			delete(clients, client)
			<-clientMutex
			return
		}
		message = strings.Replace(message, "\n", "", -1)

		// Sending message to all active clients
		if len(message) < 1 {
			continue
		} else {
			history += "[" + currentTime.Format("2006-01-02 15:04:05") + "]" + "[" + client.Username + "]" + ":" + message + "\n"
			BroadcastMessage(client, message)
		}
	}
}

func deletor(name string) []string {
	var res []string
	for _, v := range usernames {
		if v != name+"\n" {
			res = append(res, v)
		}
	}

	return res
}
