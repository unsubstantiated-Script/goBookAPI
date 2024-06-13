package models

import (
	"github.com/jinzhu/gorm"
	"goBookAPI/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Author string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
