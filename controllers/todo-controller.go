package controllers

import (
	"fmt"
	"go-todo/config"
	"go-todo/models"
	"net/http"
)

var (
	id        string
	title     string
	completed bool
	database  = config.Main()
)

func Display(w http.ResponseWriter, r *http.Request) {
	_, err := database.Query(`SELECT * FROM todos`)
	if err != nil {
		fmt.Println(err)
	}
	var todos []models.Todo

	for statement.Next() {
		err = statement.scan(&id, &title, &completed)
		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:        id,
			Title:     title,
			Completed: completed,
		}

		todos = append(todos, todo)
	}
}

//var NewTodo models.Todo
//
//func CreateTodo(w http.ResponseWriter, r *http.Request) {
//	CreateTodo := &models.Todo{}
//}
