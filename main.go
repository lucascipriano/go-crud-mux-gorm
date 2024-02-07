package main

import (
	"crud/handles"
	"crud/initializers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init(){
  initializers.ConnectDb()
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/user", handles.CreateUser).Methods(http.MethodPost)
  router.HandleFunc("/user", handles.SearchAllUsers).Methods(http.MethodGet)
  router.HandleFunc("/user/{uuid}", handles.SearchUser).Methods(http.MethodGet)
  router.HandleFunc("/user/{uuid}", handles.UpdateUser).Methods(http.MethodPut)
  router.HandleFunc("/user/{uuid}", handles.DeleteUser).Methods(http.MethodDelete)

  fmt.Println("Api rodando na porta 2121")
  log.Fatal(http.ListenAndServe(":2121", router))
}
