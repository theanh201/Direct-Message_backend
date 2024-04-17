package main

import (
	"DirectBackend/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	router.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8090", router))
}
