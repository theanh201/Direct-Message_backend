package model

import (
	"DirectBackend/entities"
	"database/sql"
	"encoding/hex"
	"fmt"
)

// Add
func FriendRequestAdd(fromId int, toId int, ek string, opkUsed string) (err error) {
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
	qr := fmt.Sprintf("INSERT INTO USER_FRIEND_REQUEST (USER_ID_FROM, USER_ID_TO, USER_FRIEND_REQUEST_EK, USER_FRIEND_REQUEST_OPK, USER_FRIEND_REQUEST_REJECTED) VALUES(%d, %d, x'%s', x'%s', 0)", fromId, toId, ek, opkUsed)
	_, err = db.Query(qr)
	return err
}

// Get
func FriendRequestGet(id int) (friendRequest []entities.FriendRequest, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return friendRequest, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return friendRequest, err
	}
	// Get all friend request
	qr := fmt.Sprintf("SELECT USER_ID_FROM, USER_FRIEND_REQUEST_EK, USER_FRIEND_REQUEST_OPK FROM USER_FRIEND_REQUEST WHERE USER_ID_TO=%d AND USER_FRIEND_REQUEST_REJECTED=0", id)
	row, err := db.Query(qr)
	if err != nil {
		return friendRequest, err
	}
	defer row.Close()
	for row.Next() {
		var tempId int
		var tempEk []byte
		var tempOpk []byte
		if err := row.Scan(&tempId, &tempEk, &tempOpk); err != nil {
			return friendRequest, err
		}
		senderInfo, err := AccGetInfo(tempId)
		if err != nil {
			return friendRequest, err
		}
		ik, err := KeyBundleGetIk(tempId)
		if err != nil {
			return friendRequest, err
		}
		ek := hex.EncodeToString(tempEk)
		opk := hex.EncodeToString(tempOpk)
		friendRequest = append(friendRequest, entities.FriendRequest{
			From:    senderInfo,
			Ek:      ek,
			Ik:      ik,
			OpkUsed: opk,
		})
	}
	return friendRequest, err
}
