package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Direct_Backend_DB string = "user:password1234@tcp(127.0.0.1:3306)/Direct_Backend_DB"

func WriteUserToDB(username string, password string) (error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return fmt.Errorf("Internal error", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Internal error", err)
	}

	qr := "INSERT INTO USER (USER_EMAIL, USER_PASSWORD) VALUES('" + username + "','" + password + "')"
	_, err = db.Query(qr)
	if err != nil {
		return fmt.Errorf("username %q: %v", username, err)
	}

	defer db.Close()
	return err
}

// return Password, Id, Error
func ReadUserPasswordFromDB(username string) (string, int, error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return "", -1, fmt.Errorf("Internal error", err)
	}

	err = db.Ping()
	if err != nil {
		return "", -1, fmt.Errorf("Internal error", err)
	}

	qr := "SELECT USER_PASSWORD, USER_ID FROM USER WHERE USER_EMAIL LIKE '" + username + "';"
	rows, err := db.Query(qr)
	if err != nil {
		return "", -1, fmt.Errorf("Username not exsist", err)
	}

	var dbPassword string
	var dbID int
	rows.Next()
	if err := rows.Scan(&dbPassword, &dbID); err != nil {
		return "", -1, err
	}

	defer db.Close()
	return dbPassword, dbID, err
}

func WriteUserTokenToDB(id int, token string, timeout string) error {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return fmt.Errorf("Internal error", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Internal error", err)
	}

	qr := "INSERT INTO USER_TOKEN(USER_TOKEN, USER_ID, USER_TOKEN_TIMEOUT) VALUES('" + token + "', " + fmt.Sprint(id) + ", '" + timeout + "');"
	_, err = db.Query(qr)
	if err != nil {
		return fmt.Errorf("Internal error", err)
	}

	defer db.Close()
	return err
}

// return valid, error
func CheckUserTokenWithDB(username string, token string) (bool, error) {
	return true, nil
}
