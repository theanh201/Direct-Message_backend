package model

import (
	"DirectBackend/entities"
	"database/sql"
	"encoding/hex"
)

// Add
func FriendRequestAdd(fromId int, toId int, ek []byte, opkUsed []byte) (err error) {
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
	rows, err := db.Query("INSERT INTO USER_FRIEND_REQUEST (USER_ID_FROM, USER_ID_TO, USER_FRIEND_REQUEST_EK, USER_FRIEND_REQUEST_OPK, USER_FRIEND_REQUEST_IS_DEL) VALUES(?, ?, ?, ?, 0)", fromId, toId, ek, opkUsed)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

// Update
func FriendRequestUpdateReject(email string, id int) (err error) {
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
	_, id2, err := AccGetUserPassword(email)
	if err != nil {
		return err
	}
	rows, err := db.Query("UPDATE USER_FRIEND_REQUEST SET USER_FRIEND_REQUEST_IS_DEL=1 WHERE USER_ID_TO=? AND USER_ID_FROM=?", id, id2)
	if err != nil {
		return err
	}
	defer rows.Close()
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
	rows, err := db.Query("SELECT USER_ID_FROM, USER_FRIEND_REQUEST_EK, USER_FRIEND_REQUEST_OPK FROM USER_FRIEND_REQUEST WHERE USER_ID_TO=? AND USER_FRIEND_REQUEST_IS_DEL=0", id)
	if err != nil {
		return friendRequest, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempId int
		var tempEk []byte
		var tempOpk []byte
		if err := rows.Scan(&tempId, &tempEk, &tempOpk); err != nil {
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
