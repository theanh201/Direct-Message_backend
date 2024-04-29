package model

import (
	"database/sql"
	"encoding/hex"
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
	rows, err := db.Query(qr)
	if err != nil {
		return err
	}
	defer rows.Close()
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
	rows, err := db.Query(qr)
	if err != nil {
		return err
	}
	defer rows.Close()
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
	qr := fmt.Sprintf("DELETE FROM USER_OPK_KEY WHERE USER_ID=%d", id)
	rows, err := db.Query(qr)
	if err != nil {
		return err
	}
	defer rows.Close()
	if err != nil {
		return err
	}
	// Add new opk
	for _, key := range opk {
		qr := fmt.Sprintf("INSERT INTO USER_OPK_KEY(USER_ID, USER_OPK_KEY, USER_OPK_KEY_IS_DEL) VALUES (%d, x'%s', 0)", id, key)
		rows, err := db.Query(qr)
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
	qr := fmt.Sprintf("SELECT USER_KEY_IK, USER_KEY_SPK FROM USER_KEY WHERE USER_ID=%d", id)
	rows, err := db.Query(qr)
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
	// Get opk
	qr = fmt.Sprintf("SELECT USER_OPK_KEY FROM USER_OPK_KEY WHERE USER_ID=%d AND USER_OPK_KEY_IS_DEL=0", id)
	rows, err = db.Query(qr)
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
	qr = fmt.Sprintf("UPDATE USER_OPK_KEY SET USER_OPK_KEY_IS_DEL=1 WHERE USER_ID=%d AND USER_OPK_KEY=x'%s'", id, opk)
	rows, err = db.Query(qr)
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
	qr := fmt.Sprintf("SELECT USER_KEY_IK FROM USER_KEY WHERE USER_ID=%d", id)
	rows, err := db.Query(qr)
	if err != nil {
		return ik, err
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&ik)
	return ik, err
}
