package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GET
func PrekeyBundleGet(w http.ResponseWriter, r *http.Request) {
	// Validate token
	_, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// getKeyBundle
	userEmail := mux.Vars(r)["email"]
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
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Update IK
	ik, err := convert32Byte(r.FormValue("ik"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = model.KeyBundleUpdateIk(id, ik)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update SPK
	spk, err := convert32Byte(r.FormValue("spk"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = model.KeyBundleUpdateSpk(id, spk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update OPK
	opk := strings.Split(r.FormValue("opk"), ",")
	var opkByteArr [][]byte
	for _, val := range opk {
		opkByte, err := convert32Byte(val)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if err != nil {
			http.Error(w, "Invalid opk", http.StatusBadRequest)
			return
		}
		opkByteArr = append(opkByteArr, opkByte)
	}
	if len(opk) > 5 {
		http.Error(w, "Max 5 opk", http.StatusBadRequest)
	}
	err = model.KeyBundleUpdateOpk(id, opkByteArr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response
	response := map[string]string{"message": "Upload keybundle successful"}
	json.NewEncoder(w).Encode(response)
}
