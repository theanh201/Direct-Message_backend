package model

import (
	"database/sql"
	"fmt"
)

func MessagePostFriendUnencrypt(idFrom int, idTo int, timeNow string, fileName string) (err error) {
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
	// Insert message
	qr := fmt.Sprintf("INSERT INTO MESSAGE(USER_ID_FROM, USER_ID_TO, MESSAGE_CONTENT, MESSAGE_SINCE, MESSAGE_IS_ENCRYPT) VALUES (%d,%d,'%s','%s',0)", idFrom, idTo, fileName, timeNow)
	rows, err := db.Query(qr)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}
