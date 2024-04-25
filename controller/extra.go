package controller

import (
	"crypto/rand"
	"encoding/hex"
	"net/mail"
)

func validMail(s string) bool {
	if len(s) > 64 {
		return false
	}
	_, err := mail.ParseAddress(s)
	return err == nil
}
func validToken(s string) bool {
	return len(s) == 64
}
func generateSecureRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
