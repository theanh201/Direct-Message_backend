package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"encoding/json"
	"net/http"
	"strings"
)

// GET
func PrekeyBundleGet(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
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
	userEmail := r.FormValue("email")
	if !validMail(userEmail) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	ik, spk, opk, err := model.KeyBundleGetByEmail(userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	keyBundle := entities.PreKyeBundle{Ik: ik, Spk: spk, Opk: opk}
	json.NewEncoder(w).Encode(keyBundle)
}

// PUT
func PrekeyBundlePut(w http.ResponseWriter, r *http.Request) {
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
	// Update IK
	ik := r.FormValue("ik")
	if !valid32Byte(ik) {
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
	if !valid32Byte(ik) {
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
		if !valid32Byte(val) {
			http.Error(w, "Invalid opk", http.StatusBadRequest)
			return
		}
	}
	if len(opk) > 5 {
		http.Error(w, "Max 5 opk", http.StatusBadRequest)
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
