package main

import (
	"crud/handles"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", handles.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user", handles.SearchAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", handles.SearchUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", handles.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", handles.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Api rodando na porta 2121")
	log.Fatal(http.ListenAndServe(":2121", router))
}
