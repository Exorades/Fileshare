package fileShare

import (
	"net"
	"fmt"
	"os"
	"log"
	"strconv"
	"io"
)
const BUFFERSIZE = 1024
func SendFileToClient(connection net.Conn, filename string) {

	fmt.Println("A client has connected!")

	defer connection.Close()

	// file
	path, _ := os.Getwd()
	filedir := fmt.Sprintf("%v/exorades/fileShare/fileShareServer/tmp/%v", path, filename)
	file, err := os.Open(filedir)

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