package fileShare

import "github.com/jinzhu/gorm"

type command struct {
	gorm.Model
	Text string
	Sender string
}