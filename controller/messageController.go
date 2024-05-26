package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ZEGOCLOUD/zego_server_assistant/token/go/src/token04"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var onlineConn = make(map[int][]*websocket.Conn)

func getCallToken(id int) (token string) {
	type RtcRoomPayLoad struct {
		RoomId       string      `json:"room_id"`        // Room ID; used for strong verification of the room ID of the interface
		Privilege    map[int]int `json:"privilege"`      // Permission switch list; used for strong verification of operation permissions of the interface
		StreamIdList []string    `json:"stream_id_list"` // Stream list; used for strong verification of the stream ID of the interface; can be empty; if empty, no stream ID verification is performed
	}
	var appId uint32 = 2105949447
	serverSecret := "c391b3f6184da3a7a0115c7ba7363f8f"
	userId := fmt.Sprint(id)
	roomId := fmt.Sprint(id)
	var effectiveTimeInSeconds int64 = 3600
	privilege := make(map[int]int)
	privilege[token04.PrivilegeKeyLogin] = token04.PrivilegeEnable    // Allow room login
	privilege[token04.PrivilegeKeyPublish] = token04.PrivilegeDisable // Do not allow streaming
	payloadData := &RtcRoomPayLoad{
		RoomId:       roomId,
		Privilege:    privilege,
		StreamIdList: nil,
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Generate token
	token, err = token04.GenerateToken04(appId, userId, serverSecret, effectiveTimeInSeconds, string(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
	return token
}

// WebSocket
func deleteElement(slice []*websocket.Conn, remove *websocket.Conn) (result []*websocket.Conn) {
	for _, elem := range slice {
		if elem != remove {
			result = append(result, elem)
		}
	}
	return result
}
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
		switch message.Case {
		case 0:
			onlineConn[idFrom] = deleteElement(onlineConn[idFrom], conn)
			onlineConn[idFrom] = append(onlineConn[idFrom], conn)
			conn.WriteMessage(websocket.TextMessage, []byte("You are now online"))
			// conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Call token: %s", getCallToken())))
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
			// Send time to sender
			conn.WriteMessage(websocket.TextMessage, []byte(timeNow))
			// Send to other client
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
				break
			}
			for _, toConn := range onlineConn[idFrom] {
				if conn == toConn {
					continue
				}
				err = toConn.WriteMessage(websocket.TextMessage, jsonData)
				if err != nil {
					conn.WriteMessage(websocket.TextMessage, []byte("Cant dilivered to other client"))
				}
			}
			// Send to user if online
			toConns, isOnline := onlineConn[idTo]
			if isOnline {
				for _, toConn := range toConns {
					err = toConn.WriteMessage(websocket.TextMessage, jsonData)
					if err != nil {
						conn.WriteMessage(websocket.TextMessage, []byte("Message not dilivered"))
					}
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
			// Send time to sender
			conn.WriteMessage(websocket.TextMessage, []byte(timeNow))
			// Send to other client
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
				break
			}
			for _, toConn := range onlineConn[idFrom] {
				if conn == toConn {
					continue
				}
				err = toConn.WriteMessage(websocket.TextMessage, jsonData)
				if err != nil {
					conn.WriteMessage(websocket.TextMessage, []byte("Cant dilivered to other client"))
				}
			}
			// Send to user if online
			toConns, isOnline := onlineConn[idTo]
			if isOnline {
				for _, toConn := range toConns {
					err = toConn.WriteMessage(websocket.TextMessage, jsonData)
					if err != nil {
						conn.WriteMessage(websocket.TextMessage, []byte("Message not dilivered"))
					}
				}
			}
		case 3:
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Call token: %s", getCallToken(idFrom))))
		}
	}
	onlineConn[idFrom] = deleteElement(onlineConn[idFrom], conn)
}

// GET
func MessageGetByEmail(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id1, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get email
	email := mux.Vars(r)["email"]
	if !validMail(email) {
		http.Error(w, "valid email not found", http.StatusBadRequest)
		return
	}
	messages, err := model.MessageGetAfterTime(id1, email, "2000-01-01 20:00:00")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}
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
func MessageGetByEmailAfterTime(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get email
	email := mux.Vars(r)["email"]
	if !validMail(email) {
		http.Error(w, "valid email not found", http.StatusBadRequest)
		return
	}
	// Get time
	time := mux.Vars(r)["time"]
	time = strings.Replace(time, "_", " ", -1)
	messages, err := model.MessageGetAfterTime(id, email, time)
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
	time = strings.Replace(time, "_", " ", -1)
	messages, err := model.MessageGetAllAfterTime(id, time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(messages)
}
func MessageDelete(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	time := mux.Vars(r)["time"]
	time = strings.Replace(time, "_", " ", -1)
	err = model.MessageDelete(id, time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := map[string]string{"message": "delete success"}
	json.NewEncoder(w).Encode(response)
}
