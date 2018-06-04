package main

import "github.com/exorades/Fileshare/fileShareServer"

func main() {
	fileShare.Server{Address:"0.0.0.0:6000"}.Start()
}
