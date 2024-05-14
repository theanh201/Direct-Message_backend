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
	router.HandleFunc("/login", controller.AccPostLogin).Methods("POST")       // done
	router.HandleFunc("/register", controller.AccPostRegister).Methods("POST") // done
	// Account get
	router.HandleFunc("/get-self-info", controller.AccGetSelfInfo).Methods("GET")
	router.HandleFunc("/get-avatar/{token}/{imgName}", controller.AccGetAvatar).Methods("GET")          // done
	router.HandleFunc("/get-background/{token}/{imgName}", controller.AccGetBackGround).Methods("GET")  // done
	router.HandleFunc("/get-by-name/{token}/{name}/{page}", controller.AccGetUserByName).Methods("GET") // done
	router.HandleFunc("/get-by-email/{token}/{email}", controller.AccGetUserByEmail).Methods("GET")     // done
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
	router.HandleFunc("/get-prekey-bundle/{token}/{email}", controller.PrekeyBundleGet).Methods("GET") // done
	router.HandleFunc("/update-prekey-bundle", controller.PrekeyBundlePut).Methods("PUT")
	// Add Friend Request
	router.HandleFunc("/add-friend-request", controller.FriendRequestPost).Methods("POST") // done
	// Get friendRequest
	router.HandleFunc("/get-friend-request/{token}", controller.FriendRequestGet).Methods("GET")
	// Accept friend request
	router.HandleFunc("/accept-friend-request", controller.FriendRequestPostAccept).Methods("POST")
	// Reject friend request
	router.HandleFunc("/reject-friend-request", controller.FriendRequestPostReject).Methods("POST")
	// Get friend list
	router.HandleFunc("/get-friend-list/{token}", controller.FriendGet).Methods("GET") // done
	// Get all message
	router.HandleFunc("/get-all-message/{token}", controller.MessageGetAll).Methods("GET")
	// Send message unencrypt to friend
	router.HandleFunc("/send-message-friend-unencrypt", controller.MessageFriendUnencrypt)
	// Get all message after time frame
	router.HandleFunc("/get-all-message-after-time/{token}/{time}", controller.MessageGetAllAfterTime).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
