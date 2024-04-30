package model

import (
	"DirectBackend/entities"
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
	rows, err := db.Query("INSERT INTO USER_FRIEND (USER_ID_1, USER_ID_2, USER_FRIEND_SINCE, USER_FRIEND_IS_DEL) VALUES(?, ?, ?, 0)", id, id2, now)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

// Get
func FriendGet(id int) (friendList []entities.Friend, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return friendList, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return friendList, err
	}
	// columm 1 = id
	rows, err := db.Query("SELECT USER_ID_2, USER_FRIEND_SINCE FROM USER_FRIEND WHERE USER_ID_1=? AND USER_FRIEND_IS_DEL=0", id)
	if err != nil {
		return friendList, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempId int
		var tempDate string
		if err := rows.Scan(&tempId, &tempDate); err != nil {
			return friendList, err
		}
		info, err := AccGetInfo(tempId)
		if err != nil {
			return friendList, err
		}
		friendList = append(friendList, entities.Friend{
			Info: entities.AccountInfoExcludePrivateStatus{
				Email:      info.Email,
				Name:       info.Name,
				Avatar:     info.Avatar,
				Background: info.Background,
			},
			Since: tempDate,
		})
	}
	// columm 2 = id
	rows, err = db.Query("SELECT USER_ID_1, USER_FRIEND_SINCE FROM USER_FRIEND WHERE USER_ID_2=? AND USER_FRIEND_IS_DEL=0", id)
	if err != nil {
		return friendList, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempId int
		var tempDate string
		if err := rows.Scan(&tempId, &tempDate); err != nil {
			return friendList, err
		}
		info, err := AccGetInfo(tempId)
		if err != nil {
			return friendList, err
		}
		friendList = append(friendList, entities.Friend{
			Info: entities.AccountInfoExcludePrivateStatus{
				Email:      info.Email,
				Name:       info.Name,
				Avatar:     info.Avatar,
				Background: info.Background,
			},
			Since: tempDate,
		})
	}
	return friendList, err
}
func FriendCheck(id1 int, id2 int) (err error) {
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
	if id1 > id2 {
		id1, id2 = id2, id1
	}
	rows, err := db.Query("SELECT USER_FRIEND_SINCE FROM USER_FRIEND WHERE USER_ID_1=? AND USER_ID_2=? AND USER_FRIEND_IS_DEL=0", id1, id2)
	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	var temp string
	if err := rows.Scan(&temp); err != nil {
		return err
	}
	return err
}
