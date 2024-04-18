package main

import (
	"DirectBackend/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", controller.AccControllerLogin).Methods("POST")
	router.HandleFunc("/register", controller.AccControllerRegister).Methods("POST")

	// router.HandleFunc("/get-info", controller.SearchControllerGetUsername).Methods("POST")
	router.HandleFunc("/update-avatar", controller.AccControllerUpdate).Methods("PUT")
	// router.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8090", router))
}
func test(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "./asset/test.txt")
	http.ServeFile(w, r, "./asset/test2.txt")
}
