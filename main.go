package main

import (
	"DirectBackend/api"
	_ "DirectBackend/api"
	"DirectBackend/apiHandler"
	_ "DirectBackend/apiHandler"
	_ "DirectBackend/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Hardcoded user credentials for demonstration purposes
var UserAccount = map[string]string{
	"user1": "password1",
	"user2": "password2",
}
var UserTokens = map[string]api.Token{}

func main() {
	// db.WriteUserToDB("user3", "password3")
	// password, err := db.ReadUserPasswordFromDB("user3")
	// fmt.Println(password, err)
	router := mux.NewRouter()
	router.HandleFunc("/login", apiHandler.LoginHandler).Methods("POST")
	// router.HandleFunc("/register", apiHandler.RegisterHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8090", router))
}
