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
	Case    int    `json:"case"`
	Token   string `json:"token"`
	Content string `json:"content"`
	Email   string `json:"email"`
}
