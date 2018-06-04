package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleMessage(msg string, conn net.Conn) {
	command := strings.TrimSpace(msg)
	commandMap := map[string]string{
		"keyword": "whoaa, I understand this",
	}
	newCommand, exists := commandMap[command]
	if !exists {
		newCommand = "Invalid command!"
	}
	conn.Write([]byte(newCommand + "\n"))
}

func handle(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = string(message)
		if message == "" {
			return
		}
		fmt.Print("Message Received: ", message)
		handleMessage(message, conn)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")
	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}
