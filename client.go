package main

import "net"
import "fmt"
import "bufio"
import (
	"os"
	"log"
	"io"
	"os/user"
)

func main() {
	conn, _ := net.Dial("tcp", "10.93.7.243:6000")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

const BUFFERSIZE = 1024
func SendFileToServer(connection net.Conn, filedest string) {
	fmt.Println(" A client has connected!")
	defer connection.Close()

	// file
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	fmt.Println( usr.HomeDir + "/" )

	filedir := fmt.Sprintf("%v/%v", usr.HomeDir, filedest))
	file, err := os.Open(filedir)

	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	// file done

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