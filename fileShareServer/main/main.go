package main

import (
	"github.com/exorades/fileShare/fileShareServer"
)

func main() {
	fileShare.Server{Address:"0.0.0.0:6000"}.Start()
}
