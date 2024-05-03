package controller

import (
	"DirectBackend/model"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const MAX_REQUEST_IMG = 10 << 20 // 10 MB

func validMail(s string) bool {
	if len(s) > 64 {
		return false
	}
	_, err := mail.ParseAddress(s)
	return err == nil
}

func generateSecureRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
func validateToken(token string) (id int, err error) {
	if len(token) != 64 {
		return id, fmt.Errorf("token not 32 byte")
	}
	tokenByte, err := hex.DecodeString(token)
	if err != nil {
		return id, err
	}
	id, err = model.UserTokenValidate(tokenByte)
	if err != nil {
		return id, err
	}
	return id, err
}
func convert32Byte(key string) (keyByte []byte, err error) {
	if len(key) != 64 {
		return keyByte, fmt.Errorf("token not 32 byte")
	}
	keyByte, err = hex.DecodeString(key)
	if err != nil {
		return keyByte, err
	}
	return keyByte, err
}
