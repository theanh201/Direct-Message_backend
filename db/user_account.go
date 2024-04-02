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
func ReadUserFromDB(username string) (string, error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return "", err
	}

	err = db.Ping()
	if err != nil {
		return "", err
	}

	qr := "SELECT USER_ACCOUNT_PASSWORD FROM USER_ACCOUNT WHERE USER_ACCOUNT_USERNAME LIKE '" + username + "';"
	rows, err := db.Query(qr)
	if err != nil {
		return "", fmt.Errorf("username not exsist")
	}

	var dbResult string
	rows.Next()
	if err := rows.Scan(&dbResult); err != nil {
		return "", err
	}

	return dbResult, nil
}
