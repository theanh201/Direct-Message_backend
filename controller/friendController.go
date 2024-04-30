package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"net/http"
)

func FriendGet(w http.ResponseWriter, r *http.Request) {
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
	firendList, err := model.FriendGet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(firendList)
}
