package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
			conn.Write([]byte(text))
			saveFile(conn)
		} else if strings.HasPrefix(text, "#upload") {
			conn.Write([]byte(text))
			SendFileToClient(conn, strings.TrimSpace(text)[8:])
		} else {
			conn.Write([]byte(text))
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: " + message)
		}
	}
}

func SendFileToClient(connection net.Conn, filename string) {

	fmt.Println("/n A client has connected!")

	defer connection.Close()

	// file
	// path, _ := os.Getwd()
	file, err := os.Open(filename)

	if err != nil {

		log.Fatalln("my program broke opening: ", err.Error())

	}

	// file done

	if err != nil {

		fmt.Println(err)

		return

	}

	fileInfo, err := file.Stat()

	if err != nil {

		fmt.Println(err)

		return

	}

	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)

	fileName := fillString(fileInfo.Name(), 64)

	fmt.Println("Sending filename and filesize!")

	connection.Write([]byte(fileSize))

	connection.Write([]byte(fileName))

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

func fillString(retunString string, toLength int) string {

	for {

		lengtString := len(retunString)

		if lengtString < toLength {

			retunString = retunString + ":"

			continue

		}

		break

	}

	return retunString

}
