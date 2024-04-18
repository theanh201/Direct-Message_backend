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

func AccControllerLogin(w http.ResponseWriter, r *http.Request) {
	var creds entities.Account
	var err error = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	password, id, err := model.ReadUserPasswordFromDB(creds.Username)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
		return
	} else if password != creds.Password {
		response := map[string]string{"message": "Invalid username or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	token := generateSecureRandomString(32)
	now := time.Now()
	timeout := now.AddDate(0, 0, 30).Format("2006-01-02 15:04:05")
	err = model.UserTokenToDB(id, token, timeout)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
	}

	response := map[string]string{"message": "Login successful", "token": token, "timeout": timeout}
	json.NewEncoder(w).Encode(response)
}

func AccControllerRegister(w http.ResponseWriter, r *http.Request) {
	var creds entities.Account
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, id, _ := model.ReadUserPasswordFromDB(creds.Username)
	if id != -1 {
		response := map[string]string{"message": " already exsist"}
		json.NewEncoder(w).Encode(response)
		return
	}

	model.WriteUserToDB(creds.Username, creds.Password)
	response := map[string]string{"message": "Create sucessful"}
	json.NewEncoder(w).Encode(response)
}

func generateSecureRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
