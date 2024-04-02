package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Direct_Backend_DB string = "user:password1234@tcp(127.0.0.1:3306)/Direct_Backend_DB"

func writeUserToDB(username string, password string) int {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	qr := "INSERT INTO USER_ACCOUNT (USER_ACCOUNT_USERNAME, USER_ACCOUNT_PASSWORD) VALUES('" + username + "','" + password + "')"
	fmt.Println(qr)
	_, err = db.Query(qr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return 0
}
