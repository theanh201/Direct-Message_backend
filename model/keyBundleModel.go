package model

import (
	"database/sql"
	"encoding/hex"
)

func KeyBundleUpdateIk(id int, ik []byte) (err error) {
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
	rows, err := db.Query("UPDATE USER_KEY SET USER_KEY_IK=? WHERE USER_ID=?", ik, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

func KeyBundleUpdateSpk(id int, spk []byte) (err error) {
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
	rows, err := db.Query("UPDATE USER_KEY SET USER_KEY_SPK=? WHERE USER_ID=?", spk, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return err
}

func KeyBundleUpdateOpk(id int, opk [][]byte) (err error) {
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
	rows, err := db.Query("DELETE FROM USER_OPK_KEY WHERE USER_ID=?", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	if err != nil {
		return err
	}
	// Add new opk
	for _, key := range opk {
		rows, err := db.Query("INSERT INTO USER_OPK_KEY(USER_ID, USER_OPK_KEY, USER_OPK_KEY_IS_DEL) VALUES (?, ?, 0)", id, key)
		if err != nil {
			return err
		}
		defer rows.Close()
	}
	return err
}

func KeyBundleGetByEmail(userEmail string) (ik string, spk string, opk string, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return ik, spk, opk, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return ik, spk, opk, err
	}
	// Get email id
	_, id, err := AccGetUserPassword(userEmail)
	if err != nil {
		return ik, spk, opk, err
	}
	// Get ik, spk
	rows, err := db.Query("SELECT USER_KEY_IK, USER_KEY_SPK FROM USER_KEY WHERE USER_ID=?", id)
	if err != nil {
		return ik, spk, opk, err
	}
	defer rows.Close()
	var ikByte []byte
	var spkByte []byte
	rows.Next()
	if err := rows.Scan(&ikByte, &spkByte); err != nil {
		return ik, spk, opk, err
	}
	ik = hex.EncodeToString(ikByte)
	spk = hex.EncodeToString(spkByte)
	// Get opk count
	rows, err = db.Query("SELECT COUNT(*) FROM USER_OPK_KEY WHERE USER_ID=? AND USER_OPK_KEY_IS_DEL=0;", id)
	if err != nil {
		return ik, spk, opk, err
	}
	defer rows.Close()
	var count int
	rows.Next()
	if err := rows.Scan(&count); err != nil {
		return ik, spk, opk, err
	}
	if count == 0 {
		return ik, spk, opk, err
	}
	// Get opk
	rows, err = db.Query("SELECT USER_OPK_KEY FROM USER_OPK_KEY WHERE USER_ID=? AND USER_OPK_KEY_IS_DEL=0", id)
	if err != nil {
		return ik, spk, opk, err
	}
	defer rows.Close()
	var opkByte []byte
	rows.Next()
	if err := rows.Scan(&opkByte); err != nil {
		return ik, spk, opk, err
	}
	opk = hex.EncodeToString(opkByte)
	// Mark opk as IS_DEL
	rows, err = db.Query("UPDATE USER_OPK_KEY SET USER_OPK_KEY_IS_DEL=1 WHERE USER_ID=? AND USER_OPK_KEY=?", id, opkByte)
	if err != nil {
		return ik, spk, opk, err
	}
	defer rows.Close()
	return ik, spk, opk, err
}

func KeyBundleGetIk(id int) (ik string, err error) {
	// Check DB
	db, err := sql.Open("mysql", Direct_Backend_DB)
	if err != nil {
		return ik, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return ik, err
	}
	// Get Ik
	rows, err := db.Query("SELECT USER_KEY_IK FROM USER_KEY WHERE USER_ID=?", id)
	if err != nil {
		return ik, err
	}
	defer rows.Close()
	rows.Next()
	var tempIk []byte
	err = rows.Scan(&tempIk)
	ik = hex.EncodeToString(tempIk)
	return ik, err
}
