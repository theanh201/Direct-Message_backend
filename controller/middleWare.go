package controller

import (
	"crypto/rand"
	"encoding/hex"
	"net/mail"
)

const MAX_REQUEST_IMG = 10 << 20 // 10 MB

func validMail(s string) bool {
	if len(s) > 64 {
		return false
	}
	_, err := mail.ParseAddress(s)
	return err == nil
}

func valid32Byte(s string) bool {
	return len(s) == 64
}

func generateSecureRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
