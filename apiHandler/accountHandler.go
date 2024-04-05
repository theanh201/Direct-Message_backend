package apiHandler

import (
	"DirectBackend/api"
	"DirectBackend/db"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds api.User
	var err error = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	password, id, err := db.ReadUserPasswordFromDB(creds.Username)
	if err != nil || password != creds.Password {
		response := map[string]string{"message": "Invalid username or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	token := generateSecureRandomString(32)
	now := time.Now()
	fmt.Println(now.Format(""))
	var timeout string = "2024-4-2  00:00:00"
	err = db.WriteUserTokenToDB(id, token, timeout)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
	}

	response := map[string]string{"message": "Login successful", "token": token, "timeout": timeout}
	json.NewEncoder(w).Encode(response)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds api.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	_, _, err = db.ReadUserPasswordFromDB(creds.Username)
	if err == nil {
		response := map[string]string{"message": "Username already exsist"}
		json.NewEncoder(w).Encode(response)
		return
	}
	db.WriteUserToDB(creds.Username, creds.Password)
	response := map[string]string{"message": "Create sucessful"}
	json.NewEncoder(w).Encode(response)
}

func generateSecureRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
