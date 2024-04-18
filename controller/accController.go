package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"
	"time"
	"unicode"
)

// 50 mb limit
const MAX_UPLOAD_SIZE = 1024 * 1024

func AccLogin(w http.ResponseWriter, r *http.Request) {
	// Decode json request
	var creds entities.Account
	var err error = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate password
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
	if !validMail(creds.Username) {
		http.Error(w, "Invalid Email", http.StatusBadRequest)
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

func AccUpdateAvatar(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	valid, id, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	} else if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Update avatar
	file, fileHeader, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()
	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("avatar", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./%s/%d%s", "avatar", id, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
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
	// Validate token
	token := r.FormValue("token")
	valid, id, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
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
	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("background", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./%s/%d%s", "background", id, filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
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
	valid, id, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
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
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "No valid email found", http.StatusBadRequest)
		return
	}
	// Response
	response := map[string]string{"message": "Email address have been updated"}
	json.NewEncoder(w).Encode(response)
}
func AccUpdatePhoneNumber(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	valid, id, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
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
			http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
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
func validMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func validPhoneNumber(s string) bool {
	if len(s) != 10 {
		return false
	}
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
func AccUpdatePassword(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	valid, id, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
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
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
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
	valid, id, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
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
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
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
