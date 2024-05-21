package server

import (
	"log"
	"time"
)

func BroadcastJoinMSG(sender *Client, message string) {
	currentTime := time.Now()
	clientMutex <- true

	for client := range clients {
		if client != sender {
			_, err := client.Connection.Write([]byte(message))
			if err != nil {
				log.Println(err)
			}
			client.Connection.Write([]byte("\n" + "[" + currentTime.Format("2006-01-02 15:04:05") + "]" + "[" + client.Username + "]" + ":"))
		}
	}
	<-clientMutex
}

func BroadcastMessage(sender *Client, message string) {
	currentTime := time.Now()
	clientMutex <- true

	for client := range clients {
		if len(message) < 1 {
			break
		}
		if client != sender {
			_, err := client.Connection.Write([]byte("\n" + "[" + currentTime.Format("2006-01-02 15:04:05") + "]" + "[" + sender.Username + "]" + ":" + message))
			if err != nil {
				log.Println(err)
			}
			client.Connection.Write([]byte("\n" + "[" + currentTime.Format("2006-01-02 15:04:05") + "]" + "[" + client.Username + "]" + ":"))
		}
	}
	<-clientMutex
}
