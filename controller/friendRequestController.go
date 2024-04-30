package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"fmt"
	"net/http"
)

// POST
func FriendRequestPost(w http.ResponseWriter, r *http.Request) {
	// Validate token
	fromId, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	ek, err := convert32Byte(r.FormValue("ek"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get opk used
	opkUsed, err := convert32Byte(r.FormValue("opkUsed"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = model.FriendRequestAdd(fromId, toId, ek, opkUsed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Add sucessful"}
	json.NewEncoder(w).Encode(response)
}
func FriendRequestPostAccept(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate email
	email := r.FormValue("email")
	if !validMail(email) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Delete friend request
	err = model.FriendRequestUpdateReject(email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Accect friend request and create group
	err = model.FriendAdd(email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message := fmt.Sprintf("You are now friend with %s", email)
	response := map[string]string{"message": message}
	json.NewEncoder(w).Encode(response)
}
func FriendRequestPostReject(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate email
	email := r.FormValue("email")
	if !validMail(email) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Delete friend request
	err = model.FriendRequestUpdateReject(email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message := fmt.Sprintf("You have delete request from %s", email)
	response := map[string]string{"message": message}
	json.NewEncoder(w).Encode(response)
}

// GET
func FriendRequestGet(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
