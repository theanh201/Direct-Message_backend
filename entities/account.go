package entities

// User represents a user in the system
type Account struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Tokens        []string
	TokensTimeOut []int
}

type AccountInfo struct {
	Email      string
	Name       string
	Avatar     string
	Background string
	IsPrivate  bool
}

type AccountInfoExcludePrivateStatus struct {
	Email      string
	Name       string
	Avatar     string
	Background string
}
