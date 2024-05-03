package entities

type Message struct {
	SenderEmail   string
	ReceiverEmail string
	Content       string
	Since         string
	IsEncrypt     bool
}
type MessageToMe struct {
	SenderEmail string
	Content     string
	Since       string
	IsEncrypt   bool
}

type WebsocketMessage struct {
	OnlineStatus bool   `json:"onlineStatus"`
	Token        string `json:"token"`
	Content      string `json:"content"`
	Email        string `json:"email"`
}
