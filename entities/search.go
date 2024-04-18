package entities

type SearchQueryName struct {
	Token      string `json:"token"`
	SearchName string `json:"searchName"`
	PageIdx    int    `json:"pageIdx"`
}
