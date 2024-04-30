package entities

type Message struct {
	SenderEmail   string
	ReceiverEmail string
	Content       string
	Since         string
	IsEncrypt     bool
}
