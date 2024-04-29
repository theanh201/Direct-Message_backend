package main

import (
	"DirectBackend/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Account post
	router.HandleFunc("/login", controller.AccPostLogin).Methods("POST")
	router.HandleFunc("/register", controller.AccPostRegister).Methods("POST")
	// Account get
	router.HandleFunc("/get-self-info", controller.AccGetSelfInfo).Methods("GET")
	router.HandleFunc("/get-avatar", controller.AccGetAvatar).Methods("GET")
	router.HandleFunc("/get-background", controller.AccGetBackGround).Methods("GET")
	router.HandleFunc("/get-by-name", controller.AccGetUserByName).Methods("GET")
	router.HandleFunc("/get-by-email", controller.AccGetUserByEmail).Methods("GET")
	// Account put
	router.HandleFunc("/update-avatar", controller.AccPutAvatar).Methods("PUT")
	router.HandleFunc("/update-background", controller.AccPutBackground).Methods("PUT")
	router.HandleFunc("/update-email", controller.AccPutEmail).Methods("PUT")
	router.HandleFunc("/update-password", controller.AccPutPassword).Methods("PUT")
	router.HandleFunc("/update-name", controller.AccPutName).Methods("PUT")
	router.HandleFunc("/update-private-status", controller.AccPutPrivateStatus).Methods("PUT")
	// Account delete
	router.HandleFunc("/delete-self", controller.AccDeleteSelf).Methods("DELETE")
	// Prekey bundle
	router.HandleFunc("/get-prekey-bundle", controller.PrekeyBundleGet).Methods("GET")
	router.HandleFunc("/update-prekey-bundle", controller.PrekeyBundlePut).Methods("PUT")
	// Add Friend Request
	router.HandleFunc("/add-friend-request", controller.FriendRequestPost).Methods("POST")
	// Get friendRequest
	router.HandleFunc("/get-friend-request", controller.FriendRequestGet).Methods("GET")
	// Accept friend request
	router.HandleFunc("/accept-friend-request", controller.FriendRequestPostAccept).Methods("POST")
	// Reject friend request
	router.HandleFunc("/reject-friend-request", controller.FriendRequestPostReject).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
