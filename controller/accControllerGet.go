package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	// Get info
	info, err := model.AccGetSelf(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response
	json.NewEncoder(w).Encode(info)
}
func AccGetAvatar(w http.ResponseWriter, r *http.Request) {
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
	// Get image
	imgName := r.FormValue("imgName")
	imgId, err := strconv.Atoi(strings.Split(imgName, ".")[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if id == imgId {
		http.ServeFile(w, r, fmt.Sprintf("./avatar/%s", imgName))
		return
	} else {
		info, err := model.AccGetSelf(imgId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if info.UserIsPrivate {
			http.Error(w, "User is private", http.StatusUnauthorized)
			return
		}
		http.ServeFile(w, r, fmt.Sprintf("./avatar/%s", imgName))
		return
	}
}
