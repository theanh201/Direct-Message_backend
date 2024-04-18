package controller

// func SearchControllerGetUsername(w http.ResponseWriter, r *http.Request) {
// 	// Decode json request
// 	var query entities.SearchQueryName
// 	var err error = json.NewDecoder(r.Body).Decode(&query)
// 	if err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}
// 	// Validate token
// 	valid, err := model.UserTokenValidate(query.Token)
// 	if err != nil {
// 		http.Error(w, fmt.Sprint(err), http.StatusUnauthorized)
// 		return
// 	} else if !valid {
// 		http.Error(w, "Token expired", http.StatusUnauthorized)
// 		return
// 	}
// 	// Search Username
// 	searchResult, err := model.UserInfoSearchByName(query.SearchName)
// }
