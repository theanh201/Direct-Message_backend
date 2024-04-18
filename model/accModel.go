package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Direct_Backend_DB string = "user:password1234@tcp(127.0.0.1:3306)/Direct_Backend_DB"

func AccModelWriteUser(username string, password string) error {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	qr := "INSERT INTO USER (USER_EMAIL, USER_PASSWORD, USER_IS_PRIVATE, USER_IS_DEL) VALUES('" + username + "','" + password + "', 0, 0)"
	_, err = db.Query(qr)
	if err != nil {
		return err
	}

	defer db.Close()
	return err
}

// return Password, Id, Error
func AccModelReadUserPassword(username string) (string, int, error) {
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

func AccModelUpdateEmail(id int, email string) (err error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	qr := fmt.Sprintf("UPDATE USER SET USER_EMAIL='%s' WHERE USER_ID=%d;", email, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	qr = fmt.Sprintf("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=%d;", id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	return err
}

func AccModelUpdatePhoneNumb(id int, phoneNumb string) (err error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	qr := fmt.Sprintf("UPDATE USER SET USER_PHONE_NUMB='%s' WHERE USER_ID=%d;", phoneNumb, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	qr = fmt.Sprintf("UPDATE USER_TOKEN SET USER_TOKEN_IS_DEL=1 WHERE USER_ID=%d;", id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	return err
}
