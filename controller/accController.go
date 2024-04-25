package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const MAX_REQUEST_IMG = 10 << 20 // 10 MB

// POST
func AccPostLogin(w http.ResponseWriter, r *http.Request) {
	// Decode json request
	var creds entities.Account
	var err error = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate username
	if !validMail(creds.Email) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate password
	if !valid32Byte(creds.Password) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Compare password
	password, id, err := model.AccGetUserPassword(creds.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response
	response := map[string]string{"message": "Login successful", "token": token, "timeout": timeout}
	json.NewEncoder(w).Encode(response)
}

func AccPostRegister(w http.ResponseWriter, r *http.Request) {
	// Decode json request
	var creds entities.Account
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate username
	if !validMail(creds.Email) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate password
	if !valid32Byte(creds.Password) {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Check if user already exsist
	_, id, _ := model.AccGetUserPassword(creds.Email)
	if id != -1 {
		response := map[string]string{"message": "Username already exsist"}
		json.NewEncoder(w).Encode(response)
		return
	}
	// Add user and response
	err = model.AccAddUser(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Create sucessful"}
	json.NewEncoder(w).Encode(response)
}

// GET
func AccGetSelfInfo(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
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
	info, err := model.AccGetInfo(id)
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
	if !valid32Byte(token) {
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
	// Get image if imgId = tokenId
	imgName := r.FormValue("imgName")
	imgId, err := strconv.Atoi(strings.Split(imgName, ".")[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if id == imgId {
		http.ServeFile(w, r, fmt.Sprintf("./avatar/%s", imgName))
		return
	}
	// Get imgage if imgId != private
	info, err := model.AccGetInfo(imgId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if info.IsPrivate {
		http.Error(w, "User is private", http.StatusUnauthorized)
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("./avatar/%s", imgName))
}

func AccGetUserByName(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	valid, _, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Get page
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if page < 0 {
		http.Error(w, "page < 0", http.StatusBadRequest)
		return
	}
	// Search by name
	name := r.FormValue("name")
	if len(name) > 64 {
		http.Error(w, "len(name) not <= 64", http.StatusBadRequest)
	}
	info, err := model.AccGetByName(name, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Reponse
	json.NewEncoder(w).Encode(info)
}

func AccGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	valid, _, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Search by name
	email := r.FormValue("email")
	if !validMail(email) {
		http.Error(w, "Valid email not found", http.StatusBadRequest)
	}
	info, err := model.AccGetByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Reponse
	json.NewEncoder(w).Encode(info)
}

func AccGetBackGround(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
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
	// Get image if imgId = tokenId
	imgName := r.FormValue("imgName")
	imgId, err := strconv.Atoi(strings.Split(imgName, ".")[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if id == imgId {
		http.ServeFile(w, r, fmt.Sprintf("./background/%s", imgName))
		return
	}
	// Get imgage if imgId != private
	info, err := model.AccGetInfo(imgId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if info.IsPrivate {
		http.Error(w, "User is private", http.StatusUnauthorized)
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("./background/%s", imgName))
}

// PUT
func AccPutAvatar(w http.ResponseWriter, r *http.Request) {
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

func AccPutBackground(w http.ResponseWriter, r *http.Request) {
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

func AccPutEmail(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
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

func AccPutPassword(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
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

func AccPutName(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
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
		http.Error(w, "len(name) > 64", http.StatusBadRequest)
		return
	}
	// Reponse
	response := map[string]string{"message": "Name have been updated"}
	json.NewEncoder(w).Encode(response)
}

func AccPutPrivateStatus(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
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
	// Update private status
	status := r.FormValue("isPrivate")
	if len(status) == 1 {
		err = model.AccUpdatePrivateStatus(id, status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "len(isPrivate) != 1", http.StatusBadRequest)
		return
	}
	// Reponse
	response := map[string]string{"message": "Private status have been updated"}
	json.NewEncoder(w).Encode(response)

}

// DELETE
func AccDeleteSelf(w http.ResponseWriter, r *http.Request) {
	// Validate token
	token := r.FormValue("token")
	if !valid32Byte(token) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	valid, id1, err := model.UserTokenValidate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}
	// Validate email
	email := r.FormValue("email")
	if !validMail(email) {
		http.Error(w, "Not valid email found", http.StatusBadRequest)
		return
	}
	// Validate Password
	password := r.FormValue("password")
	if len(password) != 64 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	// Compare password
	pwd, id2, err := model.AccGetUserPassword(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if password != pwd {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	if id1 != id2 {
		http.Error(w, "Try logout and login again", http.StatusUnauthorized)
		return
	}
	err = model.AccDelete(id1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Your account have been deleted"}
	json.NewEncoder(w).Encode(response)
}
