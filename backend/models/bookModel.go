package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	UserID      uint `gorm:"foreignKey:UserID"`
	Name        string
	Author      string
	PublishYear int
}
