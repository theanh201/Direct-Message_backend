package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"net/http"
)

func DeleteSelf(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !validToken(token) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	valid, id1, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Validate email
	email := r.FormValue("email")
	if !validMail(email) {
		http.Error(w, "Not valid email found", http.StatusBadRequest)
		return
	}
	// Validate Password
	password := r.FormValue("password")
	if len(password) != 64 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	// Compare password
	pwd, id2, err := model.AccReadUserPassword(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if password != pwd {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	if id1 != id2 {
		http.Error(w, "Try logout and login again", http.StatusUnauthorized)
		return
	}
	err = model.AccDelete(id1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Your account have been deleted"}
	json.NewEncoder(w).Encode(response)
}
