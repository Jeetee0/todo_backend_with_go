package main

import (
	"fmt"
	"log"
	"net/http"

	"mytodoapp_withgo/todo"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func main() {
	port := ":8888"
	router := mux.NewRouter()

	fmt.Println("Settin up server, enabling CORS . . .")

	router.HandleFunc("/todos", todo.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", todo.GetTodo).Methods("GET")
	router.HandleFunc("/todos", todo.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", todo.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", todo.DeleteTodo).Methods("DELETE")

	http.Handle("/", router)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                            // All origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Allowing only get, just an example
	})

	// Start the server
	fmt.Printf("Server is ready and is listening at port %s . . .", port)
	fmt.Println("")
	log.Fatal(http.ListenAndServe(port, c.Handler(router)))

}
