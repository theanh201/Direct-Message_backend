package model

import (
	"DirectBackend/entities"
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
func MessageGetAll(id int) (messages []entities.Message, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return messages, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return messages, err
	}
	// GetAll
	qr := fmt.Sprintf("SELECT USER_ID_FROM, MESSAGE_CONTENT, MESSAGE_SINCE, MESSAGE_IS_ENCRYPT FROM MESSAGE WHERE USER_ID_TO=%d", id)
	rows, err := db.Query(qr)
	if err != nil {
		return messages, err
	}
	defer rows.Close()
	cache := make(map[int]string) // Key: id, Value: email
	for rows.Next() {
		var tempMessage entities.Message
		var tempId int
		var tempIsEncrypt []byte
		if err := rows.Scan(&tempId, &tempMessage.Content, &tempMessage.Since, &tempIsEncrypt); err != nil {
			return messages, nil
		}
		tempMessage.IsEncrypt = (tempIsEncrypt[0] & 1) != 0
		_, inCache := cache[tempId]
		if !inCache {
			// Get email and add to cache
			info, err := AccGetInfo(tempId)
			if err != nil {
				return messages, err
			}
			cache[tempId] = info.Email
			tempMessage.SenderEmail = info.Email
		} else {
			// Get from cache
			tempMessage.SenderEmail = cache[tempId]
		}
		messages = append(messages, tempMessage)
	}
	return messages, err
}

func MessageGetContentPermission(contentName string) (idFrom int, idTo int, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return idFrom, idTo, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return idFrom, idTo, err
	}
	rows, err := db.Query("SELECT USER_ID_FROM, USER_ID_TO FROM MESSAGE WHERE MESSAGE_CONTENT=?", contentName)
	if err != nil {
		return idFrom, idTo, err
	}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&idFrom, &idTo); err != nil {
		return idFrom, idTo, err
	}
	return idFrom, idTo, err
}
