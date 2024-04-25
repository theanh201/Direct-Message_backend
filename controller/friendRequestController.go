package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"net/http"
)

// POST
// func FriendRequestPost(w http.ResponseWriter, r *http.Request) {
// 	// Validate token
// 	token := r.FormValue("token")
// 	if !validToken(token) {
// 		http.Error(w, "Invalid token", http.StatusBadRequest)
// 		return
// 	}
// 	valid, fromId, err := model.UserTokenValidate(token)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	} else if !valid {
// 		http.Error(w, "Token expired", http.StatusUnauthorized)
// 		return
// 	}
// 	toEmail := r.FormValue("requestToEmail")
// 	if !validMail(toEmail) {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}
// }

// GET
func FriendRequestGet(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !validToken(token) {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}
	valid, id, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Get friend request that not rejected
	friendRequest, err := model.GetFriendRequest(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(friendRequest)
}
