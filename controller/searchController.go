package controller

import (
	"DirectBackend/entities"
	"DirectBackend/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func SearchControllerGetUsername(w http.ResponseWriter, r *http.Request) {
	var query entities.SearchQueryName
	var err error = json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Validate token
	valid, err := model.UserTokenValidate(query.Token)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
		return
	} else if !valid {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}

}
