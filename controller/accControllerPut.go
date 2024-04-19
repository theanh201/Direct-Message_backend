package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const MAX_REQUEST_IMG = 10 << 20 // 10 MB

func AccUpdateAvatar(w http.ResponseWriter, r *http.Request) {
	// Limit request size
	r.Body = http.MaxBytesReader(w, r.Body, MAX_REQUEST_IMG)
	err := r.ParseMultipartForm(MAX_REQUEST_IMG)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
	// Update avatar
	file, fileHeader, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Create the uploads folder if it doesn't already exist
	err = os.MkdirAll("avatar", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a new file in the uploads directory
	fileName := fmt.Sprintf("%d%s", id, filepath.Ext(fileHeader.Filename))
	dst, err := os.Create(fmt.Sprintf("./%s/%s", "avatar", fileName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	// Update path in DB
	err = model.AccUpdateAvatar(id, fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response
	response := map[string]string{"message": "Avatar have been updated"}
	json.NewEncoder(w).Encode(response)
}

func AccUpdateBackground(w http.ResponseWriter, r *http.Request) {
	// Limit request size
	r.Body = http.MaxBytesReader(w, r.Body, MAX_REQUEST_IMG)
	err := r.ParseMultipartForm(MAX_REQUEST_IMG)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
	// Update background
	file, fileHeader, err := r.FormFile("background")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Create the uploads folder if it doesn't already exist
	err = os.MkdirAll("background", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a new file in the uploads directory
	fileName := fmt.Sprintf("%d%s", id, filepath.Ext(fileHeader.Filename))
	dst, err := os.Create(fmt.Sprintf("./%s/%s", "background", fileName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	// Update path in DB
	err = model.AccUpdateBackground(id, fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response
	response := map[string]string{"message": "Background have been updated"}
	json.NewEncoder(w).Encode(response)
}

func AccUpdateEmail(w http.ResponseWriter, r *http.Request) {
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
	// Update email
	email := r.FormValue("email")
	if validMail(email) {
		err = model.AccUpdateEmail(id, email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Not valid email found", http.StatusBadRequest)
		return
	}
	// Response
	response := map[string]string{"message": "Email address have been updated"}
	json.NewEncoder(w).Encode(response)
}

func AccUpdatePhoneNumber(w http.ResponseWriter, r *http.Request) {
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
	// Update phone number
	phoneNumb := r.FormValue("phoneNumb")
	if validPhoneNumber(phoneNumb) {
		err = model.AccUpdatePhoneNumb(id, phoneNumb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "No valid phone number found", http.StatusUnauthorized)
		return
	}
	// Reponse
	response := map[string]string{"message": "Phone number have been updated"}
	json.NewEncoder(w).Encode(response)
}
func AccUpdatePassword(w http.ResponseWriter, r *http.Request) {
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
	// Update Password
	password := r.FormValue("password")
	if len(password) == 64 {
		err = model.AccUpdatePassword(id, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	// Reponse
	response := map[string]string{"message": "Password have been updated"}
	json.NewEncoder(w).Encode(response)
}
func AccUpdateName(w http.ResponseWriter, r *http.Request) {
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
	// Update name
	name := r.FormValue("name")
	if len(name) <= 64 {
		err = model.AccUpdateName(id, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	// Reponse
	response := map[string]string{"message": "Name have been updated"}
	json.NewEncoder(w).Encode(response)
}
