package model

import (
	"DirectBackend/entities"
	"database/sql"
	"encoding/hex"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Direct_Backend_DB string = "user:password1234@tcp(127.0.0.1:3306)/Direct_Backend_DB"

func AccWriteUser(username string, password string) error {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	// Add username and password to DB
	qr := fmt.Sprintf("INSERT INTO USER(USER_EMAIL, USER_PHONE_NUMB, USER_PASSWORD, USER_NAME, USER_AVATAR, USER_BACKGROUND, USER_IS_PRIVATE, USER_IS_DEL) VALUES ('%s', '', x'%s', '', '', '', 0, 0);", username, password)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	defer db.Close()
	return err
}

// return Password, Id, Error
func AccReadUserPassword(username string) (password string, id int, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return "", -1, err
	}
	err = db.Ping()
	if err != nil {
		return "", -1, err
	}
	// Read Password and ID
	qr := "SELECT USER_PASSWORD, USER_ID FROM USER WHERE USER_EMAIL='" + username + "';"
	rows, err := db.Query(qr)
	if err != nil {
		return "", -1, err
	}
	var dbPassword []byte
	rows.Next()
	if err := rows.Scan(&dbPassword, &id); err != nil {
		return "", -1, err
	}
	password = hex.EncodeToString(dbPassword)
	// Close
	defer db.Close()
	return password, id, err
}

func AccUpdateEmail(id int, email string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update Email
	qr := fmt.Sprintf("UPDATE USER SET USER_EMAIL='%s' WHERE USER_ID=%d;", email, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Revoke token
	qr = fmt.Sprintf("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=%d;", id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	defer db.Close()
	return err
}

func AccUpdatePhoneNumb(id int, phoneNumb string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update phone number
	qr := fmt.Sprintf("UPDATE USER SET USER_PHONE_NUMB='%s' WHERE USER_ID=%d;", phoneNumb, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Revoke token
	qr = fmt.Sprintf("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=%d;", id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	defer db.Close()
	return err
}

func AccUpdatePassword(id int, password string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update password
	qr := fmt.Sprintf("UPDATE USER SET USER_PASSWORD=x'%s' WHERE USER_ID=%d;", password, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Revoke token
	qr = fmt.Sprintf("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=%d;", id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	defer db.Close()
	return err
}

func AccUpdateName(id int, name string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update name
	qr := fmt.Sprintf("UPDATE USER SET USER_NAME='%s' WHERE USER_ID=%d;", name, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	defer db.Close()
	return err
}

func AccUpdateAvatar(id int, avatar string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update name
	qr := fmt.Sprintf("UPDATE USER SET USER_AVATAR='%s' WHERE USER_ID=%d;", avatar, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	defer db.Close()
	return err
}

func AccUpdateBackground(id int, avatar string) (err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	// Update name
	qr := fmt.Sprintf("UPDATE USER SET USER_BACKGROUND='%s' WHERE USER_ID=%d;", avatar, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	defer db.Close()
	return err
}

func AccGetSelf(id int) (info entities.AccountInfo, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return info, err
	}
	err = db.Ping()
	if err != nil {
		return info, err
	}
	// Get info
	qr := fmt.Sprintf("SELECT USER_EMAIL, USER_PHONE_NUMB, USER_NAME, USER_AVATAR, USER_BACKGROUND, USER_IS_PRIVATE FROM USER WHERE USER_ID=%d", id)
	row, err := db.Query(qr)
	if err != nil {
		return info, err
	}
	var temp []byte
	row.Next()
	if err := row.Scan(&info.UserEmail, &info.UserPhoneNumber, &info.UserName, &info.UserAvatar, &info.UserBackground, &temp); err != nil {
		return info, err
	}
	info.UserIsPrivate = (temp[0] & 1) != 0
	// Close
	defer db.Close()
	return info, err
}
