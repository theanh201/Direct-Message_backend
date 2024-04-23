package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func UserTokenAddToDB(id int, token string, timeout string) error {
	// Check db
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	// Add to db
	qr := fmt.Sprintf("INSERT INTO USER_TOKEN(USER_TOKEN, USER_ID, USER_TOKEN_TIMEOUT, USER_TOKEN_IS_DEL) VALUES(x'%s', %d, '%s', 0)", token, id, timeout)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	defer db.Close()
	return err
}

func UserTokenValidate(token string) (valid bool, id int, err error) {
	// Check db
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return false, -1, err
	}
	err = db.Ping()
	if err != nil {
		return false, -1, err
	}
	// Query token
	qr := fmt.Sprintf("SELECT USER_ID, USER_TOKEN_TIMEOUT FROM USER_TOKEN WHERE USER_TOKEN = x'%s' AND USER_TOKEN_IS_DEL = 0", token)
	rows, err := db.Query(qr)
	if err != nil {
		return false, -1, err
	}
	// Decode timeout
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
	// Validate time
	now := time.Now()
	if t.Before(now) {
		return false, -1, err
	}
	// Close
	defer db.Close()
	return true, id, err
}
