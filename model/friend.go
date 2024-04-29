package model

import (
	"database/sql"
	"fmt"
	"time"
)

// add
func FriendAdd(email string, id int) (err error) {
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
	now := time.Now().Format("2006-01-02 15:04:05")
	_, id2, err := AccGetUserPassword(email)
	if err != nil {
		return err
	}
	if id > id2 {
		id, id2 = id2, id
	}
	if id == id2 {
		return fmt.Errorf("you cannot make friend with your self")
	}
	qr := fmt.Sprintf("INSERT INTO USER_FRIEND (USER_ID_1, USER_ID_2, USER_FRIEND_SINCE, USER_FRIEND_IS_DEL) VALUES(%d, %d, '%s', 0)", id, id2, now)
	_, err = db.Query(qr)
	return err
}
