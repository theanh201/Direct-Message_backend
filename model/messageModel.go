package model

import (
	"DirectBackend/entities"
	"database/sql"
)

func MessageFriendUnencrypt(idFrom int, idTo int, timeNow string, content string) (err error) {
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
	rows, err := db.Query("INSERT INTO MESSAGE(USER_ID_FROM, USER_ID_TO, MESSAGE_CONTENT, MESSAGE_SINCE, MESSAGE_IS_ENCRYPT, MESSAGE_IS_FILE) VALUES (?,?,?,?,0, 0)", idFrom, idTo, content, timeNow)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}
func MessageFriendEncrypt(idFrom int, idTo int, timeNow string, content string) (err error) {
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
	rows, err := db.Query("INSERT INTO MESSAGE(USER_ID_FROM, USER_ID_TO, MESSAGE_CONTENT, MESSAGE_SINCE, MESSAGE_IS_ENCRYPT, MESSAGE_IS_FILE) VALUES (?,?,?,?,1, 0)", idFrom, idTo, content, timeNow)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}
func MessageGetAfterTime(id1 int, email string, time string) (messages []entities.Message, err error) {
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
	// Get email id
	_, id2, err := AccGetUserPassword(email)
	if err != nil {
		return messages, err
	}
	// GetAll
	rows, err := db.Query("SELECT USER_ID_FROM, USER_ID_TO, MESSAGE_CONTENT, MESSAGE_SINCE, MESSAGE_IS_ENCRYPT, MESSAGE_IS_FILE FROM MESSAGE WHERE ((USER_ID_TO=? AND USER_ID_FROM=?) OR (USER_ID_TO=? AND USER_ID_FROM=?)) AND MESSAGE_SINCE>?", id1, id2, id2, id1, time)
	if err != nil {
		return messages, err
	}
	defer rows.Close()
	selfInfo, err := AccGetInfo(id1)
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		var tempMessage entities.Message
		var tempSender int
		var tempReceiver int
		var tempIsEncrypt []byte
		var tempIsFile []byte
		if err := rows.Scan(&tempSender, &tempReceiver, &tempMessage.Content, &tempMessage.Since, &tempIsEncrypt, &tempIsFile); err != nil {
			return messages, nil
		}
		// Caching id = email
		if tempReceiver == id1 {
			tempMessage.ReceiverEmail = selfInfo.Email
			tempMessage.SenderEmail = email
		} else {
			tempMessage.ReceiverEmail = email
			tempMessage.SenderEmail = selfInfo.Email
		}
		tempMessage.IsEncrypt = (tempIsEncrypt[0] & 1) != 0
		tempMessage.IsFile = (tempIsFile[0] & 1) != 0
		messages = append(messages, tempMessage)
	}
	return messages, err
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
	rows, err := db.Query("SELECT USER_ID_FROM, USER_ID_TO, MESSAGE_CONTENT, MESSAGE_SINCE, MESSAGE_IS_ENCRYPT, MESSAGE_IS_FILE FROM MESSAGE WHERE USER_ID_TO=? OR USER_ID_FROM=?", id, id)
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
		var tempIsFile []byte
		if err := rows.Scan(&tempSender, &tempReceiver, &tempMessage.Content, &tempMessage.Since, &tempIsEncrypt, &tempIsFile); err != nil {
			return messages, nil
		}
		// Caching id = email
		if tempReceiver == id {
			tempMessage.ReceiverEmail = selfInfo.Email
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
		tempMessage.IsEncrypt = (tempIsEncrypt[0] & 1) != 0
		tempMessage.IsFile = (tempIsFile[0] & 1) != 0
		messages = append(messages, tempMessage)
	}
	return messages, err
}
func MessageGetAllAfterTime(id int, time string) (messages []entities.Message, err error) {
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
	// GetAllAfterTime
	rows, err := db.Query("SELECT USER_ID_FROM, USER_ID_TO, MESSAGE_CONTENT, MESSAGE_SINCE, MESSAGE_IS_ENCRYPT, MESSAGE_IS_FILE FROM MESSAGE WHERE (USER_ID_TO=? OR USER_ID_FROM=?) AND MESSAGE_SINCE>?", id, id, time)
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
		var tempIsFile []byte
		if err := rows.Scan(&tempSender, &tempReceiver, &tempMessage.Content, &tempMessage.Since, &tempIsEncrypt, &tempIsFile); err != nil {
			return messages, nil
		}
		// Caching id = email
		if tempReceiver == id {
			tempMessage.ReceiverEmail = selfInfo.Email
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
		tempMessage.IsEncrypt = (tempIsEncrypt[0] & 1) != 0
		tempMessage.IsFile = (tempIsFile[0] & 1) != 0
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
func MessageDelete(id int, time string) (err error) {
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
	rows, err := db.Query("UPDATE MESSAGE SET MESSAGE_CONTENT='Deleted' WHERE USER_ID_FROM=? AND MESSAGE_SINCE=?", id, time)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}
