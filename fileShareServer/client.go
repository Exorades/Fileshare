package fileShare

import (
"bufio"
"net"
"log"
"io"
"fmt"
)

type client struct {
	name string
	room *room
	reader *bufio.Reader
	writer *bufio.Writer
	connection net.Conn
}


func(client *client) read() {
	for {
		line, err := client.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				message := message{"Client disconnected " + client.name + "\n", client}
				client.room.broadcast(message)
			} else {
				log.Printf("Exception: %v", err)
			}
			return
		} else {
			client.room.outgoing <- message{line, client}
		}
	}
}

func (client *client) write(message message) {
	messageText := fmt.Sprintf("%s: %s", message.sender.name, message.text)
	client.writer.WriteString(messageText)
	client.writer.Flush()
}


func (client *client) disconnect() {
	client.connection.Close()
	log.Printf("Disconnecting %v", client.name)
}
