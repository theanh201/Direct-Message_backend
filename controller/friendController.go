package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func FriendGet(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get friend list
	firendList, err := model.FriendGet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(firendList)
}
func FriendDelete(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id1, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get email
	email := mux.Vars(r)["email"]
	if !validMail(email) {
		http.Error(w, "valid email not found", http.StatusBadRequest)
		return
	}
	_, id2, err := model.AccGetUserPassword(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = model.FriendRemove(id1, id2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": fmt.Sprintf("You have unfriend %s", email)}
	json.NewEncoder(w).Encode(response)
}
