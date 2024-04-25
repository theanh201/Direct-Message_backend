package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"encoding/json"
	"net/http"
)

func GetPrekeyBundle(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !validToken(token) {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}
	valid, _, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// getKeyBundle
	userEmail := r.FormValue("user")
	if !validMail(userEmail) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	ik, spk, opk, err := model.KeyBundleGetByEmail(userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	keyBundle := entities.PreKyeBundle{Ik: ik, Spk: spk, Opk: opk}
	json.NewEncoder(w).Encode(keyBundle)
}
