package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var onlineConn = make(map[int]*websocket.Conn)

func MessageFriendUnencrypt(w http.ResponseWriter, r *http.Request) {
	// Web socket upgrade
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	var idFrom int
	for {
		// Read message from client
		_, jsonMessage, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		// Decode message
		var message entities.WebsocketMessage
		err = json.Unmarshal(jsonMessage, &message)
		if err != nil {
			log.Println(err)
			break
		}
		// Validate token
		idFrom, err = validateToken(message.Token)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			break
		}
		// Add to online if not online
		if !message.OnlineStatus {
			onlineConn[idFrom] = conn
			conn.WriteMessage(websocket.TextMessage, []byte("You are now online"))
			continue
		}
		// Get current time
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		// Validate mail
		if !validMail(message.Email) {
			conn.WriteMessage(websocket.TextMessage, []byte("in valid mail"))
			break
		}
		_, idTo, err := model.AccGetUserPassword(message.Email)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			break
		}
		// Check if 2 are friend
		err = model.FriendCheck(idFrom, idTo)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			break
		}
		// Upload to db
		err = model.MessageFriendUnencrypt(idFrom, idTo, timeNow, message.Content)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			break
		}
		// Send to user if online
		toConn, isOnline := onlineConn[idTo]
		if isOnline {
			err = toConn.WriteMessage(websocket.TextMessage, []byte(message.Content))
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Message not dilivered"))
			}
		}
	}
	delete(onlineConn, idFrom)
}
