package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"net/http"
	"strings"
)

func PrekeyBundleUpdate(w http.ResponseWriter, r *http.Request) {
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
	// Update IK
	ik := r.FormValue("ik")
	if !validToken(ik) {
		http.Error(w, "Invalid ik", http.StatusBadRequest)
		return
	}
	err = model.KeyBundleUpdateIk(id, ik)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update SPK
	spk := r.FormValue("spk")
	if !validToken(ik) {
		http.Error(w, "Invalid spk", http.StatusBadRequest)
		return
	}
	err = model.KeyBundleUpdateSpk(id, spk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update OPK
	opk := strings.Split(r.FormValue("opk"), ",")
	for _, val := range opk {
		if !validToken(val) {
			http.Error(w, "Invalid spk", http.StatusBadRequest)
			return
		}
	}
	err = model.KeyBundleUpdateOpk(id, opk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response
	response := map[string]string{"message": "Upload keybundle successful"}
	json.NewEncoder(w).Encode(response)
}
