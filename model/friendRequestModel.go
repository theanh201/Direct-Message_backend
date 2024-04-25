package model

import (
	"DirectBackend/entities"
	"database/sql"
	"encoding/hex"
	"fmt"
)

// Get
func GetFriendRequest(id int) (friendRequest []entities.FriendRequestResponse, err error) {
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
	qr := fmt.Sprintf("SELECT (USER_ID_FROM, USER_FRIEND_REQUEST_EK) FROM USER_FRIEND_REQUEST WHERE USER_ID_TO=%d AND USER_FRIEND_REQUEST_REJECTED=0", id)
	row, err := db.Query(qr)
	if err != nil {
		return friendRequest, err
	}
	defer row.Close()
	for row.Next() {
		var tempId int
		var tempEk []byte
		if err := row.Scan(&tempId, &tempEk); err != nil {
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
		friendRequest = append(friendRequest, entities.FriendRequestResponse{
			RequestFrom: senderInfo,
			RequestEk:   ek,
			RequestIk:   ik,
		})
	}
	return friendRequest, err
}
