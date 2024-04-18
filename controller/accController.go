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

func AccControllerLogin(w http.ResponseWriter, r *http.Request) {
	// Decode json request
	var creds entities.Account
	var err error = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate password
	password, id, err := model.AccModelReadUserPassword(creds.Username)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
		return
	} else if password != creds.Password {
		response := map[string]string{"message": "Invalid username or password"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// Generate Token and save
	token := generateSecureRandomString(32)
	now := time.Now()
	timeout := now.AddDate(0, 0, 30).Format("2006-01-02 15:04:05")
	err = model.UserTokenAddToDB(id, token, timeout)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
	}
	// Response
	response := map[string]string{"message": "Login successful", "token": token, "timeout": timeout}
	json.NewEncoder(w).Encode(response)
}

func AccControllerRegister(w http.ResponseWriter, r *http.Request) {
	// Decode json request
	var creds entities.Account
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Check if user already exsist
	_, id, _ := model.AccModelReadUserPassword(creds.Username)
	if id != -1 {
		response := map[string]string{"message": " already exsist"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// Add user and response
	model.AccModelWriteUser(creds.Username, creds.Password)
	response := map[string]string{"message": "Create sucessful"}
	json.NewEncoder(w).Encode(response)
}

func generateSecureRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func AccControllerUpdate(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	valid, id, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
		return
	} else if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Update email
	email := r.FormValue("email")
	if validMail(email) {
		err = model.AccModelUpdateEmail(id, email)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "No valid email found", http.StatusUnauthorized)
	}
	// Update phone number
	phoneNumb := r.FormValue("phoneNumb")
	if validPhoneNumber(phoneNumb) {
		err = model.AccModelUpdatePhoneNumb(id, phoneNumb)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "No valid phone number found", http.StatusUnauthorized)
	}
	// Update avatar and background
	imgs := []string{"avatar", "background"}
	for _, img := range imgs {
		file, fileHeader, err := r.FormFile(img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			continue
		}
		defer file.Close()

		// Create the uploads folder if it doesn't
		// already exist
		err = os.MkdirAll(img, os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a new file in the uploads directory
		dst, err := os.Create(fmt.Sprintf("./%s/%d%s", img, id, filepath.Ext(fileHeader.Filename)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the filesystem
		// at the specified destination
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	fmt.Fprintf(w, "Upload successful\n")
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
