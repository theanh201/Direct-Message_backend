package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// POST
func AccPostLogin(w http.ResponseWriter, r *http.Request) {
	// Decode json request
	var creds entities.Account
	var err error = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate username
	if !validMail(creds.Email) {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	// Compare password
	password, id, err := model.AccGetUserPassword(creds.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if password != creds.Password {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}
	// Generate Token and save
	token := generateSecureRandomString(32)
	tokenByte, err := hex.DecodeString(token)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	now := time.Now()
	timeout := now.AddDate(0, 0, 30).Format("2006-01-02 15:04:05")
	err = model.UserTokenAddToDB(id, tokenByte, timeout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Response
	response := map[string]string{"message": "login successful", "token": token, "timeout": timeout}
	json.NewEncoder(w).Encode(response)
}

func AccPostRegister(w http.ResponseWriter, r *http.Request) {
	// Decode json request
	var creds entities.Account
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate username
	if !validMail(creds.Email) {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	// Check if user already exsist
	_, id, _ := model.AccGetUserPassword(creds.Email)
	if id != -1 {
		http.Error(w, "username already exsist", http.StatusBadRequest)
		return
	}
	// Validate password
	password, err := convert32Byte(creds.Password)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	// Add user and response
	err = model.AccAddUser(creds.Email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "create sucessful"}
	json.NewEncoder(w).Encode(response)
}

// GET
func AccGetSelfInfo(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	id, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get image if imgId = tokenId
	imgName := mux.Vars(r)["imgName"]
	imgId, err := strconv.Atoi(strings.Split(imgName, ".")[0])
	if err != nil {
		http.Error(w, "fail to extract id from image name", http.StatusBadRequest)
	}
	if id == imgId {
		http.ServeFile(w, r, fmt.Sprintf("./Avatar/%s", imgName))
		return
	}
	// Get imgage if imgId != private
	info, err := model.AccGetInfo(imgId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if info.IsPrivate {
		http.Error(w, "user is private", http.StatusUnauthorized)
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("./Avatar/%s", imgName))
}

func AccGetUserByName(w http.ResponseWriter, r *http.Request) {
	// Validate token
	_, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get page
	page, err := strconv.Atoi(mux.Vars(r)["page"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if page < 0 {
		http.Error(w, "page < 0", http.StatusBadRequest)
		return
	}
	// Search by name
	name := mux.Vars(r)["name"]
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
	_, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Search by name
	email := mux.Vars(r)["email"]
	if !validMail(email) {
		http.Error(w, "valid email not found", http.StatusBadRequest)
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
	id, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get image if imgId = tokenId
	imgName := mux.Vars(r)["imgName"]
	imgId, err := strconv.Atoi(strings.Split(imgName, ".")[0])
	if err != nil {
		http.Error(w, "fail to extract id from image", http.StatusBadRequest)
	}
	if id == imgId {
		http.ServeFile(w, r, fmt.Sprintf("./Background/%s", imgName))
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
	http.ServeFile(w, r, fmt.Sprintf("./Background/%s", imgName))
}

// PUT
func AccPutAvatar(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	err = os.MkdirAll("Avatar", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a new file in the uploads directory
	fileName := fmt.Sprintf("%d%s", id, filepath.Ext(fileHeader.Filename))
	dst, err := os.Create(fmt.Sprintf("./%s/%s", "Avatar", fileName))
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
	// Validate token
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	err = os.MkdirAll("Background", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Create a new file in the uploads directory
	fileName := fmt.Sprintf("%d%s", id, filepath.Ext(fileHeader.Filename))
	dst, err := os.Create(fmt.Sprintf("./%s/%s", "Background", fileName))
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
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Update Password
	password := r.FormValue("password")
	if len(password) == 64 {
		pw, err := hex.DecodeString(password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = model.AccUpdatePassword(id, pw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Fail", http.StatusBadRequest)
		return
	}
	// Reponse
	response := map[string]string{"message": "Password have been updated"}
	json.NewEncoder(w).Encode(response)
}

func AccPutName(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	id, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Update private status
	status := r.FormValue("isPrivate")
	if status == "0" || status == "1" {
		st, _ := strconv.Atoi(status)
		err = model.AccUpdatePrivateStatus(id, st)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "status not 0 or 1", http.StatusBadRequest)
		return
	}
	// Reponse
	response := map[string]string{"message": "Private status have been updated"}
	json.NewEncoder(w).Encode(response)

}

// DELETE
func AccDeleteSelf(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id1, err := validateToken(r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
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
