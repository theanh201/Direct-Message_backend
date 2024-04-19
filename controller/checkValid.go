package controller

import (
	"net/mail"
	"unicode"
)

func validMail(s string) bool {
	if len(s) > 64 {
		return false
	}
	_, err := mail.ParseAddress(s)
	return err == nil
}
func validPhoneNumber(s string) bool {
	if len(s) != 10 {
		return false
	}
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
func validToken(s string) bool {
	return len(s) == 64
}
