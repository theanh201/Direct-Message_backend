package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"net/http"
)

// POST
func FriendRequestPost(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}
	valid, fromId, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Get to id
	toEmail := r.FormValue("toEmail")
	if !validMail(toEmail) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	_, toId, err := model.AccGetUserPassword(toEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Get ek
	ek := r.FormValue("ek")
	if !valid32Byte(ek) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Get opk used
	opkUsed := r.FormValue("opkUsed")
	if !valid32Byte(opkUsed) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = model.FriendRequestAdd(fromId, toId, ek, opkUsed)

}

// GET
func FriendRequestGet(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
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
	friendRequest, err := model.FriendRequestGet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(friendRequest)
}
