package main

import (
	"DirectBackend/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Account
	router.HandleFunc("/login", controller.AccLogin).Methods("POST")
	router.HandleFunc("/register", controller.AccRegister).Methods("POST")
	router.HandleFunc("/update-avatar", controller.AccUpdateAvatar).Methods("PUT")
	router.HandleFunc("/update-background", controller.AccUpdateBackground).Methods("PUT")
	router.HandleFunc("/update-email", controller.AccUpdateEmail).Methods("PUT")
	router.HandleFunc("/update-password", controller.AccUpdatePassword).Methods("PUT")
	router.HandleFunc("/update-name", controller.AccUpdateName).Methods("PUT")
	router.HandleFunc("/update-private-status", controller.AccUpdatePrivateStatus).Methods("PUT")
	router.HandleFunc("/get-self-info", controller.AccGetSelfInfo).Methods("GET")
	router.HandleFunc("/get-avatar", controller.AccGetAvatar).Methods("GET")
	router.HandleFunc("/get-background", controller.AccGetBackGround).Methods("GET")
	router.HandleFunc("/delete-self", controller.DeleteSelf).Methods("DELETE")
	// Search friend
	router.HandleFunc("/get-by-name", controller.AccGetUserByName).Methods("GET")
	router.HandleFunc("/get-by-email", controller.AccGetUserByEmail).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
