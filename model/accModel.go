package model

import (
	"DirectBackend/entities"
	"database/sql"
	"encoding/hex"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Direct_Backend_DB string = "user:password1234@tcp(127.0.0.1:3306)/Direct_Backend_DB"

// Add
func AccAddUser(email string, password string) error {
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
	qr := fmt.Sprintf("INSERT INTO USER(USER_EMAIL, USER_PASSWORD, USER_NAME, USER_AVATAR, USER_BACKGROUND, USER_IS_PRIVATE, USER_IS_DEL) VALUES ('%s', x'%s', '%s', '', '', 0, 0)", email, password, email)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Read  ID
	qr = fmt.Sprintf("SELECT USER_ID FROM USER WHERE USER_EMAIL='%s' AND USER_IS_DEL = 0", email)
	rows, err := db.Query(qr)
	if err != nil {
		return err
	}
	defer rows.Close()
	var id int
	rows.Next()
	if err := rows.Scan(&id); err != nil {
		return err
	}
	// Create empty ik spk
	qr = fmt.Sprintf("INSERT INTO USER_KEY(USER_ID, USER_KEY_IK, USER_KEY_SPK) VALUES (%d, '', '')", id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Create 5 opk
	for i := 0; i < 5; i++ {
		qr = fmt.Sprintf("INSERT INTO USER_OPK_KEY(USER_ID, USER_OPK_KEY, USER_OPK_KEY_IS_DEL) VALUES (%d, '%d', 0)", id, i)
		_, err = db.Query(qr)
		if err != nil {
			return err
		}
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
	qr := fmt.Sprintf("SELECT USER_PASSWORD, USER_ID FROM USER WHERE USER_EMAIL='%s' AND USER_IS_DEL = 0", email)
	rows, err := db.Query(qr)
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
	qr := fmt.Sprintf("SELECT USER_EMAIL, USER_NAME, USER_AVATAR, USER_BACKGROUND, USER_IS_PRIVATE FROM USER WHERE USER_ID=%d", id)
	row, err := db.Query(qr)
	if err != nil {
		return info, err
	}
	defer row.Close()
	var temp []byte
	row.Next()
	if err := row.Scan(&info.Email, &info.Name, &info.Avatar, &info.Background, &temp); err != nil {
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
	name = name + "%"
	page *= 10
	qr := fmt.Sprintf("SELECT USER_EMAIL, USER_NAME, USER_AVATAR, USER_BACKGROUND FROM USER WHERE USER_NAME LIKE '%s' AND USER_IS_PRIVATE = 0 AND USER_IS_DEL = 0 LIMIT 10 OFFSET %d", name, page)
	row, err := db.Query(qr)
	if err != nil {
		return result, err
	}
	defer row.Close()
	var info entities.AccountInfoExcludePrivateStatus
	for row.Next() {
		if err := row.Scan(&info.Email, &info.Name, &info.Avatar, &info.Background); err != nil {
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
	qr := fmt.Sprintf("SELECT USER_EMAIL, USER_NAME, USER_AVATAR, USER_BACKGROUND FROM USER WHERE USER_EMAIL LIKE '%s' AND USER_IS_DEL = 0", email)
	row, err := db.Query(qr)
	if err != nil {
		return info, err
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&info.Email, &info.Name, &info.Avatar, &info.Background); err != nil {
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
	qr := fmt.Sprintf("UPDATE USER SET USER_EMAIL='%s' WHERE USER_ID=%d", email, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Revoke token
	qr = fmt.Sprintf("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=%d", id)
	_, err = db.Query(qr)
	return err
}

func AccUpdatePassword(id int, password string) (err error) {
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
	qr := fmt.Sprintf("UPDATE USER SET USER_PASSWORD=x'%s' WHERE USER_ID=%d", password, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Revoke token
	qr = fmt.Sprintf("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=%d", id)
	_, err = db.Query(qr)
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
	qr := fmt.Sprintf("UPDATE USER SET USER_NAME='%s' WHERE USER_ID=%d", name, id)
	_, err = db.Query(qr)
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
	qr := fmt.Sprintf("UPDATE USER SET USER_AVATAR='%s' WHERE USER_ID=%d", avatar, id)
	_, err = db.Query(qr)
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
	qr := fmt.Sprintf("UPDATE USER SET USER_BACKGROUND='%s' WHERE USER_ID=%d", avatar, id)
	_, err = db.Query(qr)
	return err
}

func AccUpdatePrivateStatus(id int, status string) (err error) {
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
	qr := fmt.Sprintf("UPDATE USER SET USER_IS_PRIVATE=%s WHERE USER_ID=%d", status, id)
	_, err = db.Query(qr)
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
	qr := fmt.Sprintf("UPDATE USER SET USER_IS_DEL=1 WHERE USER_ID=%d", id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Revoke token
	qr = fmt.Sprintf("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=%d", id)
	_, err = db.Query(qr)
	return err
}
