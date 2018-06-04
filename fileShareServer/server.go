package fileShare

import (
	"net"
	"bufio"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
	"fmt"
)


type Server struct {
	Address string
}



func (server Server) Start() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", "0.0.0.0:6000")


	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}


func handle(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = string(message)
		if message == "" {
			return
		}
		fmt.Print("Message Received: ", message)
		handleCommand(message, conn)
	}
}


func (server Server) initDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "chat.db")
	if err != nil {
		panic("DB FAIL")
	}
	db.AutoMigrate(&command{})
	return db
}


//func listFiles() []string {
//
//	path, _ := os.Getwd()
//	fullPath := path + "/exorades/fileShare/fileShareServer/tmp/"
//	files, err := ioutil.ReadDir(fullPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//
//	for _, f := range files {
//		fmt.Println(f.Name())
//	}
//	return files
//}

func handleCommand(command string, conn net.Conn) {
	var file string

	msg := strings.TrimSpace(command)

	words := strings.Fields(msg)
	command = words[0]
	length := len(words)
	if length == 2 {
		file = words[1]
	}

	switch {
	case command == "#upload":
		fmt.Printf("Uploading file %v", file)
	case command == "#download":
		fmt.Printf("Downloading file %v", file)
		SendFileToClient(conn, file)
	case command == "#list":
		fmt.Printf("Listing files")
		//listFiles()
	default:
		fmt.Printf("Incorrect command")
	}

	conn.Write([]byte(command + "" + file + "\n"))
}

