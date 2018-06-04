package fileShare

import (
	"net"
	"fmt"
	"os"
	"log"
	"strconv"
	"io"
)

func sendFileToClient(connection net.Conn) {

	fmt.Println("A client has connected!")

	defer connection.Close()

	//creating file

	//for testing

	f, err := os.Open(os.Args[0])

	if err != nil {

		log.Fatalln("my program broke opening: ", err.Error())

	}

	defer f.Close()

	nf, err := os.Create("newFile.txt")

	if err != nil {

		log.Fatalln("my program broke creating: ", err.Error())

	}

	nf.Write([]byte("Hello World"))

	file := nf



	// for testing done

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