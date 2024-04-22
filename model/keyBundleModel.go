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
	_, id, err := AccReadUserPassword(userEmail)
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
	// FIXME IS_DEL FOR OPK
	opk = hex.EncodeToString(opkByte)
	return ik, spk, opk, err
}
