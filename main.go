package go_todo

import (
	"github.com/gorilla/mux"
	"go-todo/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.TodoRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
