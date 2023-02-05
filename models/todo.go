package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

type Todo struct {
	gorm.Model
	ID
	Title
	Completed
}
