package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Direct_Backend_DB string = "user:password1234@tcp(127.0.0.1:3306)/Direct_Backend_DB"

func WriteUserToDB(username string, password string) (int, error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return -1, err
	}

	err = db.Ping()
	if err != nil {
		return -1, err
	}

	qr := "INSERT INTO USER_ACCOUNT (USER_ACCOUNT_USERNAME, USER_ACCOUNT_PASSWORD) VALUES('" + username + "','" + password + "')"
	_, err = db.Query(qr)
	if err != nil {
		return -1, fmt.Errorf("username %q: %v", username, err)
	}

	defer db.Close()
	return 0, nil
}

// return Password, Error
func ReadUserPasswordFromDB(username string) (string, int, error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return "", 0, err
	}

	err = db.Ping()
	if err != nil {
		return "", 0, err
	}

	qr := "SELECT USER_ACCOUNT_PASSWORD, USER_ACCOUNT_ID FROM USER_ACCOUNT WHERE USER_ACCOUNT_USERNAME LIKE '" + username + "';"
	rows, err := db.Query(qr)
	if err != nil {
		return "", 0, fmt.Errorf("username not exsist")
	}

	var dbPassword string
	var dbID int
	rows.Next()
	if err := rows.Scan(&dbPassword, &dbID); err != nil {
		return "", 0, err
	}

	return dbPassword, dbID, nil
}

func WriteUserTokenToDB(id int, token string, timeout string) error {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	qr := "INSERT INTO USER_TOKEN(USER_TOKEN_ID, USER_ACCOUNT_ID, USER_TOKEN_EXPIRE_DATE) VALUES('" + token + "', " + fmt.Sprint(id) + ", '" + timeout + "');"
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	return nil
}

// return valid, error
func CheckUserTokenWithDB(username string, token string) (bool, error) {
	return true, nil
}
