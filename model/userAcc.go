package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Direct_Backend_DB string = "user:password1234@tcp(127.0.0.1:3306)/Direct_Backend_DB"

func WriteUserToDB(username string, password string) (error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	qr := "INSERT INTO USER (USER_EMAIL, USER_PASSWORD, USER_IS_DEL) VALUES('" + username + "','" + password + "', 0)"
	_, err = db.Query(qr)
	if err != nil {
		return err
	}

	defer db.Close()
	return err
}

// return Password, Id, Error
func ReadUserPasswordFromDB(username string) (string, int, error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return "", -1, err
	}

	err = db.Ping()
	if err != nil {
		return "", -1, err
	}

	qr := "SELECT USER_PASSWORD, USER_ID FROM USER WHERE USER_EMAIL LIKE '" + username + "';"
	rows, err := db.Query(qr)
	if err != nil {
		return "", -1, err
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

