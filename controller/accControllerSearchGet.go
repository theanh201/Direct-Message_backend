package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"net/http"
)

func AccGetUserByName(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !validToken(token) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	valid, _, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Search by name
	name := r.FormValue("name")
	if len(name) > 64 {
		http.Error(w, "len(name) not <= 64", http.StatusBadRequest)
	}
	info, err := model.AccGetUserByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Reponse
	json.NewEncoder(w).Encode(info)
}

func AccGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !validToken(token) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	valid, _, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Search by name
	email := r.FormValue("email")
	if !validMail(email) {
		http.Error(w, "Valid email not found", http.StatusBadRequest)
	}
	info, err := model.AccGetUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Reponse
	json.NewEncoder(w).Encode(info)
}
