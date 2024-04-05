package main

import (
	"DirectBackend/apiHandler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// db.WriteUserToDB("user3", "password3")
	// password, err := db.ReadUserPasswordFromDB("user3")
	// fmt.Println(password, err)
	router := mux.NewRouter()
	router.HandleFunc("/login", apiHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/register", apiHandler.RegisterHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8090", router))
}
