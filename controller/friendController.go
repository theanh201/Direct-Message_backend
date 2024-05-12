package controller

import (
	"DirectBackend/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func FriendGet(w http.ResponseWriter, r *http.Request) {
	// Validate token
	id, err := validateToken(mux.Vars(r)["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get friend list
	firendList, err := model.FriendGet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(firendList)
}
