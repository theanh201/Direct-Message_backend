package model

import (
	"DirectBackend/entities"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var Direct_Backend_DB string = "root:Python5979@tcp(localhost:3306)/CHATDB"

// Add
func AccAddUser(email string, password []byte) error {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	// Add username and password to DB
	rows, err := db.Query("INSERT INTO USER(USER_EMAIL, USER_PASSWORD, USER_NAME, USER_AVATAR, USER_BACKGROUND, USER_IS_PRIVATE, USER_IS_DEL) VALUES (?, ?, ?, '', '', 0, 0)", email, password, strings.Split(email, "@")[0])
	if err != nil {
		return err
	}
	defer rows.Close()
	// Read  ID
	_, id, err := AccGetUserPassword(email)
	if err != nil {
		return err
	}
	// Create empty ik spk
	rows, err = db.Query("INSERT INTO USER_KEY(USER_ID, USER_KEY_IK, USER_KEY_SPK) VALUES (?, '', '')", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	// Create 5 opk
	for i := 0; i < 5; i++ {
		temp := fmt.Sprintf("%d", i)
		if err != nil {
			return err
		}
		rows, err = db.Query("INSERT INTO USER_OPK_KEY(USER_ID, USER_OPK_KEY, USER_OPK_KEY_IS_DEL) VALUES (?, ?, 0)", id, temp)
		if err != nil {
			return err
		}
		defer rows.Close()
	}
	return err
}

// Get
func AccGetUserPassword(email string) (password string, id int, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return "", -1, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return "", -1, err
	}
	// Read Password and ID
	rows, err := db.Query("SELECT USER_PASSWORD, USER_ID FROM USER WHERE USER_EMAIL=? AND USER_IS_DEL = 0", email)
	if err != nil {
		return "", -1, err
	}
	defer rows.Close()
	var dbPassword []byte
	rows.Next()
	if err := rows.Scan(&dbPassword, &id); err != nil {
		return "", -1, err
	}
	password = hex.EncodeToString(dbPassword)
	return password, id, err
}

func AccGetInfo(id int) (info entities.AccountInfo, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return info, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return info, err
	}
	// Get info
	rows, err := db.Query("SELECT USER_EMAIL, USER_NAME, USER_AVATAR, USER_BACKGROUND, USER_IS_PRIVATE FROM USER WHERE USER_ID=?", id)
	if err != nil {
		return info, err
	}
	defer rows.Close()
	var temp []byte
	rows.Next()
	if err := rows.Scan(&info.Email, &info.Name, &info.Avatar, &info.Background, &temp); err != nil {
		return info, err
	}
	info.IsPrivate = (temp[0] & 1) != 0
	return info, err
}

func AccGetByName(name string, page int) (result []entities.AccountInfoExcludePrivateStatus, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return result, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return result, err
	}
	// Get info
	name = "%" + name + "%"
	page *= 10
	// Query
	rows, err := db.Query("SELECT USER_EMAIL, USER_NAME, USER_AVATAR, USER_BACKGROUND FROM USER WHERE USER_NAME LIKE ? AND USER_IS_PRIVATE = 0 AND USER_IS_DEL = 0 LIMIT 10 OFFSET ?", name, page)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	var info entities.AccountInfoExcludePrivateStatus
	for rows.Next() {
		if err := rows.Scan(&info.Email, &info.Name, &info.Avatar, &info.Background); err != nil {
			return result, err
		}
		result = append(result, info)
	}
	return result, err
}

func AccGetByEmail(email string) (info entities.AccountInfoExcludePrivateStatus, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return info, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return info, err
	}
	// Get info
	rows, err := db.Query("SELECT USER_EMAIL, USER_NAME, USER_AVATAR, USER_BACKGROUND FROM USER WHERE USER_EMAIL LIKE ? AND USER_IS_DEL = 0", email)
	if err != nil {
		return info, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&info.Email, &info.Name, &info.Avatar, &info.Background); err != nil {
			return info, err
		}
	}
	return info, err
}

// Update
func AccUpdateEmail(id int, email string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update Email
	rows, err := db.Query("UPDATE USER SET USER_EMAIL=? WHERE USER_ID=?", email, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	// Revoke token
	rows, err = db.Query("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

func AccUpdatePassword(id int, password []byte) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update password
	rows, err := db.Query("UPDATE USER SET USER_PASSWORD=? WHERE USER_ID=?", password, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	// Revoke token
	rows, err = db.Query("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

func AccUpdateName(id int, name string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update name
	rows, err := db.Query("UPDATE USER SET USER_NAME=? WHERE USER_ID=?", name, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

func AccUpdateAvatar(id int, avatar string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update avatar
	rows, err := db.Query("UPDATE USER SET USER_AVATAR=? WHERE USER_ID=?", avatar, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

func AccUpdateBackground(id int, avatar string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update background
	rows, err := db.Query("UPDATE USER SET USER_BACKGROUND=? WHERE USER_ID=?", avatar, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

func AccUpdatePrivateStatus(id int, status int) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update private status
	rows, err := db.Query("UPDATE USER SET USER_IS_PRIVATE=? WHERE USER_ID=?", status, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

// Delete
func AccDelete(id int) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update delete status
	rows, err := db.Query("UPDATE USER SET USER_IS_DEL=1 WHERE USER_ID=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	// Revoke token
	rows, err = db.Query("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}
