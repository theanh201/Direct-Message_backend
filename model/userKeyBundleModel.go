package model

import (
	"database/sql"
	"fmt"
)

func KeyBundleUpdateIk(id int, ik string) (err error) {
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
	// Update Ik
	qr := fmt.Sprintf("UPDATE USER_KEY SET USER_KEY_IK = x'%s' WHERE USER_ID=%d", ik, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	return err
}

func KeyBundleUpdateSpk(id int, spk string) (err error) {
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
	// Update Spk
	qr := fmt.Sprintf("UPDATE USER_KEY SET USER_KEY_SPK = x'%s' WHERE USER_ID=%d", spk, id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Close
	return err
}

func KeyBundleUpdateOpk(id int, opk []string) (err error) {
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
	// Remove old opk
	qr := fmt.Sprintf("DELETE FROM USER_OPK_KEY WHERE USER_ID=%d;", id)
	_, err = db.Query(qr)
	if err != nil {
		return err
	}
	// Add new opk
	for _, key := range opk {
		qr := fmt.Sprintf("INSERT INTO USER_OPK_KEY(USER_ID, USER_OPK_KEY, USER_OPK_KEY_IS_DEL) VALUES (%d, x'%s', 0)", id, key)
		_, err = db.Query(qr)
		if err != nil {
			return err
		}
	}
	// Close
	return err
}
