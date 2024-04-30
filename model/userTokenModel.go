package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func UserTokenAddToDB(id int, token []byte, timeout string) error {
	// Check db
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	// Add to db
	rows, err := db.Query("INSERT INTO USER_TOKEN(USER_TOKEN, USER_ID, USER_TOKEN_TIMEOUT, USER_TOKEN_IS_DEL) VALUES(?, ?, ?, 0)", token, id, timeout)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

func UserTokenValidate(token []byte) (id int, err error) {
	// Check db
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return id, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return id, err
	}
	// Query token
	rows, err := db.Query("SELECT USER_ID, USER_TOKEN_TIMEOUT FROM USER_TOKEN WHERE USER_TOKEN=? AND USER_TOKEN_IS_DEL=0", token)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	// Decode timeout
	var timeout string
	rows.Next()
	if err := rows.Scan(&id, &timeout); err != nil {
		return id, err
	}
	var layout string = "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timeout)
	if err != nil {
		return id, err
	}
	// Validate time
	now := time.Now()
	if t.Before(now) {
		return id, fmt.Errorf("token time out")
	}
	return id, err
}
