package fileShare

import (
"net"
"log"
"strings"
"github.com/jinzhu/gorm"
	"fmt"
)

type room struct {
	name string
	outgoing chan message
	joining chan *client
	clients []*client
	shutdown chan struct{}
	listener net.Listener
	db *gorm.DB
}


func (room *room) broadcast(message message) {
	for _, client := range room.clients {
		if client.connection != message.sender.connection {
			log.Println(message)
			client.write(message)
		}
	}
}

func (room *room) addClient(client *client) {
	room.clients = append(room.clients, client)
}


func handleCommand(command string) string {
	var file string

	msg := strings.TrimSpace(command)

	words := strings.Fields(msg)
	command = words[0]
	length := len(words)
	if length == 2 {
		file = words[1]
	}
	log.Println(words)


	switch {
		case command == "#upload":
			fmt.Printf("Uploading file %v", file)
		case command == "#download":
			fmt.Printf("Downloading file %v", file)
		case command == "#list":
			fmt.Printf("Listing files")
		default:
			log.Fatal("Incorrect command")
	}

	return command + file + "\n"
}

func (room *room) listen() {
	go func() {
		log.Println("Room Listen")
		for {
			select {
			case message := <- room.outgoing:
				if strings.HasPrefix(message.text, "#") {
					handleCommand(message.text)
				} else {
					room.broadcast(message)
					room.db.Create(&command{Text: message.text, Sender: message.sender.name})
				}
			}
		}
	}()
}
