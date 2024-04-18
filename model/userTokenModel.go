package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func UserTokenAddToDB(id int, token string, timeout string) error {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	qr := "INSERT INTO USER_TOKEN(USER_TOKEN, USER_ID, USER_TOKEN_TIMEOUT, USER_TOKEN_IS_DEL) VALUES(x'" + token + "', " + fmt.Sprint(id) + ", '" + timeout + "', 0);"
	_, err = db.Query(qr)
	if err != nil {
		return err
	}

	defer db.Close()
	return err
}

func UserTokenValidate(token string) (valid bool, id int, err error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return false, -1, err
	}

	err = db.Ping()
	if err != nil {
		return false, -1, err
	}

	qr := "SELECT USER_ID, USER_TOKEN_TIMEOUT FROM USER_TOKEN WHERE USER_TOKEN = x'" + token + "' AND USER_TOKEN_IS_DEL = 0;"
	rows, err := db.Query(qr)
	if err != nil {
		return false, -1, err
	}

	var timeout string
	rows.Next()
	if err := rows.Scan(&id, &timeout); err != nil {
		return false, -1, err
	}
	var layout string = "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timeout)
	if err != nil {
		return false, -1, err
	}

	now := time.Now()
	if t.Before(now) {
		return false, -1, err
	}

	defer db.Close()
	return true, id, err
}
