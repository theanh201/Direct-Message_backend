package main

import (
	"DirectBackend/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", controller.AccLogin).Methods("POST")
	router.HandleFunc("/register", controller.AccRegister).Methods("POST")
	router.HandleFunc("/get-self-info", controller.AccGetSelfInfo).Methods("GET")
	// router.HandleFunc("/get-info", controller.SearchControllerGetUsername).Methods("POST")
	router.HandleFunc("/update-avatar", controller.AccUpdateAvatar).Methods("PUT")
	router.HandleFunc("/update-background", controller.AccUpdateBackground).Methods("PUT")
	router.HandleFunc("/update-email", controller.AccUpdateEmail).Methods("PUT")
	router.HandleFunc("/update-phone-number", controller.AccUpdatePhoneNumber).Methods("PUT")
	router.HandleFunc("/update-password", controller.AccUpdatePassword).Methods("PUT")
	router.HandleFunc("/update-name", controller.AccUpdateName).Methods("PUT")

	// router.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8090", router))
}
func test(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./asset/test.txt")
	http.ServeFile(w, r, "./asset/test2.txt")
}
