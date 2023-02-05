package routes

import (
	"github.com/gorilla/mux"
	"go-todo/controllers"
)

var TodoRoute = func(router mux.Router) {
	router.HandleFunc("/todo/", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo", controllers.GetTodo).Methods("GET") //fetch
	router.HandleFunc("/todo/{todoId}", controllers.GetTodoById).Methods("GET")
	router.HandleFunc("/todo/{todoId}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{todoId}", controllers.Delete).Methods("DELETE")
}

func main() {

}
