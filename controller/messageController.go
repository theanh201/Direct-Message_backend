package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func MessagePostFriendUnencrypt(w http.ResponseWriter, r *http.Request) {
	// Limit request size
	r.Body = http.MaxBytesReader(w, r.Body, MAX_REQUEST_IMG)
	err := r.ParseMultipartForm(MAX_REQUEST_IMG)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}
	valid, idFrom, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	// Get email
	email := r.FormValue("email")
	if !validMail(email) {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}
	_, idTo, err := model.AccGetUserPassword(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// Check if 2 are friend
	err = model.FriendCheck(idFrom, idTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Get Content
	file, fileHeader, err := r.FormFile("content")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Create the uploads folder if it doesn't already exist
	err = os.MkdirAll("Message", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a new file in the uploads directory
	fileName := fmt.Sprintf("%d_%s%s", idFrom, timeNow, filepath.Ext(fileHeader.Filename))
	dst, err := os.Create(fmt.Sprintf("./%s/%s", "Message", fileName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	// Upload to db
	err = model.MessagePostFriendUnencrypt(idFrom, idTo, timeNow, fileName)
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
	response := map[string]string{"message": "Message sent"}
	json.NewEncoder(w).Encode(response)
}

func MessageGetAll(w http.ResponseWriter, r *http.Request) {
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
	// Get message
	messages, err := model.MessageGetAll(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

func MessageGetContent(w http.ResponseWriter, r *http.Request) {
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
	// Get content file name
	contentName := r.FormValue("content")
	idFrom, idTo, err := model.MessageGetContentPermission(contentName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if id != idFrom && id != idTo {
		http.Error(w, "You are not allow to access this information", http.StatusUnauthorized)
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("./Message/%s", contentName))
}
