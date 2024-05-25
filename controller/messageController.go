package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var onlineConn = make(map[int]*websocket.Conn)

// WebSocket
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
		log.Printf("recv: %s", jsonMessage)
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
		switch message.Case {
		case 0:
			onlineConn[idFrom] = conn
			conn.WriteMessage(websocket.TextMessage, []byte("You are now online"))
			continue
		case 1:
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
				sender, err := model.AccGetInfo(idFrom)
				if err != nil {
					conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
					break
				}
				// Define your data
				data := entities.Message{
					IsEncrypt:     false,
					IsFile:        false,
					Content:       message.Content,
					ReceiverEmail: message.Email,
					SenderEmail:   sender.Email,
					Since:         timeNow,
				}
				// Marshal data to JSON
				jsonData, err := json.Marshal(data)
				if err != nil {
					fmt.Println("Error marshalling JSON:", err)
					return
				}
				err = toConn.WriteMessage(websocket.TextMessage, jsonData)
				// err = toConn.WriteJSON(jsonData)
				if err != nil {
					conn.WriteMessage(websocket.TextMessage, []byte("Message not dilivered"))
				}
			}
		// send encrypt message
		case 2:
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
			err = model.MessageFriendEncrypt(idFrom, idTo, timeNow, message.Content)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
				break
			}
			// Send to user if online
			toConn, isOnline := onlineConn[idTo]
			if isOnline {
				sender, err := model.AccGetInfo(idFrom)
				if err != nil {
					conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
					break
				}
				// Define your data
				data := entities.Message{
					IsEncrypt:     true,
					IsFile:        false,
					Content:       message.Content,
					ReceiverEmail: message.Email,
					SenderEmail:   sender.Email,
					Since:         timeNow,
				}
				// Marshal data to JSON
				jsonData, err := json.Marshal(data)
				if err != nil {
					fmt.Println("Error marshalling JSON:", err)
					return
				}
				err = toConn.WriteMessage(websocket.TextMessage, jsonData)
				if err != nil {
					conn.WriteMessage(websocket.TextMessage, []byte("Message not dilivered"))
				}
			}
		}
	}
	delete(onlineConn, idFrom)
}

// GET
func MessageGetAll(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get message
	messages, err := model.MessageGetAll(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

func MessageGetAllAfterTime(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	time := mux.Vars(r)["time"]
	messages, err := model.MessageGetAllAfterTime(id, time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(messages)
}
