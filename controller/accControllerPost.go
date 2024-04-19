package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func AccLogin(w http.ResponseWriter, r *http.Request) {
	// Decode json request
	var creds entities.Account
	var err error = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate username
	if len(creds.Username) > 64 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate password
	if len(creds.Password) != 64 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Compare password
	password, id, err := model.AccReadUserPassword(creds.Username)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	} else if password != creds.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	// Generate Token and save
	token := generateSecureRandomString(32)
	now := time.Now()
	timeout := now.AddDate(0, 0, 30).Format("2006-01-02 15:04:05")
	err = model.UserTokenAddToDB(id, token, timeout)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	// Response
	response := map[string]string{"message": "Login successful", "token": token, "timeout": timeout}
	json.NewEncoder(w).Encode(response)
}

func AccRegister(w http.ResponseWriter, r *http.Request) {
	// Decode json request
	var creds entities.Account
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate username
	if len(creds.Username) > 64 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Check username availavble
	if !validMail(creds.Username) {
		http.Error(w, "Invalid Email", http.StatusBadRequest)
		return
	}
	if len(creds.Password) != 64 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Check if user already exsist
	_, id, _ := model.AccReadUserPassword(creds.Username)
	if id != -1 {
		response := map[string]string{"message": " already exsist"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// Add user and response
	model.AccWriteUser(creds.Username, creds.Password)
	response := map[string]string{"message": "Create sucessful"}
	json.NewEncoder(w).Encode(response)
}

func generateSecureRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
