package model

import (
	"DirectBackend/entities"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func UserInfoSearchByName(name string) (searchResult []entities.UserInfo, err error) {
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// qr := db.Query("")
	// _, err = db.Query(qr)
	// if err != nil {
	// 	return err
	// }
	return nil, err
}
