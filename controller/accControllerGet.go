package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"net/http"
)

func AccGetSelfInfo(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !validToken(token) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
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
	info, err := model.AccGetSelf(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response
	json.NewEncoder(w).Encode(info)
}
