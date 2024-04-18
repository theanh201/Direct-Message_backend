package entities

// User represents a user in the system
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Tokens        []string
	TokensTimeOut []int
}

type AccountInfo struct {
	UserEmail       string
	UserPhoneNumber string
	UserName        string
	UserAvatar      string
	UserBackground  string
	UserIsPrivate   bool
}
