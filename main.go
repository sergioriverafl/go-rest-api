package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sergioriverafl/go-rest-api/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Index Routes
	router.HandleFunc("/", handlers.IndexRoute)

	// Tasks Routes
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetOneTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")

	fmt.Println("Server started on port ", 65535)
	log.Fatal(http.ListenAndServe(":65535", router))
}
