package main

import (
	"DirectBackend/api"
	"DirectBackend/db"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Hardcoded user credentials for demonstration purposes
var UserAccount = map[string]string{
	"user1": "password1",
	"user2": "password2",
}
var UserTokens = map[string]api.Token{}

func main() {
	// db.WriteUserToDB("user3", "password3")
	password, err := db.ReadUserFromDB("user3")
	fmt.Println(password, err)
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/register", registerHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds api.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	password, exist := UserAccount[creds.Username]
	if !exist || password != creds.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token := generateSecureRandomString(64)
	tokens, exist := UserTokens[creds.Username]

	if !exist {
		UserTokens[creds.Username] = api.Token{
			Tokens:        []string{token},
			TokensTimeOut: []int{30},
		}
	} else {
		UserTokens[creds.Username] = api.Token{
			Tokens:        append(tokens.Tokens, token),
			TokensTimeOut: append(tokens.TokensTimeOut, 30),
		}
	}
	response := map[string]string{"message": "Login successful", "token": token, "timeout": "30"}
	json.NewEncoder(w).Encode(response)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var creds api.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	_, exist := UserAccount[creds.Username]
	if exist {
		http.Error(w, "Username already exsist", http.StatusUnauthorized)
		return
	}
	UserAccount[creds.Username] = creds.Password
	response := map[string]string{"message": "Create sucess"}
	json.NewEncoder(w).Encode(response)
}

func generateSecureRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
