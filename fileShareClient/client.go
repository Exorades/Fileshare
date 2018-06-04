package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/user"
	"strconv"
	"strings"
)

const BUFFERSIZE = 1024

func saveFile(connection net.Conn) {
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)
	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")
	newFile, err := os.Create(fileName)
	print(fileName)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	var receivedBytes int64
	for {
		if (fileSize - receivedBytes) < BUFFERSIZE {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
			break
		}
		io.CopyN(newFile, connection, BUFFERSIZE)
		receivedBytes += BUFFERSIZE
	}
	fmt.Println("Received file completely!")
}

func main() {
	conn, _ := net.Dial("tcp", "10.93.7.243:6000")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		if strings.HasPrefix(text, "#download") {
			saveFile(conn)
		} else if strings.HasPrefix(text, "#upload") {
			SendFileToServer(conn, strings.TrimSpace(text))
		} else {
			conn.Write([]byte(text))
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: " + message)
		}
	}
}

func SendFileToServer(connection net.Conn, filedest string) {
	fmt.Println(" A client has connected!")
	defer connection.Close()
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr.HomeDir + "/")
	filedir := fmt.Sprintf("%v/%v", usr.HomeDir, filedest)
	file, err := os.Open(filedir)
	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	sendBuffer := make([]byte, BUFFERSIZE)
	fmt.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	fmt.Println("File has been sent, closing connection!")
	return
}
