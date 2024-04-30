package model

import (
	"DirectBackend/entities"
	"database/sql"
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
	rows, err := db.Query("INSERT INTO MESSAGE(USER_ID_FROM, USER_ID_TO, MESSAGE_CONTENT, MESSAGE_SINCE, MESSAGE_IS_ENCRYPT) VALUES (?,?,?,?,0)", idFrom, idTo, fileName, timeNow)
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
	rows, err := db.Query("SELECT USER_ID_FROM, USER_ID_TO, MESSAGE_CONTENT, MESSAGE_SINCE, MESSAGE_IS_ENCRYPT FROM MESSAGE WHERE USER_ID_TO=? OR USER_ID_FROM=?", id, id)
	if err != nil {
		return messages, err
	}
	defer rows.Close()
	selfInfo, err := AccGetInfo(id)
	if err != nil {
		return messages, err
	}
	cache := make(map[int]string) // Key: id, Value: email
	for rows.Next() {
		var tempMessage entities.Message
		var tempSender int
		var tempReceiver int
		var tempIsEncrypt []byte
		if err := rows.Scan(&tempSender, &tempReceiver, &tempMessage.Content, &tempMessage.Since, &tempIsEncrypt); err != nil {
			return messages, nil
		}
		if tempReceiver == id {
			tempMessage.ReceiverEmail = selfInfo.Email
			tempMessage.IsEncrypt = (tempIsEncrypt[0] & 1) != 0
			_, inCache := cache[tempSender]
			if !inCache {
				info, err := AccGetInfo(tempSender)
				if err != nil {
					return messages, err
				}
				cache[tempSender] = info.Email
				tempMessage.SenderEmail = info.Email
			} else {
				tempMessage.SenderEmail = cache[tempSender]
			}
		} else {
			tempMessage.SenderEmail = selfInfo.Email
			// receiver email
			_, inCache := cache[tempReceiver]
			if !inCache {
				info, err := AccGetInfo(tempReceiver)
				if err != nil {
					return messages, err
				}
				cache[tempReceiver] = info.Email
				tempMessage.ReceiverEmail = info.Email
			} else {
				tempMessage.ReceiverEmail = cache[tempReceiver]
			}
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
