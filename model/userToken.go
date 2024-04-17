package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func WriteUserTokenToDB(id int, token string, timeout string) error {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	qr := "INSERT INTO USER_TOKEN(USER_TOKEN, USER_ID, USER_TOKEN_TIMEOUT, USER_TOKEN_IS_DEL) VALUES(x'" + token + "', " + fmt.Sprint(id) + ", '" + timeout + "', 0);"
	_, err = db.Query(qr)
	if err != nil {
		return err
	}

	defer db.Close()
	return err
}

// // return valid, error
// func CheckUserTokenWithDB(username string, token string) (bool, error) {
// 	return true, nil
// }
