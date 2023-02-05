package models

import (
	"github.com/jinzhu/gorm"
	"go-todo/config"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	ID        string `gorm: ""json:"id"`
	Title     string `gorm: ""json:"title"`
	Completed bool   `gorm: ""json:"completed"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

func (t *Todo) CreateTodo() *Todo {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func GetTodo() []Todo {
	var Todo []Todo
	db.Find(&Todo)
	return Todo
}

func GetTodoById(Id int64) (*Todo, *gorm.DB) {
	var getTodo Todo
	db := db.Where("ID=?", Id).Find(&getTodo)
	return &getTodo, db
}
func DeleteTodo(ID int64) Todo {
	var todo Todo
	db.Where("ID=?", ID).Delete(todo)
	return todo
}
